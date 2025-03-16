package routes

import (
	"wasm/app/drivers"
	"wasm/app/templates"

	"github.com/brightsidedeveloper/goat"
	"github.com/brightsidedeveloper/goat/goatRouter"
)

var Router *goatRouter.Router

func RouterProvider() {

	r := goatRouter.NewRouter()

	r.Register("/", func(params map[string]string) {
		goat.RenderRoot(drivers.App())
	})

	r.Register("/page", func(params map[string]string) {
		goat.RenderRoot(templates.Page())
	})

	r.Register("/404", func(params map[string]string) {
		goat.RenderRoot(templates.NotFound())
	})

	r.SetupEventListeners()

	Router = r
}
