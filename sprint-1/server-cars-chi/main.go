package main

import (
	"io"
	"log"
	"net/http"
	"strings"
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

func getCarsList() []string {
	var list = make([]string, 0, len(cars))
	for _, car := range cars {
		list = append(list, car)
	}

	return list
}

func getCar(id string) string {
	if car, ok := cars[id]; ok {
		return car
	}

	return "unknown " + id
}

func carsHandle(w http.ResponseWriter, _ *http.Request) {
	carsList := getCarsList()
	_, err := io.WriteString(w, strings.Join(carsList, ", "))
	if err != nil {
		log.Println(err)

		return
	}
}

func carHandle(w http.ResponseWriter, r *http.Request) {
	// carID := chi.URLParam(r, "id")

	// carID := r.PathValue("id")

	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "carID param is missing", http.StatusBadRequest)

		return
	}

	_, err := w.Write([]byte(getCar(carID)))
	if err != nil {
		log.Println(err)

		return
	}
}

func main() {
	// router := chi.NewRouter()
	// router.Get("/cars/", carsHandle)
	// router.Get("/car/{id}", carHandle)

	// router := http.NewServeMux()
	// router.HandleFunc("GET /cars", carsHandle)
	// router.HandleFunc("GET /car/{id}", carHandle)

	http.HandleFunc("/cars", carsHandle)
	http.HandleFunc("/car", carHandle)

	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //nolint:gosec // example
}
