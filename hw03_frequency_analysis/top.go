package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

const maxSize = 10

func sanitize(input *string) *string {
	inp := *input

	inp = strings.ToLower(inp)

	specSpaceRe := regexp.MustCompile(`[\t\r\n\v\f]`)
	inp = specSpaceRe.ReplaceAllLiteralString(inp, " ")

	spaceRe := regexp.MustCompile(`[[:space:]]{2,}`)
	inp = spaceRe.ReplaceAllLiteralString(inp, " ")

	hyphenRe := regexp.MustCompile(`(\pL+)-(\pL+)`)
	inp = hyphenRe.ReplaceAllString(inp, "$1   $2")

	punctRe := regexp.MustCompile(`[[:punct:]]`)
	inp = punctRe.ReplaceAllLiteralString(inp, "")

	hyphenBackRe := regexp.MustCompile(`(\pL+) {3}(\pL+)`)
	inp = hyphenBackRe.ReplaceAllString(inp, "$1-$2")

	inp = spaceRe.ReplaceAllLiteralString(inp, " ")

	return &inp
}

// Top10 return top 10 most seen words.
func Top10(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	inp := sanitize(&input)
	input = *inp

	frequencies := map[string]int{}
	for _, word := range strings.Split(input, " ") {
		frequencies[word]++
	}

	i := 0
	words := make([]string, len(frequencies))
	for word := range frequencies {
		words[i] = word
		i++
	}

	sort.Slice(words, func(i, j int) bool { return frequencies[words[i]] > frequencies[words[j]] })

	var size int
	if len(words) > maxSize {
		size = maxSize
	} else {
		size = len(words)
	}

	return words[:size]
}
