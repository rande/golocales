package main

import "fmt"

type Number struct {
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
}

func AttachNumberSymbols(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var defaultNumber *Number
	var defaultSystem = "latn"

	// 1 - we need to find the default number system: latn is the default number system
	if locale.IsRoot { // nothing is yet loaded correctly, so we need to find the default values
		for _, t := range ldml.Numbers.Symbols {
			if t.NumberSystem == defaultSystem {
				defaultNumber = &Number{
					System:                 t.NumberSystem,
					MinusSign:              t.MinusSign,
					PlusSign:               t.PlusSign,
					Exponential:            t.Exponential,
					SuperscriptingExponent: t.SuperscriptingExponent,
					Decimal:                t.Decimal,
					Group:                  t.Group,
					PercentSign:            t.PercentSign,
					ApproximatelySign:      t.ApproximatelySign,
					Infinity:               t.Infinity,
					TimeSeparator:          t.TimeSeparator,
				}
			}
		}

	} else if locale.Parent != nil && locale.Parent.Numbers[defaultSystem] != nil {
		// in a non root locale, we attach the default values from the parent
		defaultNumber = locale.Parent.Numbers[defaultSystem]
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

		locale.Numbers[t.NumberSystem] = &Number{
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
		}
	}
}