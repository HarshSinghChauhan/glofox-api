# Glofox Class Booking API

This project is a simple RESTful API built with GoLang to manage fitness studio classes and class bookings. It maintains all data in-memory and does not require authentication.

## Features

### 1. Create Classes

- Endpoint: `POST /classes`
- Allows studio owners to create recurring classes between two dates.
- Each date in the range creates a unique class instance with specified capacity.

### 2. Book a Class

- Endpoint: `POST /bookings`
- Allows members to book a class on a specified date by providing their name and desired date.

## Initialize dependencies

- go mod tidy

## Run the server

- go run main.go

## Project structure

```
glofox/
├── main.go                  # Application entry point
├── handlers/
│   ├── class_handler.go     # Handler for class creation
│   └── booking_handler.go   # Handler for bookings
├── models/
│   └── models.go            # Data models (Class, Booking)
├── store/
│   └── memory_store.go      # In-memory concurrent-safe data store
├── internal/
│   ├── const/               # Constants for log messages and error codes
│   │   └── const.go
│   └── dto/                 # Response DTOs
│       └── response.go
├── handlers/
│   └── class_handler_test.go    # Unit tests for class handler
│   └── booking_handler_test.go  # Unit tests for booking handler
├── go.mod
└── README.md

```

## API Documentation

## 1. Create Classes

### POST /classes

- Request Body: 
```
{
  "name": "Yoga",
  "start_date": "2025-06-01",
  "end_date": "2025-06-05",
  "capacity": 1
}
```

### API response

- Status: 201 Created
- Response: 
```
{
    "code": "0",
    "message": "Classes created successfully for all dates",
    "date": [
        "2025-06-01",
        "2025-06-02",
        "2025-06-03",
        "2025-06-04",
        "2025-06-05"
    ]
}
```

## 2. Book a Class

### POST /bookings

- Request Body 
```
{
  "name": "Alice",
  "date": "2025-06-03"
}
```

### API response

- Status: 201 Created
- Response
```
{
    "code": "0",
    "message": "Booking created successfully"
}
```

## Testing
### Covering create class

```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCreateClassHandler$ glofox/handlers

=== RUN   TestCreateClassHandler
=== RUN   TestCreateClassHandler/Valid_class_creation
--- PASS: TestCreateClassHandler/Valid_class_creation (0.00s)
=== RUN   TestCreateClassHandler/Invalid_date_format
--- PASS: TestCreateClassHandler/Invalid_date_format (0.00s)
=== RUN   TestCreateClassHandler/Empty_name
--- PASS: TestCreateClassHandler/Empty_name (0.00s)
=== RUN   TestCreateClassHandler/Invalid_date_range_(end_before_start)
--- PASS: TestCreateClassHandler/Invalid_date_range_(end_before_start) (0.00s)
--- PASS: TestCreateClassHandler (0.00s)
PASS
ok      glofox/handlers 0.425s
```

### Covering booking class

```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCreateBookingHandler$ glofox/handlers

=== RUN   TestCreateBookingHandler
=== RUN   TestCreateBookingHandler/Valid_booking
--- PASS: TestCreateBookingHandler/Valid_booking (0.00s)
=== RUN   TestCreateBookingHandler/Missing_name
--- PASS: TestCreateBookingHandler/Missing_name (0.00s)
=== RUN   TestCreateBookingHandler/Invalid_date_format
--- PASS: TestCreateBookingHandler/Invalid_date_format (0.00s)
=== RUN   TestCreateBookingHandler/Class_not_found
--- PASS: TestCreateBookingHandler/Class_not_found (0.00s)
--- PASS: TestCreateBookingHandler (0.00s)
PASS
ok      glofox/handlers 0.244s
```

## Design and Architecture

- Modular structure: Separate packages for handlers, models, and store to keep code organized and maintainable.
- In-memory storage: Uses thread-safe maps with sync.Mutex for concurrent access safety.
- RESTful design: Proper HTTP methods and status codes used.
- Validation: Input is validated with appropriate error messages.
- No authentication: APIs are open as per requirements.
- No persistence: Data lost on server restart, suitable for prototype or demo.

## Performance Considerations

- In-memory maps provide O(1) lookups for classes and bookings.
- Mutex locking ensures safe concurrent access without race conditions.
- Suitable for small to medium load scenarios (e.g., prototype or initial implementation).

## Testing

- Manual testing can be done with curl, Postman, or any REST client.

- Example commands:

## Create class: 
```
curl --location 'http://localhost:8080/classes' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Yoga",
  "start_date": "2025-06-01",
  "end_date": "2025-06-05",
  "capacity": 1
}'
```

## Book a class:
```
curl --location 'http://localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Alice",
  "date": "2025-06-03"
}'
```

## Known limitations 

- Currently no database persistence; all data is lost on restart.
- No authentication or authorization.
- Overbooking is allowed without limits.
- No pagination or listing APIs.
- No cancellation or modification endpoints.

## Future enhancements

- Add persistent storage (e.g., PostgreSQL, Redis).
- Add authentication.
- Implement booking limits and availability checks.
- Add API documentation (Swagger/OpenAPI).
- Add automated tests and CI integration.