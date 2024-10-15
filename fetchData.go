package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MetaData struct {
	Symbol     string `json:"symbol"`
	Interval   string `json:"interval"`
	Currency   string `json:"currency"`
	ExchangeTZ string `json:"exchange_timezone"`
	Exchange   string `json:"exchange"`
	MicCode    string `json:"mic_code"`
	Type       string `json:"type"`
}

type StockValue struct {
	Datetime string `json:"datetime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}

type StockData struct {
	Meta   MetaData     `json:"meta"`
	Values []StockValue `json:"values"`
}

func fetchData(ticker string) {
	url := "https://api.twelvedata.com/time_series?symbol=" + ticker + ",EUR/USD,ETH/BTC:Huobi,TRP:TSX&interval=1min&apikey=" + secrets.TwelvedataApiKey
	println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(io.Reader(resp.Body)) // response body is []byte

	var result map[string]StockData

	// Unmarshal the JSON into the map
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for key, value := range result {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

}
