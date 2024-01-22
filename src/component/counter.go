package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Counter struct {
	app.Compo
}

func (c *Counter) OnMount(ctx app.Context) {
	ctx.SetState("count", 0)
}

func (c *Counter) Render() app.UI {
	return app.Div().Body(
		&CounterDisplay{},
		&CounterButton{},
	)
}
