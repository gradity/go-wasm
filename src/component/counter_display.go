package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type CounterDisplay struct {
	app.Compo

	count int
}

func (cd *CounterDisplay) OnMount(ctx app.Context) {
	ctx.ObserveState("count").Value(&cd.count)
}

func (cd *CounterDisplay) Render() app.UI {

	return app.P().Textf("Count: %v", cd.count)
}
