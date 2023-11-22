// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

const DefaultNumberSystem = "latn"

type Number struct {
	Symbols               map[string]*Symbol
	Decimals              map[string]FormatGroup
	Currencies            map[string]FormatGroup // numbering system => format => format
	DefaultNumberSystem   string
	MinimumGroupingDigits int
}

// share with currency, number and percent
type NumberFormat struct {
	Type                  string
	Count                 string
	Pattern               string
	Alt                   string
	PrimaryGroupingSize   int
	SecondaryGroupingSize int
	StandardPattern       string
}

func processPattern(pattern string) string {
	// Strip the grouping info.
	pattern = strings.ReplaceAll(pattern, "#,##,##", "")
	pattern = strings.ReplaceAll(pattern, "#,##", "")
	pattern = strings.ReplaceAll(pattern, "#", "0")

	return pattern
}

func AttachNumber(locale *Locale, cldr *CLDR, ldml *Ldml) {

	locale.Number.MinimumGroupingDigits = ifEmptyInt(ldml.Numbers.MinimumGroupingDigits, "1")
	locale.Number.DefaultNumberSystem = ldml.Numbers.DefaultNumberingSystem

	// Parent must have a valid configuration
	if locale.Parent != nil {
		locale.Number.DefaultNumberSystem = locale.Parent.Number.DefaultNumberSystem
	}

	AttachNumberSymbols(locale, cldr, ldml)
	AttachNumberDecimals(locale, cldr, ldml)
	AttachNumberCurrencies(locale, cldr, ldml)
}

// This function an adaptation of https://github.com/bojanz/currency
// All credits goes to Bojan Zivanovic and contributors
func AttachPattern(format *NumberFormat) {
	if !strings.Contains(format.Pattern, "#") {
		return
	}

	format.PrimaryGroupingSize = 0
	format.SecondaryGroupingSize = 0

	patternParts := strings.Split(format.Pattern, ";")
	if strings.Contains(patternParts[0], ",") {
		r, _ := regexp.Compile("#+0")
		primaryGroup := r.FindString(patternParts[0])
		format.PrimaryGroupingSize = len(primaryGroup)
		format.SecondaryGroupingSize = format.PrimaryGroupingSize
		numberGroups := strings.Split(patternParts[0], ",")
		if len(numberGroups) > 2 {
			// This pattern has a distinct secondary group size.
			format.SecondaryGroupingSize = len(numberGroups[1])
		}
	}

	// Strip the grouping info from the patterns, now that it is available separately.
	format.StandardPattern = processPattern(format.Pattern)
}
