package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Order struct {
	ID            int     `json:"id"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	FilledSize    float64 `json:"filledSize"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`
	Status        string  `json:"status"`
	Type          string  `json:"type"`
}

func openOrders(coin string, print bool) []Order {
	limit := "100"

	r := request("GET", "orders/history?market="+coin+"/USD&&limit="+limit)

	var response Response
	var orders []Order
	response.Result = &orders
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	if print {
		for _, order := range orders {
			fmt.Printf("%v \t %v \t Size/Price: %v %v \n", order.Market, order.Side, order.Size, order.AvgFillPrice)
		}
	}

	return orders

}

// List all buys that match current coin holding and show P&L
func profitLoss() {

	orders := openOrders("MIDBULL", false)
	price := recentClose("MIDBULL", false)

	balance := getBalances("MIDBULL")
	// TODO: Get coin balance dynamically (not by current key value)
	total := balance[3].Total

	var size float64
	var sumPercentage float64
	var countTotal float64
	for _, order := range orders {

		// Ignore Sell orders
		if order.Side == "sell" {
			continue
		}

		// Break loop if we exceed our current wallet size
		if fmt.Sprintf("%.4f", size) >= fmt.Sprintf("%.4f", total) {
			break
		}

		// Add our fees back to our orders, so we can match the wallert total correctly
		size = size + order.FilledSize

		// Round values for buy & current values
		usdBuyValue := math.Round(order.Size*order.AvgFillPrice*100) / 100
		usdCurrentValue := math.Round(order.FilledSize*price[0].Close*100) / 100

		// Work out the percentage increase or decrease of our buys
		perc := math.Round((usdCurrentValue-usdBuyValue)/usdBuyValue*100*100) / 100

		fmt.Printf("|%v| $%-v \t |Current:| $%v  \t |%%+-| %v%%\n", order.Side, usdBuyValue, usdCurrentValue, perc)

		sumPercentage = sumPercentage + (perc / 100)
		countTotal++

	}

	fmt.Printf("\nOverall P&L Percentage %.2f", (sumPercentage/countTotal)*100)

}
