package handler

import (
	"net/http"

	"AY1st/model"
	"AY1st/registry"

	"github.com/gin-gonic/gin"
)

// GetHealth is get health check
func GetHealth(c *gin.Context) {
	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	healthCheckSearvice := servicer.NewHealthCheck()

	var input model.HealthCheckSearchInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	output, err := healthCheckSearvice.GetHealth(input.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
