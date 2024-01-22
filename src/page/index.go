package page

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Index struct {
	app.Compo
}

func (p *Index) Render() app.UI {
	return app.Body()
}
