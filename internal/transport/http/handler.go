package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/internal/transport/converters"
)

// GetPricingList handles HTTP requests to retrieve pricing data.
func (h *Handler) GetPricingList(w http.ResponseWriter, r *http.Request) {
	var request converters.PricingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	date, err := timeConvert(request.Date)
	if err != nil {
		h.log.Error("Error converting date: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if date.IsZero() {
		h.log.Error("Date is zero")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the pricing data from the service
	resp, err := h.service.GetPricing(r.Context(), converters.ToDomainPricing(date))
	if err != nil {
		h.log.Error("Error getting pricing: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the content type to JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(converters.DomainPricingToResponsePricing(resp)); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
