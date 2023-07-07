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

func ShowTickets(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get("page")
	pageNumber, _ := strconv.Atoi(pageNumberStr)
	if pageNumber < 1 {
		pageNumber = 1
	}

	perPage := 10
	offset := (pageNumber - 1) * perPage

	query := "select * from info order by id limit $1 offset $2"
	rows, err := database.DB.Query(query, perPage, offset)
	errors.CheckError(err)

	defer rows.Close()

	var tickets []config.Tickets
	for rows.Next() {
		var ticket config.Tickets
		err := rows.Scan(&ticket.Id, &ticket.Username, &ticket.Information, &ticket.Ticket)
		errors.CheckError(err)

		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	countQuery := "select count(*) from info"
	var totalTickets int
	err = database.DB.QueryRow(countQuery).Scan(&totalTickets)
	errors.CheckError(err)

	totalPages := totalTickets/perPage + 1

	page := config.PageTickets{
		Tickets:      tickets,
		CurrentPage:  pageNumber,
		TotalPage:    totalPages,
		CountTickets: totalTickets,
	}

	tmpl, err := template.New("showTickets.html").Funcs(template.FuncMap{
		"printf": func(format string, a ...interface{}) string {
			return fmt.Sprintf(format, a...)
		},
		"minus": minusTickets,
		"plus":  plusTickets,
	}).ParseFiles("templates/showTickets.html", "templates/header.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "show_tickets", page)
}

func minusTickets(a, b int) int {
	return a - b
}

func plusTickets(a, b int) int {
	return a + b
}
