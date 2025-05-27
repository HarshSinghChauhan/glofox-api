package handlers

import (
	"encoding/json"
	"glofox/models"
	"glofox/store"
	"net/http"
	"time"
)

type bookingInput struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var input bookingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, input.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Validate that at least one class exists on this date
	found := false
	for _, class := range store.ListClasses() {
		for _, d := range class.Dates {
			if d.Equal(date) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		http.Error(w, "No class available on this date", http.StatusNotFound)
		return
	}

	booking := models.Booking{
		Name: input.Name,
		Date: date,
	}
	store.AddBooking(booking)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}
