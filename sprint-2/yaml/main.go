package main

import (
	"log"

	"gopkg.in/yaml.v3"
)

type (
	AccountBalance struct {
		AccountIDHash []byte           `yaml:"accountIdHash,flow"`
		Amounts       []CurrencyAmount `yaml:"amounts,omitempty"`
		IsBlocked     bool             `yaml:"isBlocked"`
	}

	CurrencyAmount struct {
		Amount   int64  `yaml:"amount"`
		Decimals uint8  `yaml:"decimals"`
		Symbol   string `yaml:"symbol"`
	}
)

func main() {
	balance := AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"}, //nolint:mnd // example
			{Amount: 2510, Decimals: 2, Symbol: "USD"},    //nolint:mnd // example
		},
		IsBlocked: true,
	}

	out, err := yaml.Marshal(balance)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(out))
}
