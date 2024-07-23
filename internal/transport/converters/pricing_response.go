package converters

import "app/internal/domain"

type PricingResponse struct {
	Amount  int             `json:"amount"`
	Details []PricingDetail `json:"details"`
}

type PricingDetail struct {
	OrganizationName string  `json:"organization_name"`
	Rate             float32 `json:"rate"`
}

func DomainPricingToResponsePricing(pricingArr []domain.Pricing) []PricingResponse {
	pricingMap := make(map[int][]PricingDetail)

	for _, pricing := range pricingArr {
		detail := PricingDetail{
			OrganizationName: pricing.OrganizationName,
			Rate:             pricing.Rate,
		}
		pricingMap[pricing.TransferAmount] = append(pricingMap[pricing.TransferAmount], detail)
	}

	var responses []PricingResponse
	for amount, details := range pricingMap {
		response := PricingResponse{
			Amount:  amount,
			Details: details,
		}
		responses = append(responses, response)
	}

	return responses
}
