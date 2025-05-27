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

func TestCreateClassHandler(t *testing.T) {
	store.Classes = make(map[string]models.Class)

	tests := []struct {
		name         string
		payload      map[string]any
		expectedCode int
	}{
		{
			name: "Valid class creation",
			payload: map[string]any{
				"name":       "Yoga",
				"start_date": "2025-05-28",
				"end_date":   "2025-05-30",
				"capacity":   10,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "Invalid date format",
			payload: map[string]any{
				"name":       "Yoga",
				"start_date": "28-05-2025",
				"end_date":   "2025-05-30",
				"capacity":   10,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Empty name",
			payload: map[string]any{
				"name":       "",
				"start_date": "2025-05-28",
				"end_date":   "2025-05-30",
				"capacity":   10,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Invalid date range (end before start)",
			payload: map[string]any{
				"name":       "Yoga",
				"start_date": "2025-05-30",
				"end_date":   "2025-05-28",
				"capacity":   10,
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/class", bytes.NewReader(payload))
			w := httptest.NewRecorder()

			CreateClassHandler(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
			}
		})
	}
}
