package dto

type ErrorAPIResponse struct {
	StatusCode string `json:"code"`
	Error      string `json:"error"`
}

type BookingSuccessAPIResponse struct {
	StatusCode string `json:"code"`
	Message    string `json:"message"`
}

type CreateSuccessAPIResponse struct {
	StatusCode string   `json:"code"`
	Message    string   `json:"message"`
	DateRange  []string `json:"date"`
}
