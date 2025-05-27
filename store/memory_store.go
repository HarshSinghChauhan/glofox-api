package store

import (
	"glofox/models"
	"sync"
	_ "time"
)

var (
	classes  = make(map[string]models.Class)
	bookings = make([]models.Booking, 0)
	mu       sync.Mutex
)

func AddClass(c models.Class) {
	mu.Lock()
	defer mu.Unlock()
	classes[c.Name] = c
}

func GetClass(name string) (models.Class, bool) {
	mu.Lock()
	defer mu.Unlock()
	c, exists := classes[name]
	return c, exists
}

func ListClasses() []models.Class {
	mu.Lock()
	defer mu.Unlock()
	result := make([]models.Class, 0, len(classes))
	for _, v := range classes {
		result = append(result, v)
	}
	return result
}

func AddBooking(b models.Booking) {
	mu.Lock()
	defer mu.Unlock()
	bookings = append(bookings, b)
}
