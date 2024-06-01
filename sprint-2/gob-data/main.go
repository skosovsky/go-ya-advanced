package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type (
	SendData struct {
		Value   int
		Balance *float64
		Name    string
		private int
	}

	GetData struct {
		Name    string
		Balance float64
		Ext     []byte
		value   int //nolint:unused // example
	}
)

func main() {
	floatValue := 50.0
	data := SendData{Value: 100, //nolint:mnd // example
		Balance: &floatValue,
		Name:    "Василий Кузнецов",
		private: 1}

	var buffer bytes.Buffer

	log.Printf("out: %+v\n", data)

	if err := gob.NewEncoder(&buffer).Encode(data); err != nil {
		log.Fatal(err)
	}

	var out GetData
	if err := gob.NewDecoder(&buffer).Decode(&out); err != nil {
		log.Fatal(err)
	}

	log.Printf("out: %+v\n", out)
}
