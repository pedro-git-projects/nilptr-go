package app

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func (app *App) setupHandlers() {
	static := http.Dir("../static")
	styles := http.Dir("../styles")
	app.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static)))
	app.router.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(styles)))
	app.router.Handle("/blog", http.HandlerFunc(app.handleBlogPost))
}

func (app *App) handleBlogPost(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("../posts/sample.md")
	if err != nil {
		app.errorLogger.Println("Failed to read file: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := app.mdParser.Convert(content, &buf); err != nil {
		app.errorLogger.Println("Failed to parse Markdown: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, buf.String())

}
