package main

type Locale struct {
	IsRoot      bool
	IsBase      bool
	Parent      *Locale
	Code        string
	Name        string
	Territory   string
	Territories map[string]Territory
	Currencies  map[string]Currency
	TimeZones   map[string]TimeZone
	Parents     []*Locale
	Numbers     map[string]*Number
	Decimals    map[string]*FormatGroup
}

func LoadLocale(cldr *CLDR, ldml *Ldml) *Locale {
	locale := &Locale{
		IsRoot:    ldml.Identity.Language.Type == "root",
		Code:      ldml.Identity.Language.Type,
		Name:      ldml.Identity.Language.Type,
		Territory: ldml.Identity.Territory.Type,
		Parents:   []*Locale{},
		Parent:    nil,
		Numbers:   map[string]*Number{},
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
	AttachNumberSymbols(locale, cldr, ldml)
	AttachNumberDecimals(locale, cldr, ldml)

	return locale
}
