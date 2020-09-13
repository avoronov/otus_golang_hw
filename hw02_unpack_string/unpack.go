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
	var escape bool
	var prev rune
	var result strings.Builder

	for _, curr := range input {
		if prev == 0 {
			prev = curr
			continue
		}

		if unicode.IsDigit(prev) {
			if !escape {
				return "", ErrInvalidString
			}
			escape = false
		}

		if unicode.IsDigit(curr) {
			if escape {
				prev = curr
				continue
			}

			repeatCount, err := strconv.Atoi(string(curr))
			if err != nil {
				return "", err
			}

			tmp := strings.Repeat(string(prev), repeatCount)
			result.WriteString(tmp)

			prev = 0
			continue
		}

		if curr == '\\' {
			if escape {
				escape = false
				continue
			}

			escape = true
		}

		result.WriteRune(prev)
		prev = curr
	}

	if prev != 0 {
		if unicode.IsDigit(prev) && !escape {
			return "", ErrInvalidString
		}
		if escape && prev == '\\' {
			return "", ErrInvalidString
		}
		result.WriteRune(prev)
	}

	return result.String(), nil
}
