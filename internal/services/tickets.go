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

func Tickets(w http.ResponseWriter, r *http.Request) {
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

	var articles []config.Article
	for rows.Next() {
		var article config.Article
		err := rows.Scan(&article.Id, &article.Username, &article.Information, &article.Ticket)
		errors.CheckError(err)

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	countQuery := "select count(*) from info"
	var totalArticles int
	err = database.DB.QueryRow(countQuery).Scan(&totalArticles)
	errors.CheckError(err)

	totalPages := totalArticles/perPage + 1

	page := config.Page{
		Article:      articles,
		CurrentPage:  pageNumber,
		TotalPage:    totalPages,
		CountTickets: totalArticles,
	}

	tmpl, err := template.New("tickets.html").Funcs(template.FuncMap{
		"printf": func(format string, a ...interface{}) string {
			return fmt.Sprintf(format, a...)
		},
		"minus": minus,
		"plus":  plus,
	}).ParseFiles("templates/tickets.html", "templates/header.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "tickets", page)
}

func minus(a, b int) int {
	return a - b
}

func plus(a, b int) int {
	return a + b
}
