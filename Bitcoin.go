package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BitcoinData struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"
	client := &http.Client{}
	res, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var bitcoinData BitcoinData
	json.Unmarshal([]byte(body), &bitcoinData)

	fmt.Println("Bitcoin Daten vom: ", bitcoinData.Time.Updated)
	fmt.Println("Aktueller Kurs in EUR: ", bitcoinData.Bpi.EUR.Rate)

	fmt.Print("Bitte drücke eine beliebige Taste: ")
	var sth string
	fmt.Scan(&sth)
}
