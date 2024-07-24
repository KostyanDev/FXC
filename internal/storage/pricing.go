package storage

import (
	"context"

	"github.com/sirupsen/logrus"

	"app/internal/storage/dto"

	"app/internal/domain"
)

// GetPricingByDate retrieves pricing data for a given date from the database.
func (s storage) GetPricingByDate(ctx context.Context, data domain.RequestPricing) ([]domain.Pricing, error) {
	var pricingArr dto.PricingArr
	query := `
        SELECT 
            organization_name AS organization_name,
            transfer_amount AS transfer_amount,
            fx_rate AS fx_rate
        FROM pricing
        WHERE DATE(date) = DATE(?)
    `
	err := s.ext.SelectContext(ctx, &pricingArr, query, data.Date)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch pricing data")
		return nil, err
	}
	// Convert the results to the domain representation and return
	return pricingArr.ToDomain(), nil
}
