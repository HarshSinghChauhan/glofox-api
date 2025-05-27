package handlers

import (
	"encoding/json"
	cc "glofox/internal/const"
	"glofox/internal/dto"
	"glofox/models"
	"glofox/store"
	"log"
	"net/http"
	"strings"
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
		log.Printf("Failed to decode class creation request: %v", err)
		writeError(w, http.StatusBadRequest, cc.ErrInvalidBodyCode, cc.ErrInvalidBody)
		return
	}

	// Validate & sanitize
	if err := validateClassInput(input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error, err.StatusCode)
		return
	}

	start, _ := time.Parse("2006-01-02", input.StartDate)
	end, _ := time.Parse("2006-01-02", input.EndDate)

	// Check for duplicate classes
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if classExists(d.Format("2006-01-02")) {
			writeError(w, http.StatusConflict, cc.ErrDuplicateClassCode, cc.ErrDuplicateClass+": "+d.Format("2006-01-02"))
			return
		}
	}

	log.Printf("Class created successfully: %s from %s to %s", input.Name, input.StartDate, input.EndDate)

	// Prepare and add classes in goroutine
	go addClassByDate(input, start, end)

	writeSuccess(w, http.StatusCreated, dto.CreateSuccessAPIResponse{
		StatusCode: cc.ApiSuccessCode,
		Message:    cc.LogClassCreated,
		DateRange:  generateDateStrings(start, end),
	})
}

func validateClassInput(input classInput) *dto.ErrorAPIResponse {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return &dto.ErrorAPIResponse{StatusCode: cc.ErrMissingClassNameCode, Error: cc.ErrMissingClassName}
	}
	if input.Capacity <= 0 {
		return &dto.ErrorAPIResponse{StatusCode: cc.ErrInvalidCapacityCode, Error: cc.ErrInvalidCapacity}
	}
	start, err1 := time.Parse("2006-01-02", input.StartDate)
	end, err2 := time.Parse("2006-01-02", input.EndDate)
	if err1 != nil || err2 != nil {
		return &dto.ErrorAPIResponse{StatusCode: cc.ErrInvalidDateFormatCode, Error: cc.ErrInvalidDateFormat}
	}
	if end.Before(start) {
		return &dto.ErrorAPIResponse{StatusCode: cc.ErrInvalidDateRangeCode, Error: cc.ErrInvalidDateRange}
	}
	if start.Before(time.Now().Truncate(24 * time.Hour)) {
		return &dto.ErrorAPIResponse{StatusCode: cc.ErrPastDateClassCode, Error: cc.ErrPastDateClass}
	}
	return nil
}

func addClassByDate(input classInput, start, end time.Time) {
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
}
