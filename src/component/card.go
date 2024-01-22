package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Card struct {
	app.Compo
}

func (c *Card) Render() app.UI {

	return app.Div().Body(
		app.Text("this is the card"),
		&Hello{
			name: "some card",
		},
	)
}
