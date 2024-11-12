package fintoc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrice(t *testing.T) {
	stock := &Stock{
		Name:     "StockA",
		Quantity: 100,
		Prices: map[string]float64{
			"2024-01-01": 100.0,
			"2024-12-31": 120.0,
		},
	}

	tests := []struct {
		name      string
		date      string
		expected  float64
		expectErr bool
	}{
		{"Valid date 1", "2024-01-01", 100.0, false},
		{"Valid date 2", "2024-12-31", 120.0, false},
		{"Invalid date", "2023-12-31", 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, _ := time.Parse("2006-01-02", tt.date)
			price, err := stock.Price(date)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, price)
			}
		})
	}
}

func TestProfit(t *testing.T) {
	portfolio := NewPortfolio([]*Stock{
		{
			Name:     "StockA",
			Quantity: 50,
			Prices: map[string]float64{
				"2024-01-01": 100.0,
				"2024-12-31": 120.0,
			},
		},
		{
			Name:     "StockB",
			Quantity: 30,
			Prices: map[string]float64{
				"2024-01-01": 150.0,
				"2024-12-31": 180.0,
			},
		},
	})

	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-12-31")

	profit, err := portfolio.Profit(startDate, endDate)
	require.NoError(t, err)
	require.Equal(t, 3900.0, profit)
}

func TestAnnualizedReturn(t *testing.T) {
	portfolio := NewPortfolio([]*Stock{
		{
			Name:     "StockA",
			Quantity: 50,
			Prices: map[string]float64{
				"2024-01-01": 100.0,
				"2024-12-31": 120.0,
			},
		},
		{
			Name:     "StockB",
			Quantity: 30,
			Prices: map[string]float64{
				"2024-01-01": 150.0,
				"2024-12-31": 180.0,
			},
		},
	})

	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-12-31")

	annualReturn, err := portfolio.AnnualizedReturn(startDate, endDate)
	require.NoError(t, err)
	require.InEpsilon(t, 0.2, annualReturn, 0.01) // Annualized return of approx 20% with 1% tolerance
}

func TestNewPortfolio(t *testing.T) {
	stocks := []*Stock{
		{Name: "StockA", Quantity: 50},
		{Name: "StockB", Quantity: 30},
	}

	portfolio := NewPortfolio(stocks)
	require.Equal(t, 2, len(portfolio.Stocks))
	require.Contains(t, portfolio.Stocks, "StockA")
	require.Contains(t, portfolio.Stocks, "StockB")
}
