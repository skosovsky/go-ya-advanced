package main

import "net/http"

func StatusHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		return
	}
}
