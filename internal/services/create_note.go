package services

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createNote.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "createNote", nil)
}
