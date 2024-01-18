package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type topNavigation struct {
	app.Compo
}

func NewTopNavigation() *topNavigation {
	return &topNavigation{}
}

func (t *topNavigation) Render() app.UI {
	return app.Div().Body(
		app.Nav().Body(
			app.A().Href("/").Body(
				t.logo(),
			),

			app.H1().Text("Docker Card").Class("Header-title"),

			app.Div().Body(
				t.menu(
					app.Li().Body(
						app.A().Href("/login").Text("Login"),
					).Class("Header-login-button"),
				),
			).Class("Header-right-content"),
		).Class("Header-nav"),
	)
}

func (t *topNavigation) logo() app.UI {
	return app.Figure().Body(
		app.Img().Src("/web/icon/logo.png").Alt("logo").Class("Header-logo"),
	)
}

func (t *topNavigation) menu(content ...app.UI) app.UI {
	return app.Ul().Body(
		content...,
	).Class("Header-menu")
}
