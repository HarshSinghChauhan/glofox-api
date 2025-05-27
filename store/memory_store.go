package store

import (
	"glofox/models"
	"sync"
)

var (
	Classes  = make(map[string]models.Class) // Key: date (YYYY-MM-DD)
	Bookings = make([]models.Booking, 0)
	Mutex    sync.Mutex
)

// AddClassesByDate adds multiple classes, one per day between start and end
func AddClassesByDate(datedClasses map[string]models.Class) {
	Mutex.Lock()
	defer Mutex.Unlock()
	for date, class := range datedClasses {
		Classes[date] = class
	}
}

// GetClassByDate returns class info if available on that date
func GetClassByDate(date string) (models.Class, bool) {
	Mutex.Lock()
	defer Mutex.Unlock()
	c, exists := Classes[date]
	return c, exists
}

// ListClasses returns all scheduled classes
func ListClasses() []models.Class {
	Mutex.Lock()
	defer Mutex.Unlock()
	result := make([]models.Class, 0, len(Classes))
	for _, v := range Classes {
		result = append(result, v)
	}
	return result
}

// AddBooking stores a booking in memory
func AddBooking(b models.Booking) {
	Mutex.Lock()
	defer Mutex.Unlock()
	Bookings = append(Bookings, b)
}
