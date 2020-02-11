package repository

import (
	"AY1st/model"

	"github.com/go-xorm/xorm"
)

// HealthCheckInterface is health check (debug)
type HealthCheckInterface interface {
	GetHealth(cond string) (*model.HealthCheck, error)
}

// HealthCheck is health check (debug)
type HealthCheck struct {
	engine xorm.EngineInterface
}

// NewHealthCheck initializes HealthCheck
func NewHealthCheck(engine xorm.EngineInterface) *HealthCheck {
	h := HealthCheck{
		engine: engine,
	}
	return &h
}

// GetHealth is health check with DB
func (h *HealthCheck) GetHealth(cond string) (*model.HealthCheck, error) {

	out := &model.HealthCheck{}

	session := h.engine.NewSession()

	session = session.Table("health_check")

	if cond != "" {
		session.Where("id = ?", cond)
	}

	_, err := session.Get(out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
