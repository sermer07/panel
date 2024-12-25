package handlers

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "World Cup",
		Message: "FIFA will never regret it",
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}
