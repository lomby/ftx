package main

import "encoding/json"

type Account struct {
	Username string  `json:"username"`
	MakerFee float64 `json:"makerFee"`
	TakerFee float64 `json:"takerFee"`
}

func account() Account {
	r := request("GET", "account")

	var response Response
	var account Account
	response.Result = &account
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&response)

	return account

}
