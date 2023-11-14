package main

import (
	"strings"
)

func ParseValidityValues(message string) []string {
	partials := []string{}

	message = strings.ReplaceAll(message, "\n\t", " ")

	for _, word := range strings.Split(message, " ") {
		word = strings.TrimSpace(word)
		if word == "" || word == "\n" || word == "\t" {
			continue
		}

		if !strings.Contains(word, "~") {
			partials = append(partials, word)
			continue
		}

		parts := strings.Split(word, "~")
		from := parts[0]
		to := byte(parts[1][0])

		partials = append(partials, from)

		letter := byte(from[1])

		for letter <= to-1 {
			letter++
			partials = append(partials, string([]byte{from[1], letter}))
		}
	}

	return partials
}
