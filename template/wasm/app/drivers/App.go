package drivers

import (
	"syscall/js"
	"wasm/app/api"
	"wasm/app/templates"

	"github.com/brightsidedeveloper/goat"
)

var message string

func App() goat.TemplJoint {

	msg, err := api.GetJoke()

	if msg != "" {
		goat.Log("Also, the server wanted you to know: ", msg)
	}

	goat.JSFunc("goat", func(this js.Value, args []js.Value) any {
		message = "You are GOATED! 🐐"
		goat.Render("App", templates.App(message, nil))
		return nil
	})

	message = msg

	return templates.App(message, err)
}
