package service

import (
	"context"

	"app/internal/domain"
)

func (s Service) GetPricing(ctx context.Context, data domain.RequestPricing) ([]domain.Pricing, error) {
	return s.storage.GetPricingByDate(ctx, data)
}
