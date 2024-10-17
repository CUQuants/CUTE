package main

import (
	"fmt"
	"net/http"
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

	http.HandleFunc("/buy", buyHandler)
	http.HandleFunc("/sell", sellHandler)
	http.HandleFunc("/step", stepHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
