package model

// HealthCheck for debug table
type HealthCheck struct {
	ID     uint64 `json:"id"`
	Health string `json:"health"`
}

// HealthCheckSearchInput for debug table
type HealthCheckSearchInput struct {
	ID string `form:"id"`
}
