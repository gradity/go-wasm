package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type CounterButtonSelf struct {
	app.Compo
	count int
}

func (cb *CounterButtonSelf) OnMount(ctx app.Context) {
	cb.count = 0
}

func (cb *CounterButtonSelf) Render() app.UI {

	return app.Div().Body(
		app.H1().Text("Self Counter Button"),
		app.P().Text(cb.count),
		app.Button().Text("+").OnClick(cb.onClickPlus),
		app.Button().Text("-").OnClick(cb.onClickMinus),
	)
}

func (cb *CounterButtonSelf) onClickPlus(ctx app.Context, e app.Event) {
	cb.count++
}

func (cb *CounterButtonSelf) onClickMinus(ctx app.Context, e app.Event) {
	cb.count--
}
