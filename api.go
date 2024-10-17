package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BuyRequest struct {
	Symbol   string `json:"symbol"`
	Quantity int    `json:"quantity"`
}

type SellRequest struct {
	Symbol   string `json:"symbol"`
	Quantity int    `json:"quantity"`
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	var req BuyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := buyStock(req.Symbol, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Bought " + strconv.Itoa(req.Quantity) + " shares of " + req.Symbol,
	})
}

func sellHandler(w http.ResponseWriter, r *http.Request) {
	var req SellRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Ensure the body is closed

	if err := sellStock(req.Symbol, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sold " + strconv.Itoa(req.Quantity) + " shares of " + req.Symbol,
	})
}

func stepHandler(w http.ResponseWriter, r *http.Request) {
	stepLoop()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Step loop executed.",
	})
}
