package main

var secrets *Secrets

func main() {
	var err error
	secrets, err = loadSecrets("keys.json")
	if err != nil {
		panic(err)
	}
	testData, err := fetchStockData("AAPL", "1min", 5, 1673560800)
	println(testData.Meta.Symbol)
	println(testData.Values[0].Datetime)
	println(testData.Values[0].Close)

}
