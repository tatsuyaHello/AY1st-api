package service

import (
	"AY1st/model"
	"AY1st/repository"
)

// HealthCheckInterface is health check
type HealthCheckInterface interface {
	GetHealth(cond string) (*model.HealthCheck, error)
}

// HealthCheck is health check
type HealthCheck struct {
	HealthCheckRepo repository.HealthCheckInterface
}

// NewHealthCheck is health check
func NewHealthCheck(healthCheckRepo repository.HealthCheckInterface) *HealthCheck {
	h := HealthCheck{
		HealthCheckRepo: healthCheckRepo,
	}
	return &h
}

// GetHealth is health check
func (h *HealthCheck) GetHealth(cond string) (*model.HealthCheck, error) {
	return h.HealthCheckRepo.GetHealth(cond)
}
