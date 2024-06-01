package jsons

//go:generate easyjson -all jsons.go
type (
	AccountBalance struct {
		AccountIDHash []byte           `json:"accountIdHash"`
		Amounts       []CurrencyAmount `json:"amounts,omitempty"`
		IsBlocked     bool             `json:"isBlocked"`
	}

	CurrencyAmount struct {
		Amount   int64  `json:"amount"`
		Decimals int8   `json:"decimals"`
		Symbol   string `json:"symbol"`
	}
)
