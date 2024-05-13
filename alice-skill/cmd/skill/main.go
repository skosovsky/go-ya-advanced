package main

import (
	"fmt"
	"net/http"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`
	{
		"response": {
			"text": "Sorry, I'm noob"
		},
		"version": "1.0"
	}
	`))
}

func run() error {
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(webhook)) //nolint:gosec // it's example

	return fmt.Errorf("could not start server: %w", err)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
