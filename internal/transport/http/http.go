package http

import (
	"context"

	"app/internal/service"

	"github.com/sirupsen/logrus"
)

type Service interface{}

type Handler struct {
	ctx     context.Context
	log     *logrus.Logger
	service Service
}

func New(ctx context.Context, log *logrus.Logger, service *service.Service) *Handler {
	return &Handler{
		ctx:     ctx,
		log:     log,
		service: service,
	}
}
