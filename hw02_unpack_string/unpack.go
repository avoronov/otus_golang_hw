package hw02_unpack_string // nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString is the error raised when invalid string is given.
var ErrInvalidString = errors.New("invalid string")

// Unpack is the function which perform simple unpacking.
func Unpack(input string) (string, error) {
	var prev rune
	var result strings.Builder

	for _, curr := range input {
		if prev == 0 {
			prev = curr
			continue
		}

		if unicode.IsDigit(prev) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(curr) {
			repeatCount, err := strconv.Atoi(string(curr))
			if err != nil {
				return "", err
			}

			tmp := strings.Repeat(string(prev), repeatCount)
			result.WriteString(tmp)

			prev = 0

			continue
		}

		if unicode.IsLetter(curr) {
			result.WriteRune(prev)
			prev = curr

			continue
		}

		return "", ErrInvalidString
	}

	if prev != 0 {
		if unicode.IsDigit(prev) {
			return "", ErrInvalidString
		}
		result.WriteRune(prev)
	}

	return result.String(), nil
}
