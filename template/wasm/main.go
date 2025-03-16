package main

import (
	"wasm/app/routes"

	"github.com/brightsidedeveloper/goat"
)

func main() {
	done := make(chan struct{}, 0)

	goat.Log("Hi from WASM!")

	routes.RouterProvider()

	<-done
}
