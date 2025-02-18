package main

import (
	"fmt"

	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Dashboard")
	})

	fmt.Println("Server running on url: http://localhost:90")
	http.ListenAndServe(":90", nil)

}
