package handlers

import (
	"encoding/json"
	"glofox/internal/dto"
	"net/http"
	"time"
)

func writeError(w http.ResponseWriter, status int, code string, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
		StatusCode: code,
		Error:      message,
	})
}

func writeSuccess(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func generateDateStrings(start, end time.Time) []string {
	var dates []string
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}
	return dates
}
