// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

type FormatGroup struct {
	Default []*NumberFormat
	Long    []*NumberFormat
	Short   []*NumberFormat
}

type NumberFormat struct {
	Type                  string
	Count                 string
	Pattern               string
	PrimaryGroupingSize   int
	SecondaryGroupingSize int
	StandardPattern       string
}

func processPattern(pattern string) string {
	// Strip the grouping info.
	pattern = strings.ReplaceAll(pattern, "#,##,##", "")
	pattern = strings.ReplaceAll(pattern, "#,##", "")

	return pattern
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

func AttachNumberDecimals(locale *Locale, cldr *CLDR, ldml *Ldml) {
	for _, t := range ldml.Numbers.DecimalFormats {
		// no symbol is defined, so we skip
		if t.NumberSystem == "" {
			continue
		}

		// iterate over: short, long, and empty (ie: default)
		for _, f := range t.DecimalFormatLength {
			code := ifEmptyString(f.Type, "default")

			// this is an alias, pointing to the latn (to check by evaluating the path)
			// ignoring the path for now.
			if t.Alias.Source != "" {
				continue
			}

			locale.Decimals[t.NumberSystem] = &FormatGroup{}

			// a code is defined in the locale, so we need to override the default values
			switch code {
			case "long":
				locale.Decimals[t.NumberSystem].Long = []*NumberFormat{}
			case "short":
				locale.Decimals[t.NumberSystem].Short = []*NumberFormat{}
			case "default":
				locale.Decimals[t.NumberSystem].Default = []*NumberFormat{}
			}

			// then we iterate over the patterns
			for _, p := range f.DecimalFormat.Pattern {
				format := &NumberFormat{
					Type:    p.Type,
					Count:   p.Count,
					Pattern: p.Text,
				}

				AttachPattern(format)

				switch code {
				case "long":
					locale.Decimals[t.NumberSystem].Long = append(locale.Decimals[t.NumberSystem].Long, format)
				case "short":
					locale.Decimals[t.NumberSystem].Short = append(locale.Decimals[t.NumberSystem].Short, format)
				case "default":
					fmt.Printf("%s => %#v\n", locale.Code, format)
					locale.Decimals[t.NumberSystem].Default = append(locale.Decimals[t.NumberSystem].Default, format)
				}
			}
		}
	}
}
