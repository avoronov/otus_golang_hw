package hw02_unpack_string // nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString is the error raised when invalid string is given.
var ErrInvalidString = errors.New("invalid string")

func doStep(curr rune, prev rune, result *strings.Builder) (rune, error) {
	if prev == 0 {
		return curr, nil
	}

	if unicode.IsDigit(curr) {
		if !unicode.IsLetter(prev) {
			return 0, ErrInvalidString
		}

		repeatCount, err := strconv.Atoi(string(curr))
		if err != nil {
			return 0, err
		}

		tmp := strings.Repeat(string(prev), repeatCount)
		result.WriteString(tmp)

		return 0, nil
	}

	if unicode.IsLetter(curr) {
		if !unicode.IsLetter(prev) {
			return 0, ErrInvalidString
		}

		result.WriteRune(prev)

		return curr, nil
	}

	if curr == 0 {
		if !unicode.IsLetter(prev) {
			return 0, ErrInvalidString
		}

		result.WriteRune(prev)

		return 0, nil
	}

	return 0, ErrInvalidString
}

// Unpack is the function which perform simple unpacking.
func Unpack(input string) (string, error) {
	var prev rune
	var err error
	var result strings.Builder

	for _, curr := range input {
		prev, err = doStep(curr, prev, &result)
		if err != nil {
			return "", err
		}
	}

	if prev != 0 {
		_, err = doStep(0, prev, &result)
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}
