// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package golocales

type Territory string
type Currency string
type TimeZone string
type NumberSystem string
type Locale struct {
	Name                  string
	Territories           map[string]Territory
	Currencies            map[string]Currency
	TimeZones             map[string]TimeZone
	Symbols               map[NumberSystem]*Symbol
	Decimal               map[NumberSystem]FormatGroup
	Parent                *Locale
	MinimumGroupingDigits int
	DefaultNumberSystem   string
}

type Symbol struct {
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
