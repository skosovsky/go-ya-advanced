package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Subj struct {
	Product string `json:"nane"`
	Price   int    `json:"price"`
}

func JSONHandler(res http.ResponseWriter, _ *http.Request) {
	subj := Subj{"Milk", 50}

	resJSON, err := json.Marshal(subj)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(resJSON)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", JSONHandler)

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
