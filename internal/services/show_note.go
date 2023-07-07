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

func ShowNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/note.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	rows, err := database.DB.Query(fmt.Sprintf("select * from notes where id_note = '%s'", vars["id"]))
	errors.CheckError(err)

	showNote := config.Notes{}

	for rows.Next() {
		var post config.Notes
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			continue
		}

		showNote = post
	}

	t.ExecuteTemplate(w, "note", showNote)
}
