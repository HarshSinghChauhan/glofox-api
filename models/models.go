package models

import "time"

type Class struct {
	Name      string      `json:"name"`
	StartDate time.Time   `json:"start_date"`
	EndDate   time.Time   `json:"end_date"`
	Capacity  int         `json:"capacity"`
	Dates     []time.Time `json:"dates"`
}

type Booking struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
