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
			name: "Empty class name",
			payload: map[string]any{
				"name":       "",
				"start_date": "2025-05-28",
				"end_date":   "2025-05-30",
				"capacity":   10,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Zero capacity",
			payload: map[string]any{
				"name":       "Yoga",
				"start_date": "2025-05-28",
				"end_date":   "2025-05-30",
				"capacity":   0,
			},
			expectedCode: http.StatusBadRequest,
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
			name: "End date before start date",
			payload: map[string]any{
				"name":       "Yoga",
				"start_date": "2025-05-30",
				"end_date":   "2025-05-28",
				"capacity":   10,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid body format",
			payload:      nil, // simulate bad JSON
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.payload == nil {
				req = httptest.NewRequest(http.MethodPost, "/classes", bytes.NewBuffer([]byte("invalid-json")))
			} else {
				body, _ := json.Marshal(tt.payload)
				req = httptest.NewRequest(http.MethodPost, "/classes", bytes.NewReader(body))
			}
			w := httptest.NewRecorder()

			CreateClassHandler(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
			}
		})
	}
}
