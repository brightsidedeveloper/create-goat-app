package components

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/brightsidedeveloper/goat"
)

func App(ctx context.Context, props any) goat.GoatNode {

	return goat.GoatNode{
		Tag: "div",
		Attrs: map[string]string{
			"style": "display:flex;justify-content:center;align-items:center;gap:10px;flex-direction:column;padding:20px;",
		},
		Children: []goat.GoatNode{
			Counter(ctx, CounterProps{
				InitialCount: 0,
			}),
			Counter(ctx, CounterProps{
				InitialCount: 5,
			}),
		},
	}
}

type CounterProps struct {
	InitialCount int
}

func Counter(ctx context.Context, props CounterProps) goat.GoatNode {
	getCount, setCount := goat.UseState(ctx, props.InitialCount)

	increment := js.FuncOf(func(this js.Value, args []js.Value) any {
		setCount(getCount() + 1)
		return nil
	})

	return goat.NewGoatNode("div", nil, nil, []goat.GoatNode{
		{Text: fmt.Sprintf("Count: %d", getCount())},
		{
			Tag: "button",
			Attrs: map[string]string{
				"style": "margin-left: 10px",
			},
			Events: map[string]js.Func{
				"click": increment,
			},
			Text: "Click me",
		},
	}, "")
}
