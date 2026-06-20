# Cinema Ticket & Seat Booking System

A high-performance, real-time cinema seat reservation system written in Go, featuring a lease-based locking mechanism powered by Redis to handle concurrent bookings without double-booking.

---

## ⚡ Concurrency & Lock Architecture

When multiple users attempt to reserve the same seat simultaneously, the application guarantees that only one user acquires the hold. 

### 1. Lease-Based Locks (TTL)
* Seat holds are temporary and default to a **2-minute Time-To-Live (TTL)**.
* If a user holds a seat but fails to confirm the purchase within 2 minutes, Redis automatically deletes the hold, making the seat available to others.

### 2. Atomic Lease Acquisition (NX Mode)
* In [redis_store.go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/internal/booking/redis_store.go), the [hold](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/internal/booking/redis_store.go#L68) function utilizes the Redis `SET` command with the `NX` flag:
  ```go
  res := s.rdb.SetArgs(ctx, key, val, redis.SetArgs{
      Mode: "NX", // Set if Not Exists
      TTL:  defaultHoldTTL,
  })
  ```
* This operation is guaranteed to be atomic by Redis. If the key already exists, the hold fails, and the server returns `ErrSeatAlreadyBooked` ([domain.go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/internal/booking/domain.go#L10)).

### 3. Redis Key Schema
* **Seat Key**: `seat:{movieID}:{seatID}` &rarr; Maps to a serialized JSON representing the [Booking](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/internal/booking/domain.go#L14) struct.
  - *Held State*: Key has a TTL (e.g. 2 minutes) and Status is `"held"`.
  - *Confirmed State*: Key has no TTL (persisted) and Status is `"confirmed"`.
* **Session Key**: `session:{sessionID}` &rarr; Maps to the corresponding seat key for reverse lookup during confirmation or release.

---

## 🛠️ API Reference

The server runs on port `8080`. Routes are configured using standard library multiplexing in [main.go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/cmd/main.go).

| Method | Endpoint | Description | Request Body |
|---|---|---|---|
| **GET** | `/movies` | Returns list of configured movies (e.g., Inception, Dune) | None |
| **GET** | `/movies/{movieID}/seats` | Returns current status of all seats for a movie | None |
| **POST** | `/movies/{movieID}/seats/{seatID}/hold` | Acquires a temporary hold lease on a seat | `{"user_id": "string"}` |
| **PUT** | `/sessions/{sessionID}/confirm` | Confirms the seat purchase, persisting the booking | `{"user_id": "string"}` |
| **DELETE** | `/sessions/{sessionID}` | Explicitly cancels a hold and releases the seat | `{"user_id": "string"}` |

---

## 🎨 Interactive Frontend

The client-side single-page app is built using vanilla HTML5, CSS3, and JavaScript, located in [index.html](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/static/index.html).
* **Real-time Synchronization**: Polls the `/movies/{movieID}/seats` endpoint every 2 seconds to capture seat state changes from other clients.
* **Countdown Timer**: Automatically displays a visual timer once a seat is held, changing to red (urgent) in the final 60 seconds.
* **Visual Legend**:
  - **Dark Gray**: Available seat.
  - **Yellow**: Held by you.
  - **Orange**: Held by another user.
  - **Red**: Confirmed/Booked seat.

---

## 🏃 Getting Started

### Prerequisites
* Go (v1.22+)
* Docker & Docker Compose

### Step 1: Start Redis Caching Store
Deploy Redis and Redis Commander using the provided [docker-compose.yaml](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/docker-compose.yaml):
```bash
docker-compose up -d
```
* **Redis Store**: `localhost:6379`
* **Redis Commander GUI**: [http://localhost:8081](http://localhost:8081)

### Step 2: Build & Start the Go Server
Run the main file from the project directory:
```bash
go run cmd/main.go
```
* The web portal is hosted at [http://localhost:8080](http://localhost:8080).
