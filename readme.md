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

glofox/
├── main.go                  # Application entry point
├── handlers/
│   ├── class_handler.go     # Handler for class creation
│   └── booking_handler.go   # Handler for bookings
├── models/
│   └── models.go            # Data models (Class, Booking)
├── store/
│   └── memory_store.go      # In-memory concurrent-safe data store
├── go.mod                   # Go module file
└── README.md                # This file

## API Documentation

## 1. Create Classes

### POST /classes

- Request Body: 
{
  "name": "Pilates",
  "start_date": "2025-06-01T00:00:00Z",
  "end_date": "2025-06-03T00:00:00Z",
  "capacity": 10
}

### API response

- Status: 201 Created
- Body 
{
  "message": "Classes created successfully"
}


## 2. Book a Class

### POST /bookings

- Request Body 
{
  "name": "Alice",
  "date": "2025-06-02T00:00:00Z"
}

### API response

- Status: 201 Created
- Body
{
  "message": "Booking created successfully"
}


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
- curl -X POST http://localhost:8080/classes \
  -H "Content-Type: application/json" \
  -d '{"name":"Yoga","start_date":"2025-06-01T00:00:00Z","end_date":"2025-06-03T00:00:00Z","capacity":10}'

## Book a class:
- curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"name":"Bob","date":"2025-06-02T00:00:00Z"}'


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