# 🎟️ Ticket Booking Backend (Golang)

This is the backend service for the [Ticket Booking App (React Native + Expo)](https://github.com/pyroblazer/ticket-booking-mobile-react-native-expo), built using **Golang**, **Fiber**, **GORM**, and **JWT** authentication. It provides RESTful APIs for event listings, ticket bookings, and user management.

---

## 🧩 Features

* ✅ JWT-based authentication (Login & Register)
* 🎫 Create, update, delete, and view events
* 🗓️ Book and manage tickets linked to users and events
* 🧾 REST API designed for a mobile frontend
* 🔐 Role-protected routes with middleware
* 🌐 CORS configuration to support frontend access

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/pyroblazer/ticket-booking-backend-golang.git
cd ticket-booking-backend-golang
```

### 2. Environment Variables

Create a `.env` file from `.env.example`:

```bash
cp .env.example .env
```

**`.env.example`:**

```env
SERVER_PORT=8080

DB_HOST=db
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_SSLMODE=disable

ALLOWED_ORIGINS='http://localhost:8081, http://127.0.0.1:8081'
```

> For local development outside Docker, you may need:

```env
JWT_SECRET=your_jwt_secret
DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

---

## 🐳 Docker Setup

### Start development environment:

```bash
docker-compose up --build
```

This will:

* Build the Go backend with [Air](https://github.com/air-verse/air) for hot-reloading
* Start a PostgreSQL database with persistent volume and health check
* Expose backend at [http://localhost:8080](http://localhost:8080)

### `docker-compose.yaml`

```yaml
version: "3.9"

services:
  app:
    tty: true
    restart: always
    image: ticket-booking
    container_name: ticket-booking
    build: .
    ports:
      - 8080:8080
    env_file:
      - .env
    networks:
      - application
    depends_on:
      db:
        condition: service_healthy
    volumes:
      - .:/src/app
    command: air -c .air.toml

  db:
    image: postgres:alpine
    container_name: ticket-booking-db
    environment:
      - POSTGRES_HOST=${DB_HOST}
      - POSTGRES_DB=${DB_NAME}
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/var/lib/postgresql/data
    networks:
      - application
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${DB_USER}"]
      interval: 10s
      timeout: 5s
      retries: 5

networks:
  application:

volumes:
  postgres-db:
```

### `Dockerfile`

```Dockerfile
FROM golang:1.22.2-alpine3.19

WORKDIR /src/app

RUN go install github.com/air-verse/air@v1.52.3

COPY . .

RUN go mod tidy
```

---

## 🧪 Run Locally (without Docker)

```bash
go run cmd/api/main.go
```

---

## 📡 API Endpoints

### 🔑 Auth

| Method | Endpoint             | Description       |
| ------ | -------------------- | ----------------- |
| POST   | `/api/auth/register` | Register new user |
| POST   | `/api/auth/login`    | Login and get JWT |

### 📅 Event (requires JWT)

| Method | Endpoint         | Description        |
| ------ | ---------------- | ------------------ |
| GET    | `/api/event`     | Get all events     |
| GET    | `/api/event/:id` | Get a single event |
| POST   | `/api/event`     | Create a new event |
| PUT    | `/api/event/:id` | Update an event    |
| DELETE | `/api/event/:id` | Delete an event    |

### 🎟 Ticket (requires JWT)

| Method | Endpoint          | Description            |
| ------ | ----------------- | ---------------------- |
| GET    | `/api/ticket`     | Get all user's tickets |
| GET    | `/api/ticket/:id` | Get a specific ticket  |
| POST   | `/api/ticket`     | Book a ticket          |
| PUT    | `/api/ticket/:id` | Update a ticket        |

> All `/event` and `/ticket` routes require a `Bearer <JWT>` token in the `Authorization` header.

---

## 📲 Frontend

The mobile frontend is built using React Native + Expo:

👉 [Ticket Booking Mobile App](https://github.com/pyroblazer/ticket-booking-mobile-react-native-expo)

---

## ✅ Tech Stack

* **Backend**: Golang, Fiber
* **Database**: PostgreSQL
* **ORM**: GORM
* **Authentication**: JWT
* **Dev Tools**: Docker, Air (for hot-reload)
* **Frontend**: React Native Expo (separate repo)

---

## 🧑‍💻 Contributing

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Submit a Pull Request (PR)

---

## 🙋 Contact

Made with ❤️ by [@pyroblazer](https://github.com/pyroblazer)
