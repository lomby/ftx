package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID            int     `json:"id"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`
	Status        string  `json:"status"`
}

func openOrders(coin string) {
	limit := "100"
	r := request("GET", "orders/history?market="+coin+"/USD&&limit="+limit)

	var response Response
	var orders []Order
	response.Result = &orders
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	for _, order := range orders {
		fmt.Println(order)
	}

	fmt.Println(orders)
}
