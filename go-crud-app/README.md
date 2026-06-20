# Movie CRUD REST API

A simple, lightweight RESTful API built in Go for managing a collection of movies. This project demonstrates in-memory CRUD operations using Gorilla Mux.

---

## 🛠️ Tech Stack & Architecture

* **Routing**: [Gorilla Mux](https://github.com/gorilla/mux) for robust URL path parameters and HTTP method matching.
* **Storage**: In-memory global slice `movies` ([main.go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go-crud-app/main.go#L32)) of `Movie` structs. 
* **Data Persistence**: **None**. Data is lost whenever the server process terminates.

---

## 📋 Data Structures

The models are defined in [main.go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go-crud-app/main.go#L15-L27):

### Movie Struct
```go
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
```

### Director Struct
```go
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
```

---

## 🚀 API Endpoints

The API server runs on port `8000`.

| Method | Endpoint | Description | Sample JSON Request Body |
|---|---|---|---|
| **GET** | `/movies` | Returns the entire list of movies | None |
| **GET** | `/movies/{id}` | Returns a single movie matching the given ID | None |
| **POST** | `/movies` | Creates a new movie with a random auto-generated ID | `{"isbn":"112233", "title":"Inception", "director":{"firstname":"Christopher", "lastname":"Nolan"}}` |
| **PUT** | `/movies/{id}` | Replaces/updates the movie matching the given ID | `{"isbn":"112233", "title":"Interstellar", "director":{"firstname":"Christopher", "lastname":"Nolan"}}` |
| **DELETE** | `/movies/{id}` | Deletes the movie matching the given ID from the slice | None |

---

## 🏃 Getting Started

### Prerequisites
* Go (v1.22+)

### Step 1: Run the Server
From the `go-crud-app` directory, run:
```bash
go run main.go
```
The server will start listening at [http://localhost:8000](http://localhost:8000).

### Step 2: Test Endpoints
You can use `curl` or Postman to interact with the API:
```bash
# Retrieve all movies
curl http://localhost:8000/movies

# Add a movie
curl -X POST -H "Content-Type: application/json" -d '{"isbn":"987654", "title":"The Dark Knight", "director":{"firstname":"Christopher", "lastname":"Nolan"}}' http://localhost:8000/movies
```
