package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// UserViewHandler — хендлер, который нужно протестировать.
func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id is empty", http.StatusBadRequest)

			return
		}

		user, ok := users[userID]
		if !ok {
			http.Error(w, "user not found", http.StatusNotFound)

			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "can't provide a json. internal error", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonUser)
		if err != nil {
			return
		}
	}
}

// User — основной объект для теста.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	users := make(map[string]User)
	user1 := User{
		ID:        "user1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	user2 := User{
		ID:        "user2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["user1"] = user1
	users["user2"] = user2

	http.HandleFunc("/users/{user_id}/", UserViewHandler(users))
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //nolint:gosec //example
}
