package main

import (
	"net/http"

	"github.com/SrBigotones/proxy-challenge/controllers/proxy"
	"github.com/SrBigotones/proxy-challenge/controllers/stats"
	"github.com/gorilla/mux"
)

func main() {

	println("Starting server")

	r := mux.NewRouter()
	stats.RegisterRouter(r)
	proxy.RegisterRouter(r)

	http.ListenAndServe(":8080", r)
}
