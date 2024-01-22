package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hello struct {
	app.Compo

	name string
}

func (h *Hello) Render() app.UI {
	if h.name != "" {
		h.name = app.Window().URL().Query().Get("name")
	}

	return app.If(
		h.name != "",
		app.H1().Textf("hello %s", h.name),
	).Else(
		app.H1().Text("hello world"),
	)

}
