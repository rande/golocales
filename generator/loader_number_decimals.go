// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

type FormatGroup struct {
	Default []*NumberFormat
	Long    []*NumberFormat
	Short   []*NumberFormat
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

			// a code is defined in the locale, so we need to override the default values
			if locale.Number.Decimals[t.NumberSystem] == nil {
				locale.Number.Decimals[t.NumberSystem] = &FormatGroup{
					Long:    []*NumberFormat{},
					Short:   []*NumberFormat{},
					Default: []*NumberFormat{},
				}
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
					locale.Number.Decimals[t.NumberSystem].Long = append(locale.Number.Decimals[t.NumberSystem].Long, format)
				case "short":
					locale.Number.Decimals[t.NumberSystem].Short = append(locale.Number.Decimals[t.NumberSystem].Short, format)
				case "default":
					locale.Number.Decimals[t.NumberSystem].Default = append(locale.Number.Decimals[t.NumberSystem].Default, format)
				}
			}
		}
	}
}
