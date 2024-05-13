package main

import (
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	client := resty.New()

	_, err := client.R().SetResult(&users).Get(url)
	if err != nil {
		return
	}

	userNames := make([]string, 0, len(users))

	for _, user := range users {
		userNames = append(userNames, user.Username)
	}

	log.Println(strings.Join(userNames, " "))
}
