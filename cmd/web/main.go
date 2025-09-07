package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version    = "1.0.0"
	cssVersion = "1"
)

type config struct {
	port int
	env  string
	api  string
	db   struct {
		// How to connect to DB
		dataSrcName string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting HTTP Server in %s Mode On Port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "Server Port to Listen on")
	flag.StringVar(&config.env, "env", "development", "Application Environment {development|production}")
	flag.StringVar(&config.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	config.stripe.key = os.Getenv("STRIPE_KEY")
	config.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	temCache := make(map[string]*template.Template)

	app := &application{
		config:        config,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: temCache,
		version:       version,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
