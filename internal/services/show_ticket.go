package services

import (
	"fmt"
	"html/template"
	"net/http"
	"root/root/database"
	"root/root/errors"
	"root/root/internal/config"

	"github.com/gorilla/mux"
)

func ShowTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/ticket.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	rows, err := database.DB.Query(fmt.Sprintf("select * from info where id = '%s'", vars["id"]))
	errors.CheckError(err)

	showTicket := config.Tickets{}

	for rows.Next() {
		var post config.Tickets
		err := rows.Scan(&post.Id, &post.Username, &post.Information, &post.Ticket)
		if err != nil {
			continue
		}

		showTicket = post
	}

	t.ExecuteTemplate(w, "ticket", showTicket)
}
