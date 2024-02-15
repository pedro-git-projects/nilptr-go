package blog

import (
	"html/template"
	"regexp"
	"strings"
	"time"
)

type BlogPost struct {
	Slug         string             `json:"slug"`
	Title        string             `json:"title"`
	TemplateName string             `json:"-"`
	Content      *template.Template `json:"content,omitempty"`
	Day          int                `json:"day"`
	Month        time.Month         `json:"month"`
	Year         int                `json:"year"`
}

func generateSlug(title string) string {
	slug := strings.ReplaceAll(strings.ToLower(title), " ", "-")
	regexpForSlug := regexp.MustCompile("[^a-z0-9-]")
	slug = regexpForSlug.ReplaceAllString(slug, "")

	return slug
}

func NewPost(title string, templateName string, content *template.Template) *BlogPost {
	return &BlogPost{
		Slug:         generateSlug(title),
		Title:        title,
		TemplateName: templateName,
		Content:      content,
		Day:          time.Now().Day(),
		Month:        time.Now().Month(),
		Year:         time.Now().Year(),
	}
}
