package app

import (
	"net/http"
)

// TODO: Create templates for the base website
// organize like in HUGO

func (app *App) setupHandlers() {
	static := http.Dir("../static")
	styles := http.Dir("../styles")
	app.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static)))
	app.router.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(styles)))
	app.router.Handle("GET /posts/{slug}", http.HandlerFunc(app.handleBlogPosts))
}

func (app *App) handleBlogPosts(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, exists := app.posts[slug]
	if !exists {
		http.NotFound(w, r)
		return
	}

	template := app.templates[post.TemplateName]
	if template == nil {
		http.Error(w, "nil template", http.StatusInternalServerError)
		return
	}

	cssLinks := []string{
		"/styles/output.css",
		"/styles/prism-gruvbox-dark.css",
		"/styles/prism-gruvbox-light.css",
	}

	templateData := struct {
		CSSLinks []string
	}{
		CSSLinks: cssLinks,
	}

	err := template.Execute(w, templateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
