package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ImageData struct {
	ImageData string `json:"image_data"`
	Scale     int    `json:"scale"`
	Flip      bool   `json:"flip"`
}

type GraphData struct {
	Timestamp float64 `json:"timestamp"`
	Value     float64 `json:"value"`
}

type RobotPosition struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Heading float64 `json:"heading"`
}

type StringData struct {
	Value string `json:"value"`
}
type RobotData struct {
	SentTimestamp float64                `json:"sent_timestamp"`
	Images        map[string]ImageData   `json:"images"`
	GraphData     map[string][]GraphData `json:"graph_data"`
	StringData    map[string]StringData  `json:"string_data"`
	RobotPosition RobotPosition          `json:"robot_position"`
}

func stepLoop() {
	currentTimeUnix += stepIntervalSeconds
	t := time.Unix(currentTimeUnix, 0) // The second argument is for nanoseconds, so it's 0 here

	// Format the time to a human-readable string
	humanReadable := t.Format("2006-01-02 15:04:05 MST")
	println()

	fmt.Println("Time:", humanReadable)
	fmt.Println("Balance:", balance)
	fmt.Println("Portfolio:")
	portfolioEmpty := true
	for k, v := range portfolio {
		if v != 0 {
			fmt.Printf("    %s: %d\n", k, v)
			portfolioEmpty = false
		}
	}
	if portfolioEmpty {
		fmt.Println("   None")
	}
	fmt.Println("###################")

	var graph_data = RobotData{
		GraphData: map[string][]GraphData{
			"balance": {GraphData{Timestamp: float64(currentTimeUnix), Value: balance}},
		},
	}
	for k, v := range portfolio {
		graph_data.GraphData["holdings_"+k] = []GraphData{{Timestamp: float64(currentTimeUnix), Value: float64(v)}}
	}

	jsonData, err := json.Marshal(graph_data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = http.Post("http://localhost:3000/batch", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

}

func buyStock(symbol string, quantity int) error {
	data, err := fetchStockData(symbol, "1min", 3, currentTimeUnix)
	if err != nil {
		fmt.Println("err:", err.Error())
		return err
	}
	if len(data.Values) == 0 {
		err = fmt.Errorf("Empty values")
		println(err.Error())
		return err
	}
	_, exists := portfolio[symbol]
	if exists {
		portfolio[symbol] += quantity
	} else {
		portfolio[symbol] = quantity
	}

	unitPrice := data.Values[0].High
	price := unitPrice * float64(quantity)
	balance -= price * float64(quantity)
	fmt.Printf("You bought %d shares of %s for $%f\n", quantity, symbol, price)
	return nil
}

func sellStock(symbol string, quantity int) error {
	data, err := fetchStockData(symbol, "1min", 3, currentTimeUnix)
	if err != nil {
		fmt.Println("err:", err.Error())
		return err
	}
	if len(data.Values) == 0 {
		err = fmt.Errorf("Empty values")
		println(err.Error())
		return err
	}

	_, exists := portfolio[symbol]
	if exists {
		if quantity > portfolio[symbol] {
			return fmt.Errorf("You can't sell more of a stock than have (shorting will be added eventually)")
		}
		portfolio[symbol] -= quantity
	} else {
		return fmt.Errorf("You can't sell a stock you don't have (shorting will be added eventually)")
	}

	unitPrice := data.Values[0].High
	price := unitPrice * float64(quantity)
	balance += price * float64(quantity)
	fmt.Printf("You sold %d shares of %s for $%f\n", quantity, symbol, price)
	return nil
}
