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
	Datetime string  `json:"datetime"`
	Open     float32 `json:"open,string"`   // Convert from string to float64
	High     float32 `json:"high,string"`   // Convert from string to float64
	Low      float32 `json:"low,string"`    // Convert from string to float64
	Close    float32 `json:"close,string"`  // Convert from string to float64
	Volume   int     `json:"volume,string"` // Convert from string to int
}

type TwelveDataResponse struct {
	Meta    MetaData     `json:"meta"`
	Values  []StockValue `json:"values"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}

func fetchStockData(symbol string, interval string, n int, endTime int64) (*TwelveDataResponse, error) {
	if endTime > currentTimeUnix {
		return nil, fmt.Errorf("you can't look into the future!")
	}

	// endTimeISO := time.Unix(endTime, 0).UTC().Format(time.RFC3339)

	endDateISO := time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
	baseURL := "https://api.twelvedata.com/time_series"
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("interval", interval)
	params.Add("apikey", secrets.TwelvedataApiKey)
	params.Add("outputsize", fmt.Sprintf("%d", n))
	params.Add("end_date", endDateISO)

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

	if data.Status != "ok" {
		return nil, fmt.Errorf("%d: %s", data.Code, data.Message)
	}

	return &data, nil
}
