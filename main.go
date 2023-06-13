package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type Article struct {
	Id                            uint16
	Username, Information, Ticket string
}

var inforamations = []Article{}
var showTicket = Article{}

func main() {
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
	t, err := template.ParseFiles("templates/create.html")
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
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		addInfoToDB := `insert into "info"("name", "information", "ticket") values($1, $2, $3)`
		iticket := fmt.Sprintf("#%s", ticket)
		_, err = db.Exec(addInfoToDB, username, information, iticket)
		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func tickets(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/tickets.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from info")
	if err != nil {
		panic(err)
	}

	inforamations = []Article{}

	for rows.Next() {
		var post Article
		err := rows.Scan(&post.Id, &post.Username, &post.Information, &post.Ticket)
		if err != nil {
			continue
		}

		inforamations = append(inforamations, post)
	}

	t.ExecuteTemplate(w, "tickets", inforamations)
}

func show_ticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/ticket.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select * from info where id = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

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
