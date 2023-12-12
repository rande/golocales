// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dto

import "time"

type Territory struct {
	Name    string
	Numeric string
	Alpha3  string
}

type Currency struct {
	Symbol       string
	Name         string
	Digits       uint8
	Rounding     uint8
	CashDigits   uint8
	CashRounding uint8
	Numeric      string
}

type CalendarFormatter func(t time.Time, timeZone string) string

type Calendar struct {
	System     string
	Labels     map[string][]string
	Formatters map[string]CalendarFormatter
}

type Number struct {
	Symbols               map[string]*Symbol
	Decimals              map[string]FormatGroup
	Currencies            map[string]FormatGroup
	DefaultNumberSystem   string
	MinimumGroupingDigits int
}

type TimeZone string
type NumberSystem string
type Locale struct {
	Name        string
	Territories map[string]*Territory
	Currencies  map[string]*Currency
	TimeZones   map[string]TimeZone
	Calendars   map[string]*Calendar
	Parent      *Locale
	Number      *Number
}

type Symbol struct {
	System                 string
	MinusSign              string
	PlusSign               string
	Exponential            string
	SuperscriptingExponent string
	Decimal                string
	Group                  string
	CurrencyGroup          string
	PercentSign            string
	ApproximatelySign      string
	Infinity               string
	TimeSeparator          string
	PerMilleSign           string
}

type FormatGroup map[string][]*NumberFormat

type NumberFormat struct {
	Type                  string
	Count                 string
	Pattern               string
	Alt                   string
	PrimaryGroupingSize   int
	SecondaryGroupingSize int
	StandardPattern       string
}

func (locale *Locale) GetSymbol(system string) *Symbol {
	if symbol, ok := locale.Number.Symbols[system]; ok {
		return symbol
	}

	if locale.Parent != nil {
		return locale.Parent.GetSymbol(system)
	}

	return nil
}

func (locale *Locale) String() string {
	return locale.Name
}

func (locale *Locale) GetDecimalFormats(system, name string) []*NumberFormat {
	if systemFormat, ok := locale.Number.Decimals[system]; ok {
		if format, ok := systemFormat[name]; ok {
			if len(format) > 0 {
				return format
			}
		}
	}

	if locale.Parent != nil {
		return locale.Parent.GetDecimalFormats(system, name)
	}

	return nil
}

func (locale *Locale) GetCurrencyFormats(system, name string) []*NumberFormat {
	if systemFormat, ok := locale.Number.Currencies[system]; ok {
		if format, ok := systemFormat[name]; ok {
			if len(format) > 0 {
				return format
			}
		}
	}

	if locale.Parent != nil {
		return locale.Parent.GetCurrencyFormats(system, name)
	}

	return nil
}

func (locale *Locale) GetCalendarLabels(system, name string) []string {
	if calendar, ok := locale.Calendars[system]; ok {
		if labels, ok := calendar.Labels[name]; ok {
			if len(labels) > 0 {
				return labels
			}
		}
	}

	if locale.Parent != nil {
		return locale.Parent.GetCalendarLabels(system, name)
	}

	return nil
}
