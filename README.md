# Go Lang Marathon

Welcome to the **Go Lang Marathon** workspace! This repository contains a curated collection of Go projects demonstrating a wide range of backend architectures, concurrency patterns, database persistence, caching/distributed locking with Redis, microservices skeletons, and system utilities.

---

## 📂 Project Directory Overview

Below is the directory map of the projects in this workspace, sorted by completion status.

| Project Directory | Status | Technologies Used | Description |
|---|---|---|---|
| [cinema-ticket-booking](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking) | Completed 🟢 | Go, Redis, Gorilla Mux, HTML5/CSS3 | Real-time seat reservation with lease-based locks. |
| [docker-cleanup-daemon](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/docker-cleanup-daemon) | Completed 🟢 | Go, Docker SDK, Cron, YAML | Periodically prunes stopped containers, dangling images, and volumes. |
| [go-crud-app](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go-crud-app) | Completed 🟢 | Go, Gorilla Mux, In-Memory Slice | Simple in-memory movie CRUD REST API. |
| [bookstore-go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/bookstore-go) | Partial 🟡 | Go, GORM, MySQL, Gorilla Mux | Bookstore database REST API (routing and connection config done). |
| [Microservice-in-go](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/Microservice-in-go) | Skeleton 🔴 | Go, GraphQL, Docker Compose | Skeleton for a GraphQL microservices gate. |
| [go server](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go%20server) | Broken 🟠 | Go, net/http, HTML | Standard fileserver template, missing formHandler compile symbol. |

---

## 🚀 Highlighted Completed Projects

### 1. Cinema Seat Booking System
A high-concurrency, real-time ticket booking application that prevents double-booking using **Redis-backed lease locks**. It features a modern dark-mode frontend that polls the server for status updates.
* **Location**: [cinema-ticket-booking](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking)
* **Read the full guide**: [cinema-ticket-booking/README.md](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/cinema-ticket-booking/README.md)

### 2. Docker Cleanup Daemon
A lightweight background service that interacts with the Docker Go SDK to prune stopped containers, dangling images, and unused volumes periodically using a Cron scheduler.
* **Location**: [docker-cleanup-daemon](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/docker-cleanup-daemon)
* **Read the full guide**: [docker-cleanup-daemon/README.md](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/docker-cleanup-daemon/README.md)

### 3. Movie CRUD API
A simple, robust, in-memory REST API demonstrating basic CRUD operations using Gorilla Mux.
* **Location**: [go-crud-app](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go-crud-app)
* **Read the full guide**: [go-crud-app/README.md](file:///C:/Users/navde/Desktop/Data%20Engineering%20Project%27/Go%20Lang%20Marathon/go-crud-app/README.md)
