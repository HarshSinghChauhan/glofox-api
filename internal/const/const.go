package cc

// Api success
const (
	ApiSuccessCode = "0"
)

// Log Messages
const (
	LogClassCreated   = "Classes created successfully for all dates"
	LogBookingSuccess = "Booking created successfully"
)

// Error Messages
const (
	ErrInvalidBody       = "Invalid request body"
	ErrMissingName       = "name is required"
	ErrInvalidDate       = "valid date is required"
	ErrClassNotFound     = "class not found for the given date"
	ErrInvalidDateRange  = "Invalid start or end date"
	ErrMissingClassName  = "Name and positive capacity are required"
	ErrInvalidDateFormat = "Date format should be YYYY-MM-DD"
	ErrPastDateBooking   = "cannot book a class in the past"
	ErrDuplicateBooking  = "user already booked for this date"
	ErrClassFull         = "class is full for selected date"
	ErrInvalidCapacity   = "capacity must be greater than zero"
	ErrPastDateClass     = "cannot create class for past dates"
	ErrDuplicateClass    = "class already exists for date"
)

// Error code
const (
	ErrClassNotFoundCode     = "1001"
	ErrInvalidBodyCode       = "1002"
	ErrMissingNameCode       = "1003"
	ErrInvalidDateRangeCode  = "1004"
	ErrMissingClassNameCode  = "1005"
	ErrInvalidDateCode       = "1006"
	ErrInvalidDateFormatCode = "1007"
	ErrClassFullCode         = "1008"
	ErrDuplicateBookingCode  = "1009"
	ErrPastDateBookingCode   = "1010"
	ErrPastDateClassCode     = "1011"
	ErrInvalidCapacityCode   = "1012"
	ErrDuplicateClassCode    = "1013"
)
