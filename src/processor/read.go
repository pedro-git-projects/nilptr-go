package processor

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func (p *Processor) ReadMD() ([]byte, error) {
	content, err := os.ReadFile("../content/posts/sample.md")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := p.parser.Convert(content, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *Processor) ProcessHTML(htmlBytes []byte) ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(htmlBytes))
	if err != nil {
		return nil, err
	}

	err = walkDOM(doc.FirstChild, func(n *html.Node) error {
		switch n.Type {
		case html.ElementNode:
			switch n.Data {

			case "pre":
				hasClass := false
				for _, attr := range n.Attr {
					if attr.Key == "class" {
						attr.Val = fmt.Sprintf("%s my-4 p-4 rounded-md", attr.Val)
						hasClass = true
						break
					}
				}
				if !hasClass {
					n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "my-4 p-4 rounded-md"})
				}

			case "h1":
				hasClassName := false
				for _, attr := range n.Attr {
					if attr.Key == "class" {
						attr.Val = fmt.Sprintf("%s text-3xl font-extrabold sm:text-5xl", attr.Val)
						hasClassName = true
						break
					}
				}
				if !hasClassName {
					n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "text-3xl font-extrabold sm:text-5xl"})
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	out.WriteString("{{template \"base\" .}}\n")
	out.WriteString("{{define \"main\" }}\n")

	if err := html.Render(&out, doc); err != nil {
		return nil, err
	}

	out.WriteString("{{end}}\n")

	return out.Bytes(), nil
}

func walkDOM(n *html.Node, f func(n *html.Node) error) error {
	if err := f(n); err != nil {
		return err
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if err := walkDOM(child, f); err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) WriteToFile(processedHTML []byte, filename string) error {
	os.WriteFile(fmt.Sprintf("../pages/%s", filename), []byte(processedHTML), 0666)
	return nil
}

func (p *Processor) ProcessAndSave(filename string) error {
	html, err := p.ReadMD()
	if err != nil {
		return err
	}

	processedHTML, err := p.ProcessHTML(html)
	if err != nil {
		return err
	}

	err = p.WriteToFile(processedHTML, filename)
	if err != nil {
		return err
	}

	return nil
}
