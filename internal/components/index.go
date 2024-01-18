package components

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IndexLayout struct {
	app.Compo
}

func NewIndex() *IndexLayout {
	return &IndexLayout{}
}

func (l *IndexLayout) Render() app.UI {
	terminal := newCard()
	return app.Div().Body(
		NewHeader().Render(),
		terminal.Render(),
		app.Div().Body(
			l.buttons(),
		),
		terminal.LoadScript(),
	)
}

func (l *IndexLayout) buttons() app.UI {
	return app.Div().Body(
		l.startButton(),
		l.ViewSource(),
	).Class("hero-actions")
}

func (l *IndexLayout) startButton() app.UI {
	return app.A().Class("Primary").Href("/build").Text("Get Started!").Role("button")
}

func (l *IndexLayout) ViewSource() app.UI {
	return app.A().Class("Secondary").Href("https://github.com/murasame29").Text("View Source").Role("button")
}

func (l *IndexLayout) start(ctx app.Context, e app.Event) {
	log.Println("start")
	ctx.Navigate("/build")
}
