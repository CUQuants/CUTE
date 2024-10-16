package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
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

type TwelveDataResponse struct {
	Meta   MetaData     `json:"meta"`
	Values []StockValue `json:"values"`
}

func fetchStockData(symbol string, interval string, n int, endTime int64) (*TwelveDataResponse, error) {
	endTimeISO := time.Unix(endTime, 0).UTC().Format(time.RFC3339)

	baseURL := "https://api.twelvedata.com/time_series"
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("interval", interval)
	params.Add("apikey", secrets.TwelvedataApiKey)
	params.Add("outputsize", fmt.Sprintf("%d", n))
	params.Add("end_date", endTimeISO)

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API response status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var data TwelveDataResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &data, nil
}
