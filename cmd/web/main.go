package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nuttchai/WebApp-Golang/pkg/config"
	"github.com/nuttchai/WebApp-Golang/pkg/handlers"
	"github.com/nuttchai/WebApp-Golang/pkg/render"
)

// NOTE: It is convention that web folder is typically where Web Applications offten have their main function

/* NOTE: go.mod file tells the compiler that the application uses go modules.
it's like the package.json file used in Node.js dependency management */
const portNumber string = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	/* NOTE:
	ListenAndServe: listen HTTP with port and handler (nil)
	It returns error, but we ignore it in this case by putting _ as variable */
	_ = http.ListenAndServe(portNumber, nil)
}
