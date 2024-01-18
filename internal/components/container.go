package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/murasame29/docker-bussiness-card/internal/components/button"
	"github.com/murasame29/docker-bussiness-card/internal/components/svgs"
)

type container struct {
	app.Compo

	isEn bool
}

func NewContainer(isEn bool) *container {
	return &container{
		isEn: isEn,
	}
}

func (c *container) RenderEn() app.UI {
	card := newCard()
	return app.Div().Body(
		card.Render(),
		c.blurb(),
		c.buttons(),
		card.LoadScript(),
	).Class("container")
}

func (c *container) title() app.UI {
	if c.isEn {
		return app.H1().Text("Docker Container Terminal Bussiness Card")
	}
	return app.H1().Text("Docker Container Terminal 名刺")
}

func (c *container) blurb() app.UI {
	return app.Div().Body(
		c.title(),
		c.features(),
	).Class("hero-blurb")
}

func (c *container) features() app.UI {
	if c.isEn {
		return app.Ul().Body(
			c.checkMarkList("No need to install anything"),
			c.checkMarkList("Easy and fast? and very cool!"),
			c.checkMarkList("Infinite handouts!"),
			c.checkMarkList("There's a lot of romance in that!"),
		).Class("hero-features")
	}
	return app.Ul().Body(
		c.checkMarkList("何もインストールする必要がない"),
		c.checkMarkList("簡単で早い？そしてとてもクール！"),
		c.checkMarkList("無限の配布できる！"),
		c.checkMarkList("そこにはたくさんのロマンがある！"),
	).Class("hero-features")
}

func (c *container) checkMarkList(text string) app.UI {
	return app.Li().Body(
		svgs.NewCheckMark().Render(
			svgs.Width(12),
			svgs.Height(10),
			svgs.ViewBox("0 0 12 10"),
			svgs.Color("white"),
			svgs.Classes("hero-checkmark"),
		),
		app.Text(text),
	)
}

func (c *container) buttons() app.UI {
	b := button.NewButton()
	return app.Div().Body(
		b.ButtonA(
			"Get Started!",
			"/build",
			button.Color("black"),
			button.BGColor("#FDDD00"),
			button.Border("1px solid #FDDD00"),
			button.Classes("Primary"),
		),
		b.ButtonA(
			"View Source",
			"https://github.com/murasame29/docker-bussiness-card",
			button.Color("white"),
			button.BGColor("Transparent"),
			button.Border("1px solid white"),
			button.Classes("Secondary"),
		),
	).Class("hero-actions")
}
