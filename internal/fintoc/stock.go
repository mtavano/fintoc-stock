package fintoc

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Stock struct {
	Name     string             `json:"name"`
	Quantity int                `json:"quantity"`
	Prices   map[string]float64 `json:"prices"`
}

type Portfolio struct {
	Stocks map[string]*Stock
}

func NewPortfolio(stocks []*Stock) *Portfolio {
	stockMap := make(map[string]*Stock)
	for _, stock := range stocks {
		stockMap[stock.Name] = stock
	}
	return &Portfolio{Stocks: stockMap}
}

func (p *Portfolio) Profit(startDate, endDate time.Time) (float64, error) {
	var totalStart, totalEnd float64

	for _, stock := range p.Stocks {
		startPrice, err := stock.Price(startDate)
		if err != nil {
			return 0, err
		}

		endPrice, err := stock.Price(endDate)
		if err != nil {
			return 0, err
		}

		totalStart += startPrice * float64(stock.Quantity)
		totalEnd += endPrice * float64(stock.Quantity)
	}

	return totalEnd - totalStart, nil
}

func (p *Portfolio) AnnualizedReturn(startDate, endDate time.Time) (float64, error) {
	profit, err := p.Profit(startDate, endDate)
	if err != nil {
		return 0, err
	}

	totalStart := 0.0
	for _, stock := range p.Stocks {
		startPrice, err := stock.Price(startDate)
		if err != nil {
			return 0, err
		}

		totalStart += startPrice * float64(stock.Quantity)
	}

	if totalStart == 0 {
		return 0, errors.New("fintoc: Portfolio.AnnualizedReturn error: initial value of the portfolio is zero")
	}

	years := endDate.Sub(startDate).Hours() / (24 * 365.25)
	annualizedReturn := math.Pow(1+profit/totalStart, 1/years) - 1

	return annualizedReturn, nil
}

func (s *Stock) Price(date time.Time) (float64, error) {
	dateStr := date.Format("2006-01-02")
	price, ok := s.Prices[dateStr]
	if !ok {
		return 0, errors.New(
			fmt.Sprintf("price not available for %s on %s", s.Name, dateStr),
		)
	}

	return price, nil
}
