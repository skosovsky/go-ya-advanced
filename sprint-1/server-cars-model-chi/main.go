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

type vehicle struct {
	id    string
	brand string
	model string
}

var vehicles = make(map[string]vehicle, len(cars)) //nolint:gochecknoglobals // example

func createCarsList() {
	for id, car := range cars {
		data := strings.Split(car, " ")

		vehicles[id] = vehicle{
			id:    id,
			brand: data[0],
			model: data[1],
		}
	}
}

func getCar(id string) string {
	if car, ok := vehicles[id]; ok {
		return car.brand + " " + car.model
	}

	return "unknown " + id
}

func carsHandle(w http.ResponseWriter, _ *http.Request) {
	var carsList string
	for _, car := range vehicles {
		carsList += car.brand + " " + car.model + ", "
	}

	w.Header().Set("Content-Type", "text/plain")

	_, err := io.WriteString(w, carsList)
	if err != nil {
		log.Println(err)

		return
	}
}

func carsBrandHandle(w http.ResponseWriter, r *http.Request) {
	carBrand := chi.URLParam(r, "brand")

	var carsBrandList string
	for _, car := range vehicles {
		if car.brand == carBrand {
			carsBrandList += car.brand + " " + car.model + " "
		}
	}

	w.Header().Set("Content-Type", "text/plain")

	_, err := io.WriteString(w, carsBrandList)
	if err != nil {
		log.Println(err)

		return
	}
}

func carsBrandModelHandle(w http.ResponseWriter, r *http.Request) {
	carBrand := chi.URLParam(r, "brand")
	carModel := chi.URLParam(r, "model")

	var carsBrandModelList string
	for _, car := range vehicles {
		if car.brand == carBrand && carModel == car.model {
			carsBrandModelList += car.brand + " " + car.model + ", "
		}
	}

	w.Header().Set("Content-Type", "text/plain")

	_, err := io.WriteString(w, carsBrandModelList)
	if err != nil {
		log.Println(err)

		return
	}
}

func carHandle(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "text/plain")

	_, err := w.Write([]byte(getCar(carID)))
	if err != nil {
		log.Println(err)

		return
	}
}

func main() {
	createCarsList()

	router := chi.NewRouter()

	router.Route("/cars", func(router chi.Router) {
		router.Get("/", carsHandle)
		router.Route("/{brand}", func(router chi.Router) {
			router.Get("/", carsBrandHandle)
			router.Get("/{model}", carsBrandModelHandle)
		})
	})
	router.Get("/car/{id}", carHandle)

	log.Fatal(http.ListenAndServe("localhost:8080", router)) //nolint:gosec // example
}
