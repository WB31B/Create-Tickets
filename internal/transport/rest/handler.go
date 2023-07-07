package rest

import (
	"net/http"
	"root/root/internal/services"

	"github.com/gorilla/mux"
)

func HandleFunction() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", services.Index).Methods("GET")
	rtr.HandleFunc("/create_ticket", services.CreateTicket).Methods("GET")
	rtr.HandleFunc("/create_note", services.CreateNote).Methods("GET")
	rtr.HandleFunc("/save_ticket", services.SaveTicket).Methods("POST")
	rtr.HandleFunc("/save_note", services.SaveNote).Methods("POST")
	rtr.HandleFunc("/show_tickets", services.ShowTickets).Methods("GET")
	rtr.HandleFunc("/show_notes", services.ShowNotes).Methods("GET")
	rtr.HandleFunc("/ticket/{id:[0-9]+}", services.ShowTicket).Methods("GET")
	rtr.HandleFunc("/note/{id:[0-9]+}", services.ShowNote).Methods("GET")

	http.Handle("/", rtr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
