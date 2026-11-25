# Cinema Ticket Booking

Mini project built with Go and Redis.

## Features

- Display available movies and seats
- Hold a seat for a limited time
- Confirm or release a seat reservation
- Automatically detect held and confirmed seats
- Prevent the same seat from being booked simultaneously
- Use Redis for in-memory booking storage
- Provide a simple REST API and browser-based interface

## Tech Stack

- Go
- Redis
- Go `net/http` package
- HTML, CSS, and JavaScript
- Docker Compose

## Project Structure

```text
cmd/                  Application entry point
internal/booking/     Booking domain, service, handlers, and stores
internal/adapters/    Redis client adapter
internal/utils/       HTTP utility helpers
static/               Frontend web interface
```

## Booking Flow

1. The user selects a movie.
2. The user holds an available seat.
3. Redis stores the seat and session with a temporary expiration time.
4. The user can confirm or release the booking.
5. Confirmed seats remain in Redis and are shown as unavailable.

## How to Run

### Prerequisites

- Go
- Docker and Docker Compose

### Start Redis

```bash
docker compose up -d redis
```

Redis will run on `localhost:6379`.

### Start the Application

```bash
go run ./cmd
```

Open `http://localhost:8080` in your browser.

## Main API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/movies` | Get available movies |
| `GET` | `/movies/{movieID}/seats` | Get seat status |
| `POST` | `/movies/{movieID}/seats/{seatID}/hold` | Temporarily hold a seat |
| `PUT` | `/sessions/{sessionID}/confirm` | Confirm a held seat |
| `DELETE` | `/sessions/{sessionID}` | Release a held seat |

