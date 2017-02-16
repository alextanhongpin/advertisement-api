package common

import (
	"html/template"
	"net/http"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["create-campaign"] = template.Must(template.ParseFiles("templates/campaign/create.html", "templates/base.html"))
}

func RenderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist,", http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template, viewModel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
