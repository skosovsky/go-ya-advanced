package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{ //nolint:gochecknoglobals // example
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

func modelHandle(w http.ResponseWriter, r *http.Request) {
	carName := strings.ToLower(chi.URLParam(r, "brand") + ` ` + chi.URLParam(r, "model"))

	w.Header().Set("Content-Type", "text/plain")

	for _, car := range cars {
		if strings.ToLower(car) == carName {
			_, err := io.WriteString(w, car)
			if err != nil {
				return
			}

			return
		}
	}

	http.Error(w, "unknown model: "+carName, http.StatusNotFound)
}

func CarRouter() chi.Router { //nolint:ireturn // example
	router := chi.NewRouter()
	router.Get("/cars/{brand}/{model}", modelHandle) // GET /cars/renault/duster

	return router
}

func main() {
	err := http.ListenAndServe("localhost:8080", CarRouter()) //nolint:gosec // example
	if err != nil {
		log.Fatal(err)
	}
}
