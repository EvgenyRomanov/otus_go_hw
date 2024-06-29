package hw02unpackstring

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

const BACKSLASH = '\\'

func Unpack(s string) (string, error) {
	runes := []rune(s)
	var result string
	var backslash bool

	for i, item := range runes {
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(runes[i-1]) && i > 2 && runes[i-2] != BACKSLASH {
			return "", ErrInvalidString
		}

		if item == BACKSLASH && !backslash {
			backslash = true
			continue
		}

		if backslash && unicode.IsLetter(item) {
			return "", ErrInvalidString
		}

		if backslash {
			result += string(item)
			backslash = false
			continue
		}

		if unicode.IsDigit(item) {
			repeat := int(item - '0')

			if repeat == 0 {
				result = result[:utf8.RuneCountInString(result)-1]
				continue
			}

			for j := 0; j < repeat-1; j++ {
				result += string(runes[i-1])
			}

			continue
		}

		result += string(item)
	}

	return result, nil
}
