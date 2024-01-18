package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IndexLayout struct {
	app.Compo
}

func NewIndex() *IndexLayout {
	return &IndexLayout{}
}

func (l *IndexLayout) Render() app.UI {
	return app.Div().Body(
		NewHeader().Render(),
		NewContainer(
			true,
		).RenderEn(),
	)
}
