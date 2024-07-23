package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"app/internal/service"

	"app/internal/domain"
	"app/internal/transport/converters"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetPricingList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := logrus.New()
	mockStorage := service.NewMockStorage(mockCtrl)
	svc := service.New(context.Background(), logger, mockStorage)
	h := Handler{ctx: context.Background(), log: logger, service: svc}

	testDateStr := "2024-01-10"
	testDate, err := time.Parse("2006-01-02", testDateStr)
	if err != nil {
		t.Fatalf("Failed to parse test date: %v", err)
	}

	testReq := converters.PricingRequest{Date: testDateStr}

	requestBody, _ := json.Marshal(testReq)
	req := httptest.NewRequest("POST", "/price", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testPricing := []domain.Pricing{
		{OrganizationName: "Company1", TransferAmount: 100, Rate: 5.1},
		{OrganizationName: "Company2", TransferAmount: 100, Rate: 4.9},
	}

	testResponse := []converters.PricingResponse{
		{
			Amount: 100,
			Details: []converters.PricingDetail{
				{OrganizationName: "Company1", Rate: 5.1},
				{OrganizationName: "Company2", Rate: 4.9},
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockStorage.EXPECT().GetPricingByDate(gomock.Any(), domain.RequestPricing{Date: testDate}).Return(testPricing, nil)

		h.GetPricingList(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response []converters.PricingResponse
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, testResponse, response)
	})

	t.Run("error.decoding request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/price", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.GetPricingList(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("error.converting date", func(t *testing.T) {
		invalidDateReq := converters.PricingRequest{Date: "invalid-date"}
		requestBody, _ := json.Marshal(invalidDateReq)
		req := httptest.NewRequest("POST", "/price", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.GetPricingList(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
