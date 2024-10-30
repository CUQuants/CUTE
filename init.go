package main

import (
	"encoding/json"
	"net/http"
)

type InitConfig struct {
	ClientVersion       string `json:"clientVersion"`
	CurrentTimeUnix     int    `json:"currentTimeUnix"`
	StepIntervalSeconds int    `json:"stepIntervalSeconds"`
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	var req InitConfig
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	balance = 0
	portfolio = make(map[string]int)

	currentTimeUnix = req.CurrentTimeUnix
	stepIntervalSeconds = req.StepIntervalSeconds
}
