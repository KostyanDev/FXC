package http

import (
	"app/internal/service"
)

type Service interface{}

type Handler struct {
	service Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
