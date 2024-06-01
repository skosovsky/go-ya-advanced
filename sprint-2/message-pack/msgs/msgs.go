package msgs

//go:generate msgp
type AccountBalance struct {
	AccountIDHash []byte           `msg:"account_id_hash"`
	Amounts       []CurrencyAmount `msg:"amounts"`
	IsBlocked     bool             `msg:"is_blocked"`
}

type CurrencyAmount struct {
	Amount   int64
	Decimals int8
	Symbol   string
}
