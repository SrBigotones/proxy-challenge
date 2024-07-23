package main

import "github.com/SrBigotones/proxy-challenge/cmd/api"

func main() {
	api := api.NewApiServer("", "3000")
	api.Run()
}
