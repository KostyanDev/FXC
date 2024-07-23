package converters

import (
	"time"

	"app/internal/domain"
)

type PricingRequest struct {
	Date string `json:"date"`
}

func ToDomainPricing(date time.Time) domain.RequestPricing {
	return domain.RequestPricing{
		Date: date,
	}
}
