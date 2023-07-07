package services

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createTicket.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create_ticket", nil)
}
