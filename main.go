package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "postgres"
)

type Article struct {
	Id                            uint16
	Username, Information, Ticket string
}

type Page struct {
	Article     []Article
	CurrentPage int
	TotalPage   int
}

var inforamations = []Article{}
var showTicket = Article{}

func main() {
	initDB()
	defer CloseDB()
	handleFunc()
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	information := r.FormValue("information")
	ticket := r.FormValue("ticket")

	if username == "" || information == "" || ticket == "" {
		fmt.Fprintf(w, "Not all data is filled!")
	} else {
		addInfoToDB := `insert into "info"("name", "information", "ticket") values($1, $2, $3)`
		iticket := fmt.Sprintf("#%s", ticket)
		_, err := db.Exec(addInfoToDB, username, information, iticket)
		CheckError(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func tickets(w http.ResponseWriter, r *http.Request) {
	pageNumberStr := r.URL.Query().Get("page")
	pageNumber, _ := strconv.Atoi(pageNumberStr)
	if pageNumber < 1 {
		pageNumber = 1
	}

	perPage := 10
	offset := (pageNumber - 1) * perPage

	query := "select * from info order by id limit $1 offset $2"
	rows, err := db.Query(query, perPage, offset)
	CheckError(err)

	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Username, &article.Information, &article.Ticket)
		CheckError(err)

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	countQuery := "select count(*) from info"
	var totalArticles int
	err = db.QueryRow(countQuery).Scan(&totalArticles)
	CheckError(err)

	totalPages := totalArticles/perPage + 1

	page := Page{
		Article:     articles,
		CurrentPage: pageNumber,
		TotalPage:   totalPages,
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

func show_ticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/ticket.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	rows, err := db.Query(fmt.Sprintf("select * from info where id = '%s'", vars["id"]))
	CheckError(err)

	showTicket = Article{}

	for rows.Next() {
		var post Article
		err := rows.Scan(&post.Id, &post.Username, &post.Information, &post.Ticket)
		if err != nil {
			continue
		}

		showTicket = post
	}

	t.ExecuteTemplate(w, "ticket", showTicket)
}

func minus(a, b int) int {
	return a - b
}

func plus(a, b int) int {
	return a + b
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create/", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/tickets/", tickets).Methods("GET")
	rtr.HandleFunc("/ticket/{id:[0-9]+}", show_ticket).Methods("GET")

	http.Handle("/", rtr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	CheckError(err)
}
