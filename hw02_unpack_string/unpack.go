package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder
	var lastRune rune

	for _, symbol := range str {
		if unicode.IsDigit(symbol) {
			if lastRune == 0 || unicode.IsDigit(lastRune) {
				return "", ErrInvalidString
			}

			repeatNumber, _ := strconv.Atoi(string(symbol))
			builder.WriteString(strings.Repeat(string(lastRune), repeatNumber-1)) //nolint:gomnd
		} else {
			builder.WriteRune(symbol)
		}

		lastRune = symbol
	}

	return builder.String(), nil
}
