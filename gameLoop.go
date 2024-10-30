package main

import (
	"fmt"
	"time"
)

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
	price := unitPrice * float32(quantity)
	balance -= price * float32(quantity)
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
	price := unitPrice * float32(quantity)
	balance += price * float32(quantity)
	fmt.Printf("You sold %d shares of %s for $%f\n", quantity, symbol, price)
	return nil
}
