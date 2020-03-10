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

func profitLoss() {

	orders := openOrders("MIDBULL", false)
	price := recentClose("MIDBULL", false)

	balance := getBalances("MIDBULL")
	// TODO: Get coin ba;ance dynamically (not by current key value)
	total := balance[3].Total
	var size float64
	for _, order := range orders {

		if order.Side == "sell" {
			continue
		}

		size = size + order.FilledSize
		percAmount := (size / 100) * 0.070000
		size = size + percAmount

		if size > total {
			break
		}

		usdBuyValue := math.Round(order.Size*order.AvgFillPrice*100) / 100
		usdCurrentValue := math.Round(order.FilledSize*price[0].Close*100) / 100
		perc := math.Round((usdCurrentValue-usdBuyValue)/usdBuyValue*100*100) / 100
		fmt.Printf("|%v| %-v \t |Current:| %v  \t |%%+-| %v \n", order.Side, usdBuyValue, usdCurrentValue, perc)
	}

}
