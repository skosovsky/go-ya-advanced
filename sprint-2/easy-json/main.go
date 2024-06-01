package main

import (
	"encoding/json"
	"log"

	"github.com/mailru/easyjson"

	gojson "github.com/goccy/go-json"

	"go-ya-advanced/sprint-2/easy-json/jsons"
)

func main() {
	balance := jsons.AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []jsons.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"}, //nolint:mnd // example
			{Amount: 2510, Decimals: 2, Symbol: "USD"},    //nolint:mnd // example
		},
		IsBlocked: true,
	}

	out := generateEasyJSON(balance)
	log.Println(string(out))

	out = generateGoJSON(balance)
	log.Println(string(out))

	out = generateStdlibJSON(balance)
	log.Println(string(out))
}

func generateEasyJSON(jsn jsons.AccountBalance) []byte {
	out, err := easyjson.Marshal(jsn)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func generateGoJSON(jsn jsons.AccountBalance) []byte {
	out, err := gojson.Marshal(jsn)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func generateStdlibJSON(jsn jsons.AccountBalance) []byte {
	out, err := json.Marshal(jsn)
	if err != nil {
		log.Fatal(err)
	}

	return out
}
