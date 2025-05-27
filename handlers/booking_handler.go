package handlers

import (
	"encoding/json"
	cc "glofox/internal/const"
	"glofox/internal/dto"
	"glofox/models"
	"glofox/store"
	"net/http"
	"strings"
	"time"
)

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bookingRequest models.Booking
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		writeError(w, http.StatusBadRequest, cc.ErrInvalidBodyCode, cc.ErrInvalidBody)
		return
	}

	bookingRequest.Name = strings.TrimSpace(bookingRequest.Name)
	if bookingRequest.Name == "" {
		writeError(w, http.StatusBadRequest, cc.ErrMissingNameCode, cc.ErrMissingName)
		return
	}

	if bookingRequest.Date == "" {
		writeError(w, http.StatusBadRequest, cc.ErrInvalidDateCode, cc.ErrInvalidDate)
		return
	}

	bookingDate, err := time.Parse("2006-01-02", bookingRequest.Date)
	if err != nil {
		writeError(w, http.StatusBadRequest, cc.ErrInvalidDateFormatCode, cc.ErrInvalidDateFormat)
		return
	}
	dateStr := bookingDate.Format("2006-01-02")

	// Check if class exists
	if !classExists(dateStr) {
		writeError(w, http.StatusNotFound, cc.ErrClassNotFoundCode, cc.ErrClassNotFound)
		return
	}

	// Save booking in a goroutine
	go storeBooking(bookingRequest)

	writeSuccess(w, http.StatusCreated, dto.BookingSuccessAPIResponse{
		StatusCode: cc.ApiSuccessCode,
		Message:    cc.LogBookingSuccess,
	})
}

func classExists(date string) bool {
	store.Mutex.RLock()
	defer store.Mutex.RUnlock()
	_, exists := store.Classes[date]
	return exists
}

func storeBooking(booking models.Booking) {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()
	store.Bookings = append(store.Bookings, booking)
}
