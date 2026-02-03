package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server starts at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	path := r.URL.Path
	query := r.URL.Query().Get("id")

	fmt.Fprintf(w, "You sent a %s request to %s with ID %s\n",
		method, path, query)
}
