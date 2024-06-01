package main

import (
	"fmt"
	"log"

	"mocks/internal/core/application"
	"mocks/internal/core/services"
	"mocks/internal/infra/api/rest"
	"mocks/internal/infra/store"
	"mocks/internal/infra/store/memory"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	store, err := store.NewStore(store.Config{
		Memory: &memory.Config{},
	})
	if err != nil {
		return fmt.Errorf("failed to initialize a store: %w", err)
	}
	repo := services.NewRepo(store)
	foobar := application.NewApplication(repo)
	api := rest.NewAPI(foobar)
	if err := api.Run(); err != nil {
		return fmt.Errorf("server has failed: %w", err)
	}
	return nil
}
