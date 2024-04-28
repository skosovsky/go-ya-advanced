package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин <input type="text" name="login"></label>
            <label>Пароль <input type="password" name="password"></label>
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func isAuth(login, password string) bool {
	return login == "guest" && password == "demo"
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		login := req.FormValue("login")
		password := req.FormValue("password")

		if !isAuth(login, password) {
			http.Error(res, "incorrect login or password", http.StatusUnauthorized)
			return
		}

		_, err := io.WriteString(res, "Welcome!") // fmt.Fprint(res, "Welcome!") or res.Write([]byte("Welcome!"))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		_, err := io.WriteString(res, form)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(mainPage)) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
