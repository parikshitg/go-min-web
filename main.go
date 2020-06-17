package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(Login))
	http.Handle("/register", http.HandlerFunc(Register))
	http.Handle("/dashboard", http.HandlerFunc(Dashboard))

	http.ListenAndServe(":8080", nil)
}
