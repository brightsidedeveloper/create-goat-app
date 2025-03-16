package main

import (
	"wasm/app/components"

	"github.com/brightsidedeveloper/goat"
)

func main() {
	goat.RenderRoot("root", components.App, nil)
}
