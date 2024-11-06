package main

import (
	"fmt"
	"net/http"

	updatechecker "github.com/Christian1984/go-update-checker"
)

var secrets *Secrets

var currentTimeUnix int64
var stepIntervalSeconds int64 = 60 * 5
var balance float32 = 0
var portfolio map[string]int = make(map[string]int)

var VERSION string = "0.0.1"

func main() {
	uc := updatechecker.New("CUQuants", "CUTE", "CUTE", "", 0, false)
	uc.CheckForUpdate(VERSION)
	uc.PrintMessage()

	currentTimeUnix = 1673560800

	var err error
	secrets, err = loadSecrets("keys.json")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/buy", buyHandler)
	http.HandleFunc("/sell", sellHandler)
	http.HandleFunc("/step", stepHandler)

	http.HandleFunc("/fetch-stock-data", fetchStockDataHandler)
	http.HandleFunc("/get-portfolio", getPortfolioHandler)

	fmt.Println("Server starting on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
