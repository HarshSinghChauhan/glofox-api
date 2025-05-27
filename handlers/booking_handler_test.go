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
			name: "Missing date",
			payload: map[string]any{
				"name": "Bob",
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
				"date": "2025-06-01",
			},
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Invalid body format",
			payload:      nil, // Will simulate by passing invalid JSON
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.payload == nil {
				req = httptest.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer([]byte("invalid-json")))
			} else {
				body, _ := json.Marshal(tt.payload)
				req = httptest.NewRequest(http.MethodPost, "/booking", bytes.NewReader(body))
			}
			w := httptest.NewRecorder()

			CreateBookingHandler(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
			}
		})
	}
}
