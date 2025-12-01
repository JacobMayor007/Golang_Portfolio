package routes

import (
	"html/template"
	"myportfolio/model"
	"net/http"
)

var templates *template.Template

func init() {

	var err error
	templates, err = template.ParseFiles(
		"index.html",
		"projects.html",
		"about.html",
	)
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := model.Dev
	data.Title = "Home - " + model.Dev.Fullname
	data.Page = "home"
	renderTemplate(w, "index.html", data)
}

func Projects(w http.ResponseWriter, r *http.Request) {
	data := model.Dev
	data.Title = "Projects - " + model.Dev.Fullname
	data.Page = "projects"
	renderTemplate(w, "projects.html", data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := model.Dev
	data.Title = "About - " + model.Dev.Fullname
	data.Page = "about"
	renderTemplate(w, "about.html", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data model.Developer) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
