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
		if unicode.IsDigit(curr) {
			if !escape && prev == 0 {
				return "", ErrInvalidString
			}

			if escape {
				prev = curr
				escape = false
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
				prev = '\\'
				escape = false
				continue
			}

			if prev != 0 {
				result.WriteRune(prev)
			}
			prev = 0

			escape = true
			continue
		}

		if prev != 0 {
			result.WriteRune(prev)
		}

		prev = curr
	}

	if escape {
		return "", ErrInvalidString
	}

	if prev != 0 {
		result.WriteRune(prev)
	}

	return result.String(), nil
}
