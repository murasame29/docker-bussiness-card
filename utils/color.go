package utils

import "fmt"

func FontColor(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func ResetFontColor() string {
	return "\x1b[0m"
}

func BackGroundColor(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}

func ResetBackGroundColor() string {
	return "\x1b[0m"
}

const (
	Reset         = "\x1b[0m"
	Bold          = "\x1b[1m"
	Thin          = "\x1b[2m"
	Italic        = "\x1b[3m"
	Underline     = "\x1b[4m"
	SlowBlink     = "\x1b[5m"
	RapidBlink    = "\x1b[6m" // fish, bashは対応していない？
	Reversed      = "\x1b[7m"
	Hidden        = "\x1b[8m"
	StrikeThrough = "\x1b[9m"

	Fraktur         = "\x1b[20m"
	DubbleUnderline = "\x1b[21m"
	NormalIntensity = "\x1b[22m"
	NoItalic        = "\x1b[23m"
	NoUnderline     = "\x1b[24m"
	NoBlink         = "\x1b[25m"
	NoReversed      = "\x1b[27m"
	Reveal          = "\x1b[28m"
	NoCrossedOut    = "\x1b[29m"
)

// Fonts 10 ~ 19
const (
	PrimaryFont = "\x1b[10m"
	AltFont1    = "\x1b[11m"
	AltFont2    = "\x1b[12m"
	AltFont3    = "\x1b[13m"
	AltFont4    = "\x1b[14m"
	AltFont5    = "\x1b[15m"
	AltFont6    = "\x1b[16m"
	AltFont7    = "\x1b[17m"
	AltFont8    = "\x1b[18m"
	AltFont9    = "\x1b[19m"
)
