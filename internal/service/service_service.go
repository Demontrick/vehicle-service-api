package service

import (
	"errors"

	"github.com/Demontrick/vehicle-service-api/internal/model"
	"github.com/Demontrick/vehicle-service-api/internal/repository"
)

type ServiceService struct {
	repo *repository.ServiceRepository
}

func NewServiceService(repo *repository.ServiceRepository) *ServiceService {
	return &ServiceService{repo: repo}
}

func (s *ServiceService) CreateService(req model.CreateServiceRequest) (model.ServiceRequest, error) {
	if req.VehicleID == "" {
		return model.ServiceRequest{}, errors.New("vehicleId is required")
	}
	if req.Description == "" {
		return model.ServiceRequest{}, errors.New("description is required")
	}
	if req.Priority == "" {
		return model.ServiceRequest{}, errors.New("priority is required")
	}
	return s.repo.Create(req), nil
}

func (s *ServiceService) GetAllServices() []model.ServiceRequest {
	return s.repo.FindAll()
}

func (s *ServiceService) GetServiceByID(id string) (model.ServiceRequest, error) {
	return s.repo.FindByID(id)
}

func (s *ServiceService) UpdateStatus(id string, status model.Status) (model.ServiceRequest, error) {
	validStatuses := map[model.Status]bool{
		model.StatusPending:    true,
		model.StatusInProgress: true,
		model.StatusCompleted:  true,
		model.StatusCancelled:  true,
	}
	if !validStatuses[status] {
		return model.ServiceRequest{}, errors.New("invalid status: " + string(status))
	}
	return s.repo.UpdateStatus(id, status)
}

func (s *ServiceService) DeleteService(id string) error {
	return s.repo.Delete(id)
}