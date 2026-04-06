package service

import (
	"testing"

	"github.com/Demontrick/vehicle-service-api/internal/model"
	"github.com/Demontrick/vehicle-service-api/internal/repository"
)

func setup() *ServiceService {
	repo := repository.NewServiceRepository()
	return NewServiceService(repo)
}

func TestCreateService_Success(t *testing.T) {
	svc := setup()
	req := model.CreateServiceRequest{
		VehicleID:   "VW-001",
		Description: "Oil change required",
		Priority:    "HIGH",
	}

	result, err := svc.CreateService(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID == "" {
		t.Error("expected ID to be set")
	}
	if result.Status != model.StatusPending {
		t.Errorf("expected status PENDING, got %v", result.Status)
	}
	if result.VehicleID != "VW-001" {
		t.Errorf("expected vehicleId VW-001, got %v", result.VehicleID)
	}
}

func TestCreateService_MissingVehicleID(t *testing.T) {
	svc := setup()
	req := model.CreateServiceRequest{
		Description: "Oil change",
		Priority:    "HIGH",
	}

	_, err := svc.CreateService(req)

	if err == nil {
		t.Error("expected error for missing vehicleId")
	}
}

func TestCreateService_MissingDescription(t *testing.T) {
	svc := setup()
	req := model.CreateServiceRequest{
		VehicleID: "VW-001",
		Priority:  "HIGH",
	}

	_, err := svc.CreateService(req)

	if err == nil {
		t.Error("expected error for missing description")
	}
}

func TestGetAllServices_ReturnsEmpty(t *testing.T) {
	svc := setup()
	result := svc.GetAllServices()

	if len(result) != 0 {
		t.Errorf("expected empty list, got %v", len(result))
	}
}

func TestGetAllServices_ReturnsCreated(t *testing.T) {
	svc := setup()
	svc.CreateService(model.CreateServiceRequest{
		VehicleID:   "VW-001",
		Description: "Tyre replacement",
		Priority:    "MEDIUM",
	})

	result := svc.GetAllServices()

	if len(result) != 1 {
		t.Errorf("expected 1 service, got %v", len(result))
	}
}

func TestGetServiceByID_NotFound(t *testing.T) {
	svc := setup()

	_, err := svc.GetServiceByID("non-existent-id")

	if err == nil {
		t.Error("expected error for non-existent service")
	}
}

func TestUpdateStatus_Success(t *testing.T) {
	svc := setup()
	created, _ := svc.CreateService(model.CreateServiceRequest{
		VehicleID:   "VW-002",
		Description: "Brake inspection",
		Priority:    "HIGH",
	})

	updated, err := svc.UpdateStatus(created.ID, model.StatusInProgress)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Status != model.StatusInProgress {
		t.Errorf("expected status IN_PROGRESS, got %v", updated.Status)
	}
}

func TestUpdateStatus_InvalidStatus(t *testing.T) {
	svc := setup()
	created, _ := svc.CreateService(model.CreateServiceRequest{
		VehicleID:   "VW-003",
		Description: "Engine check",
		Priority:    "HIGH",
	})

	_, err := svc.UpdateStatus(created.ID, "INVALID_STATUS")

	if err == nil {
		t.Error("expected error for invalid status")
	}
}

func TestDeleteService_Success(t *testing.T) {
	svc := setup()
	created, _ := svc.CreateService(model.CreateServiceRequest{
		VehicleID:   "VW-004",
		Description: "AC repair",
		Priority:    "LOW",
	})

	err := svc.DeleteService(created.ID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteService_NotFound(t *testing.T) {
	svc := setup()

	err := svc.DeleteService("non-existent-id")

	if err == nil {
		t.Error("expected error for non-existent service")
	}
}