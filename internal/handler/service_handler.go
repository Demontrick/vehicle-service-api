package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Demontrick/vehicle-service-api/internal/model"
	"github.com/Demontrick/vehicle-service-api/internal/service"
	"github.com/gorilla/mux"
)

type ServiceHandler struct {
	service *service.ServiceService
}

func NewServiceHandler(service *service.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: service}
}

func (h *ServiceHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/services", h.GetAll).Methods("GET")
	r.HandleFunc("/api/services", h.Create).Methods("POST")
	r.HandleFunc("/api/services/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/api/services/{id}/status", h.UpdateStatus).Methods("PATCH")
	r.HandleFunc("/api/services/{id}", h.Delete).Methods("DELETE")
}

func (h *ServiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	services := h.service.GetAllServices()
	writeJSON(w, http.StatusOK, services)
}

func (h *ServiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	created, err := h.service.CreateService(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *ServiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	service, err := h.service.GetServiceByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, service)
}

func (h *ServiceHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req model.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updated, err := h.service.UpdateStatus(id, req.Status)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *ServiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteService(id); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}