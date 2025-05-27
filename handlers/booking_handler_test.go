package handlers

import (
	"bytes"
	"encoding/json"
	"glofox/models"
	"glofox/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBookingHandler(t *testing.T) {
	// Setup a class on a specific date
	store.Classes = make(map[string]models.Class)
	store.Bookings = []models.Booking{}
	date := "2025-05-29"
	store.Classes[date] = models.Class{
		Name:     "Yoga",
		Capacity: 10,
	}

	tests := []struct {
		name         string
		payload      map[string]any
		expectedCode int
	}{
		{
			name: "Valid booking",
			payload: map[string]any{
				"name": "Alice",
				"date": date,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "Missing name",
			payload: map[string]any{
				"date": date,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Invalid date format",
			payload: map[string]any{
				"name": "Bob",
				"date": "29-05-2025",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Class not found",
			payload: map[string]any{
				"name": "Charlie",
				"date": "2025-05-30",
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/booking", bytes.NewReader(payload))
			w := httptest.NewRecorder()

			CreateBookingHandler(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
			}
		})
	}
}
