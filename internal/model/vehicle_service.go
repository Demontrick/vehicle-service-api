package model

import "time"

type Status string

const (
	StatusPending    Status = "PENDING"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusCancelled  Status = "CANCELLED"
)

type ServiceRequest struct {
	ID          string    `json:"id"`
	VehicleID   string    `json:"vehicleId"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateServiceRequest struct {
	VehicleID   string `json:"vehicleId"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type UpdateStatusRequest struct {
	Status Status `json:"status"`
}