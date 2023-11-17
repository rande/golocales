// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

type Locale struct {
	IsRoot                bool
	IsBase                bool
	Parent                *Locale
	Code                  string
	Name                  string
	Territory             string
	Territories           map[string]Territory
	Currencies            map[string]Currency
	TimeZones             map[string]TimeZone
	Parents               []*Locale
	Symbols               map[string]*Symbol
	Decimals              map[string]*FormatGroup
	DefaultNumberSystem   string
	MinimumGroupingDigits int
}

func LoadLocale(cldr *CLDR, ldml *Ldml) *Locale {
	locale := &Locale{
		IsRoot:    ldml.Identity.Language.Type == "root",
		Code:      ldml.Identity.Language.Type,
		Name:      ldml.Identity.Language.Type,
		Territory: ldml.Identity.Territory.Type,
		Parents:   []*Locale{},
		Parent:    nil,
		Symbols:   map[string]*Symbol{},
		Decimals:  map[string]*FormatGroup{},
	}

	if cldr.RootLocale != nil {
		locale.Parents = append(locale.Parents, cldr.RootLocale)
	}

	if !locale.IsRoot {
		locale.Parent = cldr.RootLocale
	}

	if locale.Territory != "" {
		locale.Code = locale.Code + "_" + locale.Territory
	} else {
		locale.IsBase = !locale.IsRoot
	}

	AttachCurrencies(locale, cldr, ldml)
	AttachTerritories(locale, cldr, ldml)
	AttachTimeZones(locale, cldr, ldml)
	AttachNumber(locale, cldr, ldml)

	return locale
}
