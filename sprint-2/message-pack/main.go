package main

import (
	"log"

	"go-ya-advanced/sprint-2/message-pack/msgs"
)

func main() {
	balance := msgs.AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []msgs.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"}, //nolint:mnd // example
			{Amount: 2510, Decimals: 2, Symbol: "USD"},    //nolint:mnd // example
		},
		IsBlocked: true,
	}

	msg, err := balance.MarshalMsg(nil)
	if err != nil {
		log.Fatal(err)
	}

	var response msgs.AccountBalance
	if _, err = response.UnmarshalMsg(msg); err != nil {
		log.Fatal(err)
	}

	log.Println(balance)
	log.Println(response)
}
