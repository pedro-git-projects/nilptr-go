package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pedro-git-projects/nilptr/src/blog"
	"github.com/pedro-git-projects/nilptr/src/processor"
)

type App struct {
	config      Config
	router      *http.ServeMux
	errorLogger *log.Logger
	infoLogger  *log.Logger

	templates map[string]*template.Template // slug -> template
	posts     map[string]*blog.BlogPost     // slug -> blogpost
	processor *processor.Processor
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

func initializeLoggers(app *App) {
	app.infoLogger = log.New(os.Stdout, fmt.Sprintf("%s::INFO ", app.config.env), log.Ldate|log.Ltime)
	app.errorLogger = log.New(os.Stderr, fmt.Sprintf("%s::ERROR", app.config.env), log.Ldate|log.Ltime)
}

func New() *App {
	app := &App{
		config:    parseFlags(),
		router:    http.NewServeMux(),
		templates: make(map[string]*template.Template),
		processor: processor.New(),
	}
	initializeLoggers(app)
	app.setupHandlers()
	return app
}

func (app *App) loadTemplates(templatesDir string) {
	files, err := filepath.Glob(filepath.Join(templatesDir, "*.html"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		templateName := strings.TrimSuffix(filepath.Base(file), ".html")
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			panic(err)
		}
		app.templates[templateName] = tmpl
	}
}

func (app *App) Start() error {
	app.infoLogger.Printf("Starting %s server on port :%d\n", app.config.env, app.config.port)
	app.processor.ProcessAndSave("test.html")
	return http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.router)
}
