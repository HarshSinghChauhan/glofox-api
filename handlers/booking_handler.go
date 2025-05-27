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

// CreateBookingHandler handles booking requests
func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var bookingRequest models.Booking
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&bookingRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrInvalidBodyCode,
			Error:      cc.ErrInvalidBody,
		})
		return
	}

	// Validate name
	if bookingRequest.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrMissingNameCode,
			Error:      cc.ErrMissingName,
		})
		return
	}

	// Validate date
	if bookingRequest.Date == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrInvalidDateCode,
			Error:      cc.ErrInvalidDate,
		})
		return
	}

	bookingDate, err := time.Parse("2006-01-02", bookingRequest.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrInvalidDateFormatCode,
			Error:      cc.ErrInvalidDateFormat,
		})
		return
	}

	dateStr := bookingDate.Format("2006-01-02")

	// Lock and check for class existence
	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	if _, exists := store.Classes[dateStr]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorAPIResponse{
			StatusCode: cc.ErrClassNotFoundCode,
			Error:      cc.ErrClassNotFound,
		})
		return
	}

	// Save booking
	store.Bookings = append(store.Bookings, bookingRequest)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.BookingSuccessAPIResponse{
		StatusCode: cc.ApiSuccessCode,
		Message:    cc.LogBookingSuccess,
	})
}
