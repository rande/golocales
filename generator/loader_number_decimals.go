package main

import "fmt"

type FormatGroup struct {
	Default []*NumberFormat
	Long    []*NumberFormat
	Short   []*NumberFormat
}

type NumberFormat struct {
	Type    string
	Count   string
	Pattern string
}

func AttachNumberDecimals(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var defaultDecimal *FormatGroup
	var defaultSystem = "latn"

	// 1 - we need to find the default number system: latn is the default number system
	if locale.IsRoot { // nothing is yet loaded correctly, so we need to find the default values
		for _, t := range ldml.Numbers.DecimalFormats {
			if t.NumberSystem != defaultSystem {
				continue
			}

			defaultDecimal = &FormatGroup{}

			// iterate over: short, long, and empty (ie: default)
			for _, f := range t.DecimalFormatLength {
				code := ifEmptyString(f.Type, "default")

				for _, p := range f.DecimalFormat.Pattern {
					format := &NumberFormat{Type: p.Type, Count: p.Count, Pattern: p.Text}

					switch code {
					case "long":
						defaultDecimal.Long = append(defaultDecimal.Long, format)
					case "short":
						defaultDecimal.Short = append(defaultDecimal.Short, format)
					case "default":
						defaultDecimal.Default = append(defaultDecimal.Default, format)
					}
				}
			}
		}

	} else if locale.Parent != nil && locale.Parent.Decimals[defaultSystem] != nil {
		// in a non root locale, we attach the default values from the parent
		defaultDecimal = locale.Parent.Decimals[defaultSystem]
	} else {
		fmt.Printf("No default decimal system found for %s\n", locale.Code)
		return
	}

	for _, t := range ldml.Numbers.DecimalFormats {

		// no symbol is defined, so we skip
		if t.NumberSystem == "" {
			continue
		}

		// if _, ok := locale.Numbers[t.NumberSystem]; !ok {
		// 	locale.Numbers[t.NumberSystem] =
		// }

		locale.Decimals[t.NumberSystem] = &FormatGroup{}

		locale.Decimals[t.NumberSystem].Long = defaultDecimal.Long
		if len(locale.Decimals[t.NumberSystem].Long) == 0 {
			locale.Decimals[t.NumberSystem].Long = defaultDecimal.Short
		}
		locale.Decimals[t.NumberSystem].Short = defaultDecimal.Short
		locale.Decimals[t.NumberSystem].Default = defaultDecimal.Default

		// iterate over: short, long, and empty (ie: default)
		for _, f := range t.DecimalFormatLength {
			code := ifEmptyString(f.Type, "default")

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

				switch code {
				case "long":
					locale.Decimals[t.NumberSystem].Long = append(locale.Decimals[t.NumberSystem].Long, format)
				case "short":
					locale.Decimals[t.NumberSystem].Short = append(locale.Decimals[t.NumberSystem].Short, format)
				case "default":
					locale.Decimals[t.NumberSystem].Default = append(locale.Decimals[t.NumberSystem].Default, format)
				}
			}
		}
	}
}
