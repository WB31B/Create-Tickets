package services

import (
	"fmt"
	"net/http"
	"root/root/database"
	"root/root/errors"
)

func SaveTicket(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	information := r.FormValue("information")
	ticket := r.FormValue("ticket")

	if username == "" || information == "" || ticket == "" {
		fmt.Fprintf(w, "Not all data is filled!")
	} else {
		addTicketToDB := `insert into "info"("name", "information", "ticket") values($1, $2, $3)`
		iticket := fmt.Sprintf("#%s", ticket)
		_, err := database.DB.Exec(addTicketToDB, username, information, iticket)
		errors.CheckError(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SaveNote(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	if title == "" || description == "" {
		fmt.Fprintf(w, "Not all data is filled!")
	} else {
		addNoteToDB := `insert into "notes"("title", "description") values($1, $2)`
		_, err := database.DB.Exec(addNoteToDB, title, description)
		errors.CheckError(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
