package storage

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"app/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/jmoiron/sqlx"
)

func TestStorage_GetPricingByDate(t *testing.T) {
	successRes := []domain.Pricing{
		{OrganizationName: "Company1", TransferAmount: 100, Rate: 5.1},
		{OrganizationName: "Company2", TransferAmount: 100, Rate: 4.9},
	}

	testError := errors.New("test error")
	testDateStr := "2024-01-10"
	testDate, err := time.Parse("2006-01-02", testDateStr)
	if err != nil {
		t.Fatalf("Failed to parse test date: %v", err)
	}

	type args struct {
		ctx  context.Context
		data domain.RequestPricing
	}

	tests := map[string]struct {
		mock    func(mock sqlmock.Sqlmock)
		args    args
		want    []domain.Pricing
		wantErr error
	}{
		"error.query execution": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`
                    SELECT 
                        organization_name AS organization_name,
                        transfer_amount AS transfer_amount,
                        fx_rate AS fx_rate
                    FROM pricing
                    WHERE DATE(date) = DATE(?)
                `)).
					WithArgs(testDate).
					WillReturnError(testError)
			},
			args: args{
				ctx:  context.Background(),
				data: domain.RequestPricing{Date: testDate},
			},
			want:    nil,
			wantErr: testError,
		},
		"success": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`
                    SELECT 
                        organization_name AS organization_name,
                        transfer_amount AS transfer_amount,
                        fx_rate AS fx_rate
                    FROM pricing
                    WHERE DATE(date) = DATE(?)
                `)).
					WithArgs(testDate).
					WillReturnRows(
						sqlmock.NewRows([]string{"organization_name", "transfer_amount", "fx_rate"}).
							AddRow("Company1", 100, 5.1).
							AddRow("Company2", 100, 4.9),
					)
			},
			args: args{
				ctx:  context.Background(),
				data: domain.RequestPricing{Date: testDate},
			},
			want:    successRes,
			wantErr: nil,
		},
	}

	for caseName, tt := range tests {
		t.Run(caseName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected on sqlmock.New", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			logger := logrus.New()

			s := New(logger, sqlxDB)

			tt.mock(mock)

			got, err := s.GetPricingByDate(tt.args.ctx, tt.args.data)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
