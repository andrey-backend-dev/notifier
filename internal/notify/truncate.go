package notify

import (
	"fmt"
	"unicode/utf8"
)

const ellipsis = "…"

func Truncate(input string, maxSymbols int) (string, error) {
	switch {
	case maxSymbols <= 0:
		return "", fmt.Errorf("cannot truncate string when maxSymbols <= 0")
	case utf8.RuneCountInString(input) <= maxSymbols:
		return input, nil
	default:
		return string([]rune(input)[:maxSymbols-1]) + ellipsis, nil
	}
}
