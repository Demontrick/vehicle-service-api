package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/Demontrick/vehicle-service-api/internal/model"
	"github.com/google/uuid"
)

type ServiceRepository struct {
	mu       sync.RWMutex
	services map[string]model.ServiceRequest
}

func NewServiceRepository() *ServiceRepository {
	return &ServiceRepository{
		services: make(map[string]model.ServiceRequest),
	}
}

func (r *ServiceRepository) Create(req model.CreateServiceRequest) model.ServiceRequest {
	r.mu.Lock()
	defer r.mu.Unlock()

	service := model.ServiceRequest{
		ID:          uuid.New().String(),
		VehicleID:   req.VehicleID,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      model.StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	r.services[service.ID] = service
	return service
}

func (r *ServiceRepository) FindAll() []model.ServiceRequest {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.ServiceRequest, 0, len(r.services))
	for _, s := range r.services {
		result = append(result, s)
	}
	return result
}

func (r *ServiceRepository) FindByID(id string) (model.ServiceRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	service, exists := r.services[id]
	if !exists {
		return model.ServiceRequest{}, errors.New("service request not found with id: " + id)
	}
	return service, nil
}

func (r *ServiceRepository) UpdateStatus(id string, status model.Status) (model.ServiceRequest, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	service, exists := r.services[id]
	if !exists {
		return model.ServiceRequest{}, errors.New("service request not found with id: " + id)
	}

	service.Status = status
	service.UpdatedAt = time.Now()
	r.services[id] = service
	return service, nil
}

func (r *ServiceRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[id]; !exists {
		return errors.New("service request not found with id: " + id)
	}

	delete(r.services, id)
	return nil
}