package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	const maxSymbolCnt = 1
	symbolCnt := make(map[rune]int)
	var builder strings.Builder
	var lastRune rune

	for _, symbol := range str {
		isDigit := unicode.IsDigit(symbol)

		if lastRune == 0 && isDigit {
			return "", ErrInvalidString
		}

		symbolCnt[symbol]++
		if symbolCnt[symbol] > maxSymbolCnt {
			return "", ErrInvalidString
		}

		if isDigit {
			repeatNumber, _ := strconv.Atoi(string(symbol))
			repeatNumber--
			builder.WriteString(strings.Repeat(string(lastRune), repeatNumber))
		} else {
			builder.WriteRune(symbol)
		}

		lastRune = symbol
	}

	return builder.String(), nil
}
