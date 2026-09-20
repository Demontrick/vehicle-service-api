# Vehicle Service API

[![CI](https://github.com/Demontrick/vehicle-service-api/actions/workflows/ci.yml/badge.svg)](https://github.com/Demontrick/vehicle-service-api/actions/workflows/ci.yml)

## Project Overview

A RESTful microservice in Go for managing vehicle service requests. It is a proof of concept demonstrating clean architecture, Test-Driven Development and production-ready deployment practices.

---

## Engineering Approach

- **TDD**: tests written before the implementation, covering happy paths, validation and error scenarios
- **Clean Code**: single responsibility, clear naming, no unnecessary complexity
- **SOLID principles**: dependency injection through interfaces, so each layer can be tested on its own
- **CI/CD**: a GitHub Actions pipeline runs on every push and pull request to `master`
- **Docker**: multi-stage build for a small runtime image

---

## Tech Stack

- Go 1.26
- Gorilla Mux (HTTP router)
- Docker (multi-stage build)
- GitHub Actions (CI/CD)

---

## Architecture


vehicle-service-api/
├── cmd/
│   └── main.go              → application entry point
├── internal/
│   ├── handler/             → HTTP handlers (routing, request/response)
│   ├── model/               → domain structs and types
│   ├── repository/          → in-memory data store guarded by a mutex
│   └── service/             → business logic and validation
├── Dockerfile
└── .github/workflows/ci.yml


A layered handler / service / repository structure, familiar from Java systems and applied here in Go. Each layer depends on an interface, not a concrete type.

Architecture diagram: https://gitdiagram.com/demontrick/vehicle-service-api

---

## Design Decisions and Limitations

- **In-memory repository with a mutex.** This keeps the service simple, dependency-free and fast to test, and the mutex makes concurrent requests safe. The trade-off is that data does not persist across restarts and one lock guards all access.
- **Next step:** replace the in-memory repository with PostgreSQL behind the same repository interface, and run the tests against a real database.

---

## Service Request Statuses

PENDING → IN_PROGRESS → COMPLETED
        → CANCELLED

---

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/services | Get all service requests |
| POST | /api/services | Create a service request |
| GET | /api/services/{id} | Get a service request by ID |
| PATCH | /api/services/{id}/status | Update the status |
| DELETE | /api/services/{id} | Delete a service request |

---

## Example Request

`POST /api/services`

json
{
  "vehicleId": "VEH-001",
  "description": "Oil change required",
  "priority": "HIGH"
}


Response:

json
{
  "id": "0a405521-3c38-4e57-b3b8-2a1ca5989734",
  "vehicleId": "VEH-001",
  "description": "Oil change required",
  "status": "PENDING",
  "priority": "HIGH",
  "createdAt": "2026-04-06T12:41:16Z",
  "updatedAt": "2026-04-06T12:41:16Z"
}


---

## Tests

10 unit tests covering:

- Service creation: success and validation failures
- Get all services: empty and populated
- Get by ID: not found
- Status update: success and invalid status
- Delete: success and not found

bash
go test ./...

---

## How To Run

### Local

bash
git clone https://github.com/Demontrick/vehicle-service-api.git
cd vehicle-service-api
go run cmd/main.go


The API is available at http://localhost:8080

### Docker

bash
docker build -t vehicle-service-api .
docker run -p 8080:8080 vehicle-service-api


---

## CI/CD Pipeline

The GitHub Actions pipeline runs on every push and pull request to `master`:

1. Set up Go
2. Download dependencies
3. Run all tests
4. Build the binary
5. Build the Docker image

---

## Why Go

Java is my primary language. I chose Go for this project to learn a new language properly, and built it test-first in two days. Go's simplicity and opinionated structure suit clean architecture and TDD, and the project shows how quickly I can pick up an unfamiliar stack without cutting corners.

---

## Author

Aman Malik: [GitHub](https://github.com/Demontrick) · [LinkedIn](https://www.linkedin.com/in/aman-malik-b7b586242)
