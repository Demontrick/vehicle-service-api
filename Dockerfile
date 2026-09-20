FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o vehicle-service-api ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/vehicle-service-api .
EXPOSE 8080
ENTRYPOINT ["./vehicle-service-api"]
