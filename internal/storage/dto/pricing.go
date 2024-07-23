package dto

import "app/internal/domain"

type Pricing struct {
	OrganizationName string  `db:"organization_name"`
	TransferAmount   int     `db:"transfer_amount"`
	Rate             float32 `db:"fx_rate"`
}

type PricingArr []Pricing

func (p Pricing) ToDomain() domain.Pricing {
	return domain.Pricing{
		OrganizationName: p.OrganizationName,
		TransferAmount:   p.TransferAmount,
		Rate:             p.Rate,
	}
}

func (pr PricingArr) ToDomain() []domain.Pricing {
	prArr := make([]domain.Pricing, len(pr))
	for i := range pr {
		prArr[i] = pr[i].ToDomain()
	}
	return prArr
}
