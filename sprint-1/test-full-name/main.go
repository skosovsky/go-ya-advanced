package main

import (
	"log"
)

// User — пользователь в системе.
type User struct {
	FirstName string
	LastName  string
}

// FullName возвращает имя и фамилию пользователя.
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := User{
		FirstName: "Misha",
		LastName:  "Popov",
	}

	log.Println(u.FullName())
}
