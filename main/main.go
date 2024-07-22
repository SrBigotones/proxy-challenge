package main

import (
	"net/http"

	"github.com/SrBigotones/proxy-challenge/controllers/proxy"
)

// "github.com/SrBigotones/proxy-challenge/controllers/proxy"

func main() {

	println("Starting server")

	// r := http.NewServeMux()

	newRouter := proxy.CreateRouter()

	// r := proxy.CreateRouter()

	// println("Starting server")

	http.ListenAndServe(":8080", newRouter)
}
