package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"root/root/database"
	"root/root/errors"
	"root/root/internal/config"
	"strconv"
)

func ShowNotes(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get("page")
	pageNumber, _ := strconv.Atoi(pageNumberStr)
	if pageNumber < 1 {
		pageNumber = 1
	}

	perPage := 10
	offset := (pageNumber - 1) * perPage

	query := "select * from notes order by id_note limit $1 offset $2"
	rows, err := database.DB.Query(query, perPage, offset)
	errors.CheckError(err)

	defer rows.Close()

	var notes []config.Notes
	for rows.Next() {
		var note config.Notes
		err := rows.Scan(&note.Id, &note.Title, &note.Description)
		errors.CheckError(err)

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	countQuery := "select count(*) from notes"
	var totalNotes int
	err = database.DB.QueryRow(countQuery).Scan(&totalNotes)
	errors.CheckError(err)

	totalPages := totalNotes/perPage + 1

	page := config.PageNotes{
		Notes:        notes,
		CurrentPage:  pageNumber,
		TotalPage:    totalPages,
		CountTickets: totalNotes,
	}

	tmpl, err := template.New("showNotes.html").Funcs(template.FuncMap{
		"printf": func(format string, a ...interface{}) string {
			return fmt.Sprintf(format, a...)
		},
		"minus": minusNotes,
		"plus":  plusNotes,
	}).ParseFiles("templates/showNotes.html", "templates/header.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "show_notes", page)

}

func minusNotes(a, b int) int {
	return a - b
}

func plusNotes(a, b int) int {
	return a + b
}
