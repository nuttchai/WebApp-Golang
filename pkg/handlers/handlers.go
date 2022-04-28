package handlers

import (
	"net/http"

	"github.com/nuttchai/go-webapp-starting-template/pkg/config"
	"github.com/nuttchai/go-webapp-starting-template/pkg/models"
	"github.com/nuttchai/go-webapp-starting-template/pkg/render"
)

// IMPORTANT: if function start with CAP letter, it means that that function will be ACCESSIBLE outside the package

/* NOTE: HTTP HandleFunc must have at least TWO arugments
First is path URL
Second is inline function with args: response writer and request pointer */

// NOTE: Repository below is called Repository pattern that allows us to swap components out of our application with a minimal changes required to the code base
// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// NOTE: Go convention for commenting: function name should appear at the beginning
// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{}) // TemplateData{} is how we intialize an empty TemplateData
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// NOTE: Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// NOTE: Send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
