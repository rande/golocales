// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import "strings"

func AttachValidity(cldr *CLDR, supplemental *SupplementalData) {
	for _, id := range supplemental.IdValidity.ID {
		v := &Validity{
			From:   id.Type,
			Status: id.IdStatus,
			List:   ParseValidityValues(id.Text),
		}

		cldr.Validities = append(cldr.Validities, v)
	}
}

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
