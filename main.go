package main

func main() {
	secrets, err := loadSecrets("keys.json")
	if err != nil {
		panic("Couldn't load secrets")
	}

	println(secrets.TwelvedataApiKey)

	println("Hello world!")

}
