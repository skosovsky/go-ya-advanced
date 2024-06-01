package main

import (
	"log"

	"github.com/pelletier/go-toml/v2"
)

type (
	AccountBalance struct {
		AccountIDHash []byte           `toml:"accountIdHash"`
		Amounts       []CurrencyAmount `toml:"amounts,omitempty"`
		IsBlocked     bool             `toml:"isBlocked"         comment:"Deprecated" commented:"true"`
	}

	CurrencyAmount struct {
		Amount   int64  `toml:"amount"`
		Decimals uint8  `toml:"decimals"`
		Symbol   string `toml:"symbol"`
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

	out, err := toml.Marshal(balance)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(out))
}
