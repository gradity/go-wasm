package handler

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func HandleCount(ctx app.Context, a app.Action) {

	var count int
	ctx.GetState("count", &count)

	switch a.Tags.Get("operation") {
	case "add":
		count++
	case "substract":
		count--
	case "reset":
		count = 0
	}

	ctx.SetState("count", count)
}
