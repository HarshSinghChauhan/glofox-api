package handlers

import (
	"encoding/json"
	cc "glofox/internal/const"
	"glofox/internal/dto"
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
	w.Header().Set("Content-Type", "application/json")
	var input classInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrInvalidBodyCode,
			Error:      cc.ErrInvalidBody,
		})
		return
	}

	start, err1 := time.Parse("2006-01-02", input.StartDate)
	end, err2 := time.Parse("2006-01-02", input.EndDate)

	if err1 != nil || err2 != nil || end.Before(start) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrInvalidDateRangeCode,
			Error:      cc.ErrInvalidDateRange,
		})
		return
	}

	if input.Name == "" || input.Capacity <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrMissingClassNameCode,
			Error:      cc.ErrMissingClassName,
		})
		return
	}

	dateStrings := GenerateDateStrings(start, end)
	classByDate := make(map[string]models.Class)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("2006-01-02")
		classByDate[dateStr] = models.Class{
			Name:      input.Name,
			StartDate: start,
			EndDate:   end,
			Capacity:  input.Capacity,
			Dates:     []time.Time{d},
		}
	}

	store.AddClassesByDate(classByDate)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.CreateSuccessAPIResponse{
		StatusCode: cc.ApiSuccessCode,
		Message:    cc.LogClassCreated,
		DateRange:  dateStrings,
	})
}

func GenerateDateStrings(start, end time.Time) []string {
	dates := []string{}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}
	return dates
}
