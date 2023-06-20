package rest

import (
	"net/http"
	"root/root/internal/services"

	"github.com/gorilla/mux"
)

func HandleFunction() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", services.Index).Methods("GET")
	rtr.HandleFunc("/create/", services.Create).Methods("GET")
	rtr.HandleFunc("/save_article", services.Save_article).Methods("POST")
	rtr.HandleFunc("/tickets/", services.Tickets).Methods("GET")
	rtr.HandleFunc("/ticket/{id:[0-9]+}", services.Show_ticket).Methods("GET")
	rtr.HandleFunc("/create_ticket", services.Create_ticket).Methods("GET")
	rtr.HandleFunc("/show_form", services.Show_form).Methods("GET")

	http.Handle("/", rtr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
