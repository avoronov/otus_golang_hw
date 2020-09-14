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

	wordFreq := map[string]int{}
	for _, word := range strings.Split(input, " ") {
		wordFreq[word]++
	}

	type word struct {
		Name string
		Freq int
	}

	words := []word{}
	for name, freq := range wordFreq {
		words = append(words, word{name, freq})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].Freq > words[j].Freq })

	var size int
	if len(words) > maxSize {
		size = maxSize
	} else {
		size = len(words)
	}

	result := []string{}
	for _, word := range words[:size] {
		result = append(result, word.Name)
	}

	return result
}
