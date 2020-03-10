package main

import (
	"encoding/json"
	"fmt"
)

type Balance struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}

func getBalances(coin string) []Balance {
	r := request("GET", "wallet/balances")

	var response Response
	var balances []Balance
	response.Result = &balances
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	for _, balance := range balances {
		if coin == "ALL" {
			fmt.Printf("Coin/Balance: %v \t %v \n", balance.Coin, balance.Total)
		} else if coin == balance.Coin {
			fmt.Printf("Coin/Balance: %v \t %v \n", balance.Coin, balance.Total)
			break
		}
	}

	return balances

}
