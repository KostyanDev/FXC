package service

import (
	"context"

	"app/internal/domain"

	"github.com/sirupsen/logrus"
)

type Storage interface {
	GetPricingByDate(ctx context.Context, data domain.RequestPricing) ([]domain.Pricing, error)
}
type Service struct {
	ctx     context.Context
	log     *logrus.Logger
	storage Storage
}

func New(ctx context.Context, log *logrus.Logger, storage Storage) *Service {
	return &Service{
		ctx:     ctx,
		log:     log,
		storage: storage,
	}
}
