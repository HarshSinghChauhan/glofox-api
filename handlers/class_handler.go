package handlers

import (
	"encoding/json"
	"glofox/models"
	"glofox/store"
	"net/http"
	"time"
)

type classInput struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var input classInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	start, err1 := time.Parse(time.RFC3339, input.StartDate)
	end, err2 := time.Parse(time.RFC3339, input.EndDate)
	if err1 != nil || err2 != nil || end.Before(start) {
		http.Error(w, "Invalid start or end date", http.StatusBadRequest)
		return
	}
	if input.Name == "" || input.Capacity <= 0 {
		http.Error(w, "Invalid name or capacity", http.StatusBadRequest)
		return
	}

	// Generate list of dates for the class
	dates := []time.Time{}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}

	class := models.Class{
		Name:      input.Name,
		StartDate: start,
		EndDate:   end,
		Capacity:  input.Capacity,
		Dates:     dates,
	}

	store.AddClass(class)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}
