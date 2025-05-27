package main

import (
	"glofox/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/classes", handlers.CreateClassHandler).Methods("POST")
	r.HandleFunc("/bookings", handlers.CreateBookingHandler).Methods("POST")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		log.Println("Registered Route:", path)
		return nil
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
