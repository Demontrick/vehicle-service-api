# Vehicle Service API

## Project Overview

A RESTful microservice built in Go for managing vehicle service requests — demonstrating clean architecture, Test-Driven Development, and production-ready deployment practices.

Built as a proof of concept targeting enterprise automotive software engineering roles, specifically reflecting the Extreme Programming culture valued at organisations like Volkswagen Group Digital Solutions.

---

## Engineering Approach

This project was built following Extreme Programming principles:

- **TDD** — tests written before implementation, covering happy paths, validation, and error scenarios
- **Clean Code** — single responsibility, clear naming, no unnecessary complexity
- **SOLID principles** — dependency injection throughout, interfaces for testability
- **CI/CD** — GitHub Actions pipeline runs on every push to main
- **Docker** — production-ready containerisation with multi-stage build

---

## Tech Stack

- Go 1.22
- Gorilla Mux (HTTP router)
- Docker + Docker Compose
- GitHub Actions (CI/CD)

---

## Architecture
vehicle-service-api/
├── cmd/
│   └── main.go              → application entry point
├── internal/
│   ├── handler/             → HTTP handlers (routing + request/response)
│   ├── model/               → domain structs and types
│   ├── repository/          → in-memory data store with mutex for concurrency
│   └── service/             → business logic + validation
├── Dockerfile
└── .github/workflows/ci.yml
Clean separation of concerns — same layered architecture pattern used in enterprise Java systems, applied to Go.

---

## Service Request Lifecycle
PENDING → IN_PROGRESS → COMPLETED
→ CANCELLED

---

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/services | Get all service requests |
| POST | /api/services | Create new service request |
| GET | /api/services/{id} | Get service request by ID |
| PATCH | /api/services/{id}/status | Update service status |
| DELETE | /api/services/{id} | Delete service request |

---

## Example Request

POST /api/services
```json
{
  "vehicleId": "VW-001",
  "description": "Oil change required",
  "priority": "HIGH"
}
```

Response:
```json
{
  "id": "0a405521-3c38-4e57-b3b8-2a1ca5989734",
  "vehicleId": "VW-001",
  "description": "Oil change required",
  "status": "PENDING",
  "priority": "HIGH",
  "createdAt": "2026-04-06T12:41:16Z",
  "updatedAt": "2026-04-06T12:41:16Z"
}
```

---

## Tests

10 unit tests covering:
- Service creation — success and validation failures
- Get all services — empty and populated
- Get by ID — not found scenario
- Status update — success and invalid status
- Delete — success and not found

Run tests:
```bash
go test ./...
```

---

## How To Run

### Local
```bash
git clone https://github.com/Demontrick/vehicle-service-api.git
cd vehicle-service-api
go run cmd/main.go
```

API available at http://localhost:8080

### Docker
```bash
docker build -t vehicle-service-api .
docker run -p 8080:8080 vehicle-service-api
```

---

## CI/CD Pipeline

GitHub Actions pipeline on every push to main:

1. Go environment setup
2. Dependencies downloaded
3. All tests executed
4. Binary built
5. Docker image built

Pipeline status: ![CI](https://github.com/Demontrick/vehicle-service-api/actions/workflows/ci.yml/badge.svg)

---

## Why Go

Go was chosen deliberately for this project — not Java, which is my primary language. Go's simplicity, performance, and opinionated structure align naturally with Clean Code and Extreme Programming principles. Building this in Go demonstrates adaptability and openness to learn new languages — a core requirement for this role.

---

## Author

Portfolio project demonstrating Go microservice development with TDD, Clean Code, and production-ready deployment practices.