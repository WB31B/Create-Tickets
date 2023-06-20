package services

import (
	"fmt"
	"html/template"
	"net/http"
)

func Show_form(w http.ResponseWriter, r *http.Request) {
	sfticket, err := template.ParseFiles("templates/showForm.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	sfticket.ExecuteTemplate(w, "showForm", nil)
}
