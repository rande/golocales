// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

type Annotation struct {
	Type  string
	Label string
	Cp    string
	Key   string
}

type Locale struct {
	IsRoot          bool
	IsBase          bool
	Parent          *Locale
	Code            string
	Name            string
	Territory       string
	Territories     map[string]*Territory
	Currencies      map[string]*Currency
	CurrencySymbols map[string]Symbol
	TimeZones       map[string]*TimeZone
	Number          *Number
	Keys            map[string]string
	Annotations     []*Annotation
}

func LoadLocale(cldr *CLDR, ldml *Ldml) *Locale {
	locale := &Locale{
		IsRoot:    ldml.Identity.Language.Type == "root",
		Code:      ldml.Identity.Language.Type,
		Name:      ldml.Identity.Language.Type,
		Territory: ldml.Identity.Territory.Type,
		Parent:    nil,
		Number: &Number{
			Symbols:    map[string]*Symbol{},
			Decimals:   map[string]FormatGroup{},
			Currencies: map[string]FormatGroup{},
		},
		Keys:        map[string]string{},
		Territories: map[string]*Territory{},
		Currencies:  map[string]*Currency{},
	}

	if !locale.IsRoot {
		locale.Parent = cldr.RootLocale
	}

	if locale.Territory != "" {
		locale.Parent = cldr.Locales[locale.Code]
		locale.Code = locale.Code + "_" + locale.Territory
	} else {
		locale.IsBase = !locale.IsRoot
	}

	AttachKeys(locale, cldr, ldml)
	AttachAnnotations(locale, cldr, ldml)
	AttachCurrencies(locale, cldr, ldml)
	AttachTerritories(locale, cldr, ldml)
	AttachTimeZones(locale, cldr, ldml)
	AttachNumber(locale, cldr, ldml)

	return locale
}

// The keys are used to filter data in each annotation files
// currency is device is french,
//
//	<annotation cp="€">devise | EUR | euro</annotation>
func AttachKeys(locale *Locale, cldr *CLDR, ldml *Ldml) {
	for _, key := range ldml.LocaleDisplayNames.Keys.Key {
		locale.Keys[key.Type] = strings.ToLower(key.Text)
	}
}

// For now, we only attach currency key, and draft must be empty
func AttachAnnotations(locale *Locale, cldr *CLDR, ldml *Ldml) {
	annotation := &XmlAnnotation{}

	module := locale.Code
	if locale.IsRoot {
		module = "en"
	}

	// load the related annotation for the root file, however the
	// file is empty, so we fallback to the en one.
	err := LoadXml(cldr.Path+"/annotations/"+module+".xml", annotation)

	if err != nil {
		fmt.Printf("Unable to find the annotation: %s\n", module)
		return
	}

	for _, a := range annotation.Annotations.Annotation {
		if a.Draft != "" {
			continue
		}

		annotation := &Annotation{
			Type:  a.Type,
			Cp:    a.Cp,
			Label: a.Text,
		}

		locale.Annotations = append(locale.Annotations, annotation)
	}
}
