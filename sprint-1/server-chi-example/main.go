package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func carsHandle(_ http.ResponseWriter, _ *http.Request) {
}

func brandHandle(_ http.ResponseWriter, _ *http.Request) {
}

func modelHandle(_ http.ResponseWriter, _ *http.Request) {
}

func RouteFunc(router chi.Router) {
	router.Get("/", carsHandle) // GET /cars

	// Route можно вкладывать один в другой
	router.Route("/{brand}", func(router chi.Router) {
		router.Get("/", brandHandle)        // GET /cars/renault
		router.Get("/{model}", modelHandle) // GET /cars/renault/duster
	})
}

func example() {
	router := chi.NewRouter()

	router.Get("/cars", carsHandle)                  // GET /cars
	router.Get("/cars/{brand}", brandHandle)         // GET /cars/renault
	router.Get("/cars/{brand}/{model}", modelHandle) // GET /cars/renault/duster

	// то же самое можно описать, используя Route
	router.Route("/cars", func(router chi.Router) {
		router.Get("/", carsHandle) // GET /cars

		// Route можно вкладывать один в другой
		router.Route("/{brand}", func(router chi.Router) {
			router.Get("/", brandHandle)        // GET /cars/renault
			router.Get("/{model}", modelHandle) // GET /cars/renault/duster
		})
	})

	// или выделив функцию роутинга отдельно
	router.Route("/cars", RouteFunc)
}

func newCar(_ http.ResponseWriter, _ *http.Request) {
}

func getCar(_ http.ResponseWriter, _ *http.Request) {
}

func updateCar(_ http.ResponseWriter, _ *http.Request) {
}

func deleteCar(_ http.ResponseWriter, _ *http.Request) {
}

func example2() {
	router := chi.NewRouter()

	router.Post("/car", newCar)           // POST /car
	router.Get("/car/{id}", getCar)       // GET /car/1234
	router.Put("/car/{id}", updateCar)    // PUT /car/1234
	router.Delete("/car/{id}", deleteCar) // DELETE /car/1234

	// то же самое, используя Router
	router.Route("/car", func(router chi.Router) {
		router.Post("/", newCar) // POST /car

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", getCar)       // GET /car/1234
			router.Put("/", updateCar)    // PUT /car/1234
			router.Delete("/", deleteCar) // DELETE /car/1234
		})
	})
}

func main() {
	example()
	example2()
}
