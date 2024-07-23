package main

import "github.com/SrBigotones/proxy-challenge/cmd/api"

func main() {
	api := api.NewApiServer("", "8080")
	api.Run()
}
