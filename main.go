package main

var secrets *Secrets

func main() {
	secrets, _ = loadSecrets("keys.json")

	fetchData("AAPL")
}
