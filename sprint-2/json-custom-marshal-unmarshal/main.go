package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Visitor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name,omitempty"`
	Phones  []string `json:"phones,omitempty"`
	Company string   `json:"company,omitempty"`
}

func (v Visitor) MarshalJSON() ([]byte, error) {
	type CustomVisitor struct {
		ID      string   `json:"id"`
		Name    string   `json:"name,omitempty"`
		Phones  []string `json:"phones,omitempty"`
		Company string   `json:"company,omitempty"`
	}

	customVisitor := CustomVisitor{
		ID:      strconv.Itoa(v.ID),
		Name:    v.Name,
		Phones:  v.Phones,
		Company: v.Company,
	}

	data, err := json.Marshal(customVisitor)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return data, nil
}

func (v *Visitor) UnmarshalJSON(data []byte) error {
	type AliasVisitor struct {
		*Visitor
		UnmarshalJSON struct{}
	}

	aliasVisitor := AliasVisitor{
		Visitor:       v,
		UnmarshalJSON: struct{}{},
	}

	err := json.Unmarshal(data, &aliasVisitor) //nolint:musttag // pointer
	if err == nil {
		return nil
	}

	type CustomVisitor struct {
		ID      string   `json:"id"`
		Name    string   `json:"name,omitempty"`
		Phones  []string `json:"phones,omitempty"`
		Company string   `json:"company,omitempty"`
	}

	var customVisitor CustomVisitor

	if err = json.Unmarshal(data, &customVisitor); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	v.ID, err = strconv.Atoi(customVisitor.ID)
	if err != nil {
		return fmt.Errorf("failed to parse ID: %w", err)
	}
	v.Name = customVisitor.Name
	v.Phones = customVisitor.Phones
	v.Company = customVisitor.Company

	return nil
}

var visitors = map[string]Visitor{ //nolint:gochecknoglobals // example
	"1": {
		ID:      1,
		Name:    "John Doe",
		Phones:  []string{"123-456-789", "123-456-678"},
		Company: "",
	},
}

func JSONGetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func JSONPostHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	var visitor Visitor

	if err := json.NewDecoder(r.Body).Decode(&visitor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	id = strconv.Itoa(visitor.ID)
	visitors[id] = visitor

	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", JSONGetHandler)
	mux.HandleFunc("POST /", JSONPostHandler)

	go func() {
		time.Sleep(time.Second)
		resp, err := http.Post("http://localhost:8080/", "application/json", //nolint:noctx // example
			bytes.NewBufferString(`{"id": 10, "name": "Gopher", "company": "Don't Panic"}`))
		if err != nil {
			log.Println(err)

			return
		}

		err = resp.Body.Close()
		if err != nil {
			log.Println(err)

			return
		}
	}()

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // example
	if err != nil {
		log.Fatal(err)
	}
}
