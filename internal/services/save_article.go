package services

import (
	"fmt"
	"net/http"
	"root/root/database"
	"root/root/errors"
)

func Save_article(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	information := r.FormValue("information")
	ticket := r.FormValue("ticket")

	if username == "" || information == "" || ticket == "" {
		fmt.Fprintf(w, "Not all data is filled!")
	} else {
		addInfoToDB := `insert into "info"("name", "information", "ticket") values($1, $2, $3)`
		iticket := fmt.Sprintf("#%s", ticket)
		_, err := database.DB.Exec(addInfoToDB, username, information, iticket)
		errors.CheckError(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
