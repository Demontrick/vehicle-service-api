package main

import (
	"log"
	"net/http"

	"github.com/Demontrick/vehicle-service-api/internal/handler"
	"github.com/Demontrick/vehicle-service-api/internal/repository"
	"github.com/Demontrick/vehicle-service-api/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewServiceRepository()
	svc := service.NewServiceService(repo)
	h := handler.NewServiceHandler(svc)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	log.Println("Vehicle Service API running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}