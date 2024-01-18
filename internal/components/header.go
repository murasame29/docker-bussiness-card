package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type header struct {
	app.Compo
}

func NewHeader() *header {
	return &header{}
}

func (h *header) Render() app.UI {
	return app.Header().Body(
		NewTopNavigation().Render(),
	).Class("Header")
}
