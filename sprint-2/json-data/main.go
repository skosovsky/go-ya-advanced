package main

import (
	"encoding/json"
	"log"
)

type Data struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Company string `json:"comp,omitempty"`
}

func main() {
	foo := []Data{
		{ //nolint:exhaustruct // example
			ID:   10, //nolint:mnd // example
			Name: "Gopher",
		},
		{ //nolint:exhaustruct // example
			Name:    "Вася",
			Company: "Яндекс",
		},
	}

	out, err := json.Marshal(foo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(out))
}
