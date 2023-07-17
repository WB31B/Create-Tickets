package services

import (
	"fmt"
	"html/template"
	"net/http"
	"root/root/database"
	"root/root/errors"
	"root/root/internal/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	countQuery := "select count(*) from notes"
	var totalNotes int
	err = database.DB.QueryRow(countQuery).Scan(&totalNotes)
	errors.CheckError(err)

	rows, err := database.DB.Query(fmt.Sprintf("select * from notes where id_note = '%d'", totalNotes))
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

	t.ExecuteTemplate(w, "index", showNote)
}
