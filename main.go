package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

type Response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func info() {
	app.Name = "MidBull Trader"
	app.Usage = "Personal app for trading midbull"
	app.Author = "L0mby"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "price",
			Aliases: []string{"p"},
			Usage:   "Get recent close Price",
			Action: func(c *cli.Context) {
				recentClose("MIDBULL", true)
			},
		},
		{
			Name:    "balances",
			Aliases: []string{"b"},
			Usage:   "List Balances (All coins)",
			Action: func(c *cli.Context) {
				getBalances("ALL")
			},
		},
		{
			Name:    "orders",
			Aliases: []string{"o"},
			Usage:   "Open Orders",
			Action: func(c *cli.Context) {
				openOrders("MIDBULL", true)
			},
		},
		{
			Name:    "profitloss",
			Aliases: []string{"pl"},
			Usage:   "Profit and Loss",
			Action: func(c *cli.Context) {
				profitLoss()
			},
		},
	}
}

func authRequest(method, path string) (string, string) {

	secret := os.Getenv("API_SECRET")
	ts := fmt.Sprintf("%d", time.Now().UTC().UnixNano()/1000000)
	message := ts + method + "/api/" + path
	sig := hmac.New(sha256.New, []byte(secret))

	sig.Write([]byte(message))

	return hex.EncodeToString(sig.Sum(nil)), ts

}

func request(method, path string) *http.Response {

	client := &http.Client{}

	key := os.Getenv("API_KEY")

	signature, ts := authRequest(method, path)
	req, err := http.NewRequest(method, "https://ftx.com/api/"+path, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("FTX-KEY", key)
	req.Header.Add("FTX-TS", ts)
	req.Header.Add("FTX-SIGN", signature)

	r, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	return r

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	info()
	commands()

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// openOrders("MIDBULL")
	// recentClose("MIDBULL")
}
