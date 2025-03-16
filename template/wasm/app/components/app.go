package components

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/brightsidedeveloper/goat"
	"github.com/brightsidedeveloper/goat/el"
)

func App(ctx context.Context, props any) goat.GoatNode {

	div := el.Div()

	div.Attr("style", "display:flex;justify-content:center;align-items:center;gap:10px;flex-direction:column;padding:20px;")

	for i := range 3 {
		div.Child(Counter(ctx, CounterProps{InitialCount: i}))
	}

	return div.Build()

}

type CounterProps struct {
	InitialCount int
}

func Counter(ctx context.Context, props CounterProps) goat.GoatNode {
	getCount, setCount := goat.UseState(ctx, props.InitialCount)

	increment := goat.EventCB(func(el js.Value, event js.Value) {
		goat.Log(event)
		setCount(getCount() + 1)
	})

	return el.Div().
		Child(el.Text(fmt.Sprintf("Count: %d", getCount()))).
		Child(el.Button().
			Attr("style", "margin-left: 10px").
			Event("click", increment).
			Text("Click me").Build()).
		Build()

}
