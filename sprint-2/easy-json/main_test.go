package main

import (
	"testing"

	"go-ya-advanced/sprint-2/easy-json/jsons"
)

func BenchmarkEasyJSON(b *testing.B) {
	balance := jsons.AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []jsons.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = generateEasyJSON(balance)
	}
}

func BenchmarkGoJSON(b *testing.B) {
	balance := jsons.AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []jsons.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = generateGoJSON(balance)
	}
}

func BenchmarkStdlibJSON(b *testing.B) {
	balance := jsons.AccountBalance{
		AccountIDHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []jsons.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = generateStdlibJSON(balance)
	}
}
