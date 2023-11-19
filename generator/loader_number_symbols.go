// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type Symbol struct {
	MinimumGroupingDigits  int
	System                 string
	MinusSign              string
	PlusSign               string
	Exponential            string
	SuperscriptingExponent string
	Decimal                string
	Group                  string
	PercentSign            string
	ApproximatelySign      string
	Infinity               string
	TimeSeparator          string
	PerMilleSign           string
}

func AttachNumberSymbols(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var defaultNumber *Symbol

	// 1 - we need to find the default number system: latn is the default number system
	if locale.IsRoot { // nothing is yet loaded correctly, so we need to find the default values
		for _, t := range ldml.Numbers.Symbols {
			if t.NumberSystem == locale.Number.DefaultNumberSystem {
				defaultNumber = &Symbol{
					System:                 t.NumberSystem,           // latn
					MinusSign:              t.MinusSign,              // -
					PlusSign:               t.PlusSign,               // +
					Exponential:            t.Exponential,            // E
					SuperscriptingExponent: t.SuperscriptingExponent, // ×
					Decimal:                t.Decimal,                // ,
					Group:                  t.Group,                  // ,
					PercentSign:            t.PercentSign,            // %
					ApproximatelySign:      t.ApproximatelySign,      // ~
					Infinity:               t.Infinity,               // ∞
					TimeSeparator:          t.TimeSeparator,          // :
					PerMilleSign:           t.PerMille,               // ‰
				}

				locale.Number.Symbols[t.NumberSystem] = defaultNumber
			}
		}

	} else if locale.Parent != nil && locale.Parent.Number.Symbols[locale.Number.DefaultNumberSystem] != nil {
		// in a non root locale, we attach the default values from the parent
		defaultNumber = locale.Parent.Number.Symbols[locale.Number.DefaultNumberSystem]
	} else {
		fmt.Printf("No default number system found for %s\n", locale.Code)
		return
	}

	// if any value are defined in the locale, there are overriding the default values
	for _, t := range ldml.Numbers.Symbols {
		// no symbol is defined, so we skip
		if t.NumberSystem == "" {
			continue
		}
		// the parent has the same default system, do we can skip the configuration
		if locale.Parent != nil && locale.Parent.Number.DefaultNumberSystem == t.NumberSystem {
			continue
		}

		number := &Symbol{
			System:                 ifEmptyString(t.NumberSystem, defaultNumber.System),
			MinusSign:              ifEmptyString(t.MinusSign, defaultNumber.MinusSign),
			PlusSign:               ifEmptyString(t.PlusSign, defaultNumber.PlusSign),
			Exponential:            ifEmptyString(t.Exponential, defaultNumber.Exponential),
			SuperscriptingExponent: ifEmptyString(t.SuperscriptingExponent, defaultNumber.SuperscriptingExponent),
			Decimal:                ifEmptyString(t.Decimal, defaultNumber.Decimal),
			Group:                  ifEmptyString(t.Group, defaultNumber.Group),
			PercentSign:            ifEmptyString(t.PercentSign, defaultNumber.PercentSign),
			ApproximatelySign:      ifEmptyString(t.ApproximatelySign, defaultNumber.ApproximatelySign),
			Infinity:               ifEmptyString(t.Infinity, defaultNumber.Infinity),
			TimeSeparator:          ifEmptyString(t.TimeSeparator, defaultNumber.TimeSeparator),
			PerMilleSign:           ifEmptyString(t.PerMille, defaultNumber.PerMilleSign),
		}

		// don't store if the same as the parent,
		// we need to check the hierarchie too
		// @TODO: implements inheritance checks
		if number.MinusSign == defaultNumber.MinusSign &&
			number.PlusSign == defaultNumber.PlusSign &&
			number.Exponential == defaultNumber.Exponential &&
			number.SuperscriptingExponent == defaultNumber.SuperscriptingExponent &&
			number.Decimal == defaultNumber.Decimal &&
			number.Group == defaultNumber.Group &&
			number.PercentSign == defaultNumber.PercentSign &&
			number.ApproximatelySign == defaultNumber.ApproximatelySign &&
			number.Infinity == defaultNumber.Infinity &&
			number.TimeSeparator == defaultNumber.TimeSeparator &&
			number.PerMilleSign == defaultNumber.PerMilleSign {
			continue
		}

		locale.Number.Symbols[t.NumberSystem] = number
	}
}
