package main

import (
	"log"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func main() {
	balance := AccountBalance{
		state:         protoimpl.MessageState{}, //nolint:exhaustruct // system
		sizeCache:     0,
		unknownFields: nil,
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []*CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"}, //nolint:mnd // example
			{Amount: 2510, Decimals: 2, Symbol: "USD"},    //nolint:mnd // example
		},
		IsBlocked: true,
	}

	protoData, err := proto.Marshal(&balance)
	if err != nil {
		log.Fatal(err)
	}

	var newBalance AccountBalance

	if err = proto.Unmarshal(protoData, &newBalance); err != nil {
		log.Fatal(err)
	}

	log.Println(balance.String())
	log.Println(newBalance.String())
}
