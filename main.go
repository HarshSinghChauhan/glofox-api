// main.go
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

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
