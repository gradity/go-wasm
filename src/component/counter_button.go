package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type CounterButton struct {
	app.Compo
}

func (cb *CounterButton) Render() app.UI {

	return app.Div().Body(
		app.Button().Text("add").Value("add").OnClick(cb.onClickOperation),
		app.Button().Text("reset").Value("reset").OnClick(cb.onClickOperation),
		app.Button().Text("subtract").Value("substract").OnClick(cb.onClickOperation),
	)
}

func (cb *CounterButton) onClickOperation(ctx app.Context, e app.Event) {

	operation := ctx.JSSrc().Get("value").String()
	ctx.NewAction("count", app.T("operation", operation))
}
