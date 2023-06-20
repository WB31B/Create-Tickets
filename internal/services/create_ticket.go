package services

import (
	"fmt"
	"html/template"
	"net/http"
)

func Create_ticket(w http.ResponseWriter, r *http.Request) {
	cticket, err := template.ParseFiles("templates/createTicket.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	cticket.ExecuteTemplate(w, "createTicket", nil)
}
