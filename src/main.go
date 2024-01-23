package main

// Continue to read https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /login request\n")
	io.WriteString(w, "/login is unavailable!\n")
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /register request\n")
	io.WriteString(w, "/user/register is unavailable!\n")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /user/get request\n")
	io.WriteString(w, "/user/get is unavailable!\n")
}

func main() {
	http.HandleFunc("/login", getLogin)
	http.HandleFunc("/user/register", getRegister)
	http.HandleFunc("/user/get", getUser)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
