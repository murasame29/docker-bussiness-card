package svgs

type options struct {
	classes []string
	width   int
	height  int
	viewBox string
	color   string
	opacity float64
}

// Option is a function that configures a svg.
type Option func(*options)

// Classes configures the svg classes.
func Classes(classes ...string) Option {
	return func(o *options) {
		o.classes = append(o.classes, classes...)
	}
}

// Width configures the svg width.
func Width(width int) Option {
	return func(o *options) {
		o.width = width
	}
}

// Height configures the svg height.
func Height(height int) Option {
	return func(o *options) {
		o.height = height
	}
}

// ViewBox configures the svg view box.
func ViewBox(viewBox string) Option {
	return func(o *options) {
		o.viewBox = viewBox
	}
}

// Color configures the svg color.
func Color(color string) Option {
	return func(o *options) {
		o.color = color
	}
}

// Opacity configures the svg opacity.
func Opacity(opacity float64) Option {
	return func(o *options) {
		o.opacity = opacity
	}
}
