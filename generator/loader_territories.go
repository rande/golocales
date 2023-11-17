// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"slices"
	"strings"
)

type Territory struct {
	Code string
	Name string
	Alt  string
}

var TerritoriesDenyList = map[string]bool{
	// Exceptional reservations
	"AC": true, // Ascension Island
	"CP": true, // Clipperton Island
	"CQ": true, // Island of Sark
	"DG": true, // Diego Garcia
	"EA": true, // Ceuta & Melilla
	"EU": true, // European Union
	"EZ": true, // Eurozone
	"IC": true, // Canary Islands
	"TA": true, // Tristan da Cunha
	"UN": true, // United Nations
	// User-assigned
	"QO": true, // Outlying Oceania
	"XA": true, // Pseudo-Accents
	"XB": true, // Pseudo-Bidi
	"XK": true, // Kosovo
	// Misc
	"ZZ": true, // Unknown Region
}

// @see https://en.wikipedia.org/wiki/ISO_3166-1_numeric#Withdrawn_codes
var TerritoriesWithdrawnCodes = []string{
	"128", //	Canton and Enderbury Islands
	"200", //	Czechoslovakia
	"216", //	Dronning Maud Land
	"230", //	Ethiopia
	"249", //	France, Metropolitan
	"278", //	German Democratic Republic
	"280", //	Germany, Federal Republic of
	"396", //	Johnston Island
	"488", //	Midway Islands
	"530", //	Netherlands Antilles
	"532", //	Netherlands Antilles
	"536", //	Neutral Zone
	"582", //	Pacific Islands (Trust Territory)
	"590", //	Panama
	"658", //	Saint Kitts-Nevis-Anguilla
	"720", //	Yemen, Democratic
	"736", //	Sudan
	"810", //	USSR
	"849", //	United States Miscellaneous Pacific Islands
	"872", //	Wake Island
	"886", //	Yemen Arab Republic
	"890", //	Yugoslavia, Socialist Federal Republic of
	"891", //	Serbia and Montenegro
}

func AttachTerritories(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var territories map[string]Territory = make(map[string]Territory)

	list := cldr.GetValidity("region", "regular")

	if list == nil {
		fmt.Printf("No regions found\n")
		return
	}

	for _, t := range ldml.LocaleDisplayNames.Territories.Territory {
		// if _, ok := TerritoriesDenyList[t.Type]; ok {
		// 	continue
		// }

		if !slices.Contains(list.List, strings.ToUpper(t.Type)) {
			continue
		}

		// if slices.Contains(TerritoriesWithdrawnCodes, t.Type) {
		// 	continue
		// }

		// we don't keep variants or short names.
		if t.Alt != "" {
			continue
		}

		// if _, err := strconv.Atoi(t.Type); err == nil {
		// 	continue
		// }

		territories[t.Type] = Territory{
			Code: t.Type,
			Name: t.Text,
			Alt:  t.Alt,
		}
	}

	locale.Territories = territories
}
