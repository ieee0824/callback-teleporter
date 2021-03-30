package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ieee0824/callback-teleporter/restapi/operations/health_check"
)

func NewHealthCheck() *HealthCheck {
	return new(HealthCheck)
}

type HealthCheck struct {
}

func (h *HealthCheck) Get(params health_check.GetHealthCheckParams) middleware.Responder {
	return health_check.NewGetHealthCheckOK()
}
