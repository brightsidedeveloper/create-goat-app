package components

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/brightsidedeveloper/goat"
	"github.com/brightsidedeveloper/goat/el"
)

func App(ctx context.Context, props any) goat.GoatNode {

	return el.Div().
		Attr("style", "display:flex;justify-content:center;align-items:center;gap:10px;flex-direction:column;padding:20px;").
		Child(Counter(ctx, CounterProps{InitialCount: 0})).
		Child(Counter(ctx, CounterProps{InitialCount: 10})).
		Build()
}

type CounterProps struct {
	InitialCount int
}

func Counter(ctx context.Context, props CounterProps) goat.GoatNode {
	getCount, setCount := goat.UseState(ctx, props.InitialCount)

	increment := goat.Fn(func(this js.Value, args []js.Value) any {
		setCount(getCount() + 1)
		return nil
	})

	return el.Div().
		Child(el.Text(fmt.Sprintf("Count: %d", getCount()))).
		Child(el.Button().
			Attr("style", "margin-left: 10px").
			Event("click", increment).
			Text("Click me").Build()).
		Build()

}
