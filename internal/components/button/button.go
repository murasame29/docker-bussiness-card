package button

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type button struct {
	app.Compo

	bgColor string
	color   string
	border  string
	Classes []string
}

// Button is the interface that describes a button component.
type Button interface {
	ButtonA(text, path string, opts ...Option) app.UI
	Button(text string, opts ...Option) app.UI
}

// Option is a function that configures a button.
type Option func(*button)

// NewButton creates a new button.
func NewButton(ops ...Option) Button {
	b := &button{
		color:   "#fff",
		bgColor: "#000",
		border:  "1px solid #000",
	}
	return b
}

// BGColor configures the button color.
// The default color is white.
// Use CSS color name or hex value. ex: "#fff" or "white" or rgb(255, 255, 255)
func BGColor(color string) Option {
	return func(b *button) {
		b.bgColor = color
	}
}

// Color configures the button font color.
// The default color is white.
// Use CSS color name or hex value. ex: "#fff" or "white" or rgb(255, 255, 255)
func Color(color string) Option {
	return func(b *button) {
		b.color = color
	}
}

// Border configures the button border.
func Border(border string) Option {
	return func(b *button) {
		b.border = border
	}
}

// Class configures the button class.
func Class(class string) Option {
	return func(b *button) {
		b.Classes = append(b.Classes, class)
	}
}

// Classes configures the button classes.
func Classes(classes ...string) Option {
	return func(b *button) {
		b.Classes = append(b.Classes, classes...)
	}
}

// ButtonA creates a new button for "a" tag.
func (l *button) ButtonA(text, path string, opts ...Option) app.UI {
	for _, opt := range opts {
		opt(l)
	}

	button := app.A().Href(path).Text(text).Role("button")
	button.Styles(map[string]string{
		"background-color": l.bgColor,
		"color":            l.color,
		"text-decoration":  "none",
		"border":           l.border,
	})
	button.Class(l.Classes...)
	return button
}

// Button create a new button for "button" tag.
func (l *button) Button(text string, opts ...Option) app.UI {
	for _, opt := range opts {
		opt(l)
	}

	button := app.Button().Text(text).Role("button")
	button.Styles(map[string]string{
		"background-color": l.bgColor,
		"color":            l.color,
		"border":           l.border,
	})
	button.Class(l.Classes...)

	return button
}
