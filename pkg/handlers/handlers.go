package handlers

import (
	"golang-aplication-project/pkg/config"
	"golang-aplication-project/pkg/models"
	"golang-aplication-project/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository 

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//New create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the package home handler
func (m *Repository) Home (w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about package handler
func (m *Repository) About (w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"
	
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
