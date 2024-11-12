package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mtavano/fintoc-challenge/internal/fintoc"
)

func loadStocksFromFile(filename string) ([]*fintoc.Stock, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var stocks []*fintoc.Stock
	err = json.Unmarshal(data, &stocks)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func main() {
	stocks, err := loadStocksFromFile("stock.json")
	if err != nil {
		fmt.Println("Error loading stocks:", err)
		return
	}

	portfolio := fintoc.NewPortfolio(stocks)

	// dates to test
	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-06-01")

	profit, err := portfolio.Profit(startDate, endDate)
	if err != nil {
		fmt.Println("Error calculating profit:", err)
		return
	}
	fmt.Printf("Profit from %s to %s: %.2f%%\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), profit)

	annualReturn, err := portfolio.AnnualizedReturn(startDate, endDate)
	if err != nil {
		fmt.Println("Error calculating annualized return:", err)
		return
	}
	fmt.Printf("Annualized Return from %s to %s: %.2f%%\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), annualReturn*100)
}
