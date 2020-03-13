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

	_, err := strconv.Atoi(str)
	if err == nil {
		return "", ErrInvalidString
	}

	symbolCnt := make(map[rune]int)
	var builder strings.Builder
	var lastRune rune

	for _, symbol := range str {
		symbolCnt[symbol]++

		if symbolCnt[symbol] > 1 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(symbol) {
			repeatNumber, _ := strconv.Atoi(string(symbol))
			builder.WriteString(strings.Repeat(string(lastRune), repeatNumber-1))
		} else {
			builder.WriteRune(symbol)
		}

		lastRune = symbol
	}

	return builder.String(), nil
}
