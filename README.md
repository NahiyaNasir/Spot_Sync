
Spot-sync & EV Charging Reservation

A centralized platform for busy airports and shopping malls to manage parking zones, with a special focus on reserving limited EV charging spots. The system ensures fair allocation of parking spaces, prevents overbooking, and provides real-time parking availability.


---

# 🌐 Live URL


**Backend API:** `https://your-backend-url.com`

---

# ✨ Features

### Authentication

* User Registration & Login
* JWT Authentication
* Protected Routes
* Role-based Authorization (Admin/User)

### Parking Zone Management

* View available parking zones
* Real-time available parking spots calculation
* Admin can create, update, and delete parking zones

### Reservation Management

* Create parking reservations
* Prevent overbooking using database transactions
* Row-level locking (`FOR UPDATE`) to avoid race conditions
* Cancel reservations
* View reservation history

### EV Charging Spots

* Dedicated EV charging reservations
* Limited charging slot management
* Active reservation validation

### API Features

* RESTful API
* Request validation
* Standardized JSON responses
* Pagination
* Error handling

---

# 🛠 Tech Stack

## Backend

* Go (Golang)
* Gin Framework
* GORM
* PostgreSQL
* JWT Authentication

## Database

* PostgreSQL

## Tools

* Postman
* Git & GitHub

---

# 🏗 Architecture

The project follows a layered architecture to separate concerns and keep the codebase maintainable.

```
                Client
                   │
                   ▼
             HTTP Request
                   │
                   ▼
             Router (Gin)
                   │
                   ▼
          Middleware Layer
      (JWT, Validation, Logger)
                   │
                   ▼
          Controller / Handler
                   │
                   ▼
            Service Layer
        (Business Logic)
                   │
                   ▼
          Repository Layer
        (Database Queries)
                   │
                   ▼
            PostgreSQL
```

## Layer Responsibilities

### Router

Defines all API routes and maps them to handlers.

### Middleware

* JWT Authentication
* Authorization
* Request validation
* Logging

### Controller / Handler

Receives HTTP requests, validates input, and returns HTTP responses.

### Service Layer

Contains the business logic such as:

* Reservation validation
* EV charging slot management
* Preventing double booking
* Transaction handling

### Repository Layer

Handles all database operations using GORM.

### Database

Stores users, parking zones, and reservations.

---

# 🚀 Getting Started

## Prerequisites

* Go 1.24+
* PostgreSQL
* Git

---

## Clone Repository

```bash
git clone https://github.com/your-username/spot_sync.git

cd spot_sync
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Configure Environment Variables

Create a `.env` file in the project root.

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=smart_parking
DB_SSLMODE=disable

JWT_SECRET=your_jwt_secret
JWT_EXPIRE=24h
```

---

## Run the Project

```bash
go run main.go
```

or

```bash
go run cmd/main.go
```

---

# 📌 API Endpoints

## Authentication

| Method | Endpoint                | Description   |
| ------ | ----------------------- | ------------- |
| POST   | `/api/v1/auth/register` | Register user |
| POST   | `/api/v1/auth/login`    | Login user    |

---

## Parking Zones

| Method | Endpoint            | Access |
| ------ | ------------------- | ------ |
| GET    | `/api/v1/zones`     | Public |
| GET    | `/api/v1/zones/:id` | Public |
| POST   | `/api/v1/zones`     | Admin  |


---

## Reservations

| Method | Endpoint                          | Access |
| ------ | --------------------------------- | ------ |
| POST   | `/api/v1/reservations`            | driver   |
| GET    | `/api/v1/reservations`            | driver   |
| GET    | `/api/v1/reservations/:id`        | driver   |
| PUT    | `/api/v1/reservations/:id/` | driver   |

---

# 🔒 Concurrency Handling

To prevent overbooking of EV charging spots and parking spaces, the application uses:

* Database Transactions
* Row-Level Locking (`FOR UPDATE`)
* GORM Transactions
* Atomic Reservation Creation

These mechanisms ensure that multiple users cannot reserve the same parking spot simultaneously.

---

# 📂 Project Structure

```text
smart-parking/
│
├── cmd/
├── config/
├── controllers/
├── middleware/
├── models/
├── repositories/
├── routes/
├── services/
├── utils/
├── .env
├── go.mod
├── go.sum
└── main.go
```

---
