package main

import (
	"log"
	"os"
)

func main() {
	user := os.Getenv("USER")
	log.Println(user)

	user, ok := os.LookupEnv("SHELL")
	log.Println(user, ok)

	envList := os.Environ()

	for _, env := range envList {
		log.Println(env)
	}
}
