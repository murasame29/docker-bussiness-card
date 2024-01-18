package svgs

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// CheckMark is a component that renders a check mark.
type CheckMark struct {
	app.Compo
}

// NewCheckMark creates a new CheckMark.
func NewCheckMark() *CheckMark {
	return &CheckMark{}
}

// Render renders the CheckMark.
func (*CheckMark) Render(opts ...Option) app.UI {
	o := &options{
		width:   24,
		height:  24,
		color:   "#fff",
		viewBox: "0 0 24 24",
		opacity: 1,
	}
	for _, opt := range opts {
		opt(o)
	}

	return app.Raw(fmt.Sprintf(
		`
	<svg width="%d" height="%d" viewBox="%s" class="%s" xmlns="http://www.w3.org/2000/svg">
		<path fill="%s" fill-opacity="%f" d="M10.8519 0.52594L3.89189 7.10404L1.14811 4.51081L0 5.59592L3.89189 9.27426L12 1.61105L10.8519 0.52594Z"/>
	</svg>
	`, o.width, o.height, o.viewBox, strings.Join(o.classes, " "), o.color, o.opacity,
	))
}
