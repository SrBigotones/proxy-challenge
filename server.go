package main

import (
	"fmt"
	"net/http"
)

const limitPerIP = 1000
const limitCat = 10_000
const limitItems = 10

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /*", doThings)

	http.ListenAndServe(":8080", mux)
}

func doThings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New connection")

	response := fmt.Sprintf("New request from %s", r.RemoteAddr)

	fmt.Fprintf(w, response)
}
