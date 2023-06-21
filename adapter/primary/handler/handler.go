package handler

import (
	"github.com/witalok2/test-dev-golang-api/config"
	"github.com/witalok2/test-dev-golang-api/internal/service"
)

type Handler struct {
	service     service.ServiceInterface
	environment *config.Environment
}

func NewHandler(serv service.ServiceInterface, env *config.Environment) *Handler {
	return &Handler{
		service:     serv,
		environment: env,
	}
}
