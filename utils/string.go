package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	lu = "┏"
	ru = "┓"
	ld = "┗"
	rd = "┛"
	Vl = "┃"
	hl = "━"
)

func UpperLine(length int) string {
	var line string
	line += lu
	for i := 0; i < length; i++ {
		line += hl
	}
	return line + ru
}

func LowerLine(length int) string {
	var line string
	line += ld
	for i := 0; i < length; i++ {
		line += hl
	}
	return line + rd
}

func trimEscapeSquence(str string) string {
	const regexPattern = `\\x1b\[[0-9;]*[A-Za-z]`
	re := regexp.MustCompile(regexPattern)
	text := strconv.Quote(str)
	return re.ReplaceAllString(text, "")
}

func strLength(str string) int {
	length := 0
	for _, runeValue := range str {
		if utf8.RuneLen(runeValue) == 4 {
			length += 1
		}
		if utf8.RuneLen(runeValue) == 3 {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

func MaxLength(slice []string) (maxLength int) {
	for _, str := range slice {
		str = RemoveEscapeSequences(trimEscapeSquence(str))
		if maxLength < strLength(str) {
			maxLength = strLength(str)
		}
	}
	return
}

func RemoveEscapeSequences(str string) string {
	re := regexp.MustCompile(`\\[abfnrtv\\'"]`)
	return strings.ReplaceAll(strings.ReplaceAll(re.ReplaceAllString(str, ""), `"`, ""), `\`, "")
}

func FillSpace(str string, length int) string {
	tstr := RemoveEscapeSequences(trimEscapeSquence(str))
	for i := 0; i < length-strLength(tstr); i++ {
		str += " "
	}
	return str
}
