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

func Show_ticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/ticket.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	rows, err := database.DB.Query(fmt.Sprintf("select * from info where id = '%s'", vars["id"]))
	errors.CheckError(err)

	showTicket := config.Article{}

	for rows.Next() {
		var post config.Article
		err := rows.Scan(&post.Id, &post.Username, &post.Information, &post.Ticket)
		if err != nil {
			continue
		}

		showTicket = post
	}

	t.ExecuteTemplate(w, "ticket", showTicket)
}
