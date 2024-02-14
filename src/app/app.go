package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type App struct {
	config      Config
	router      *http.ServeMux
	templates   map[string]*template.Template // slug -> template
	errorLogger *log.Logger
	infoLogger  *log.Logger
	mdParser    goldmark.Markdown
}

func parseFlags() Config {
	defaultEnv := Development
	defaultPort := 8080

	var flagEnv string
	var flagPort int

	flag.StringVar(&flagEnv, "env", string(defaultEnv), "Environment")
	flag.IntVar(&flagPort, "port", defaultPort, "Port number")
	flag.Parse()

	envFromEnvVar := os.Getenv("ENV")
	var env Environment
	switch envFromEnvVar {
	case "development":
		env = Development
	case "staging":
		env = Staging
	case "production":
		env = Production
	default:
		env = defaultEnv
	}

	return Config{
		env:  env,
		port: flagPort,
	}
}

func initializeParser(app *App) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	app.mdParser = md
}

func initializeLoggers(app *App) {
	app.infoLogger = log.New(os.Stdout, fmt.Sprintf("%s::INFO ", app.config.env), log.Ldate|log.Ltime)
	app.errorLogger = log.New(os.Stderr, fmt.Sprintf("%s::ERROR", app.config.env), log.Ldate|log.Ltime)
}

func New() *App {
	app := &App{
		config:    parseFlags(),
		router:    http.NewServeMux(),
		templates: make(map[string]*template.Template),
	}
	initializeParser(app)
	initializeLoggers(app)
	app.setupHandlers()
	return app
}

func (app *App) Start() error {
	app.infoLogger.Printf("Starting %s server on port :%d\n", app.config.env, app.config.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.router)
}
