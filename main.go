package main

import (
	"time"

	"golang.org/x/exp/rand"
)

var secrets *Secrets

var currentTimeUnix int64
var stepIntervalSeconds int64 = 60 * 5
var balance float32 = 0
var portfolio map[string]int = make(map[string]int)

func main() {

	currentTimeUnix = 1673560800

	var err error
	secrets, err = loadSecrets("keys.json")
	if err != nil {
		panic(err)
	}
	// testData, err := fetchStockData("AAPL", "1min", 5, 1673560800)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// println(testData.Meta.Symbol)
	// fmt.Println("len", len(testData.Values))
	// println(testData.Values[len(testData.Values)-1].High)
	// println(testData.Values[0].Datetime)
	// println(testData.Values[0].Close)

	// _ = buyStock("AAPL", 1)

	for i := 0; i < 60; i++ {
		action := rand.Intn(3)
		if action == 0 {
			_ = buyStock("AAPL", 1)
		} else if action == 1 {
			_ = sellStock("AAPL", 1)
		}
		//do nothing on action == 3
		stepLoop()
		time.Sleep(16 * time.Second)
	}
}
