package main

import (
	"encoding/json"
	"fmt"
)

type Market struct {
	Last float64 `json:"last"`
}

type Price struct {
	StartTime string  `json:"startTime"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}

func market(coin string) {
	r := request("GET", "markets/"+coin+"/USD")

	var response Response
	var market Market
	response.Result = &market
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	fmt.Println(market)
}

func recentClose(coin string) {
	r := request("GET", "markets/"+coin+"/USD/candles?resolution=15&limit=1")

	var response Response
	var price []Price
	response.Result = &price
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	fmt.Println(price[0].Close)
}
