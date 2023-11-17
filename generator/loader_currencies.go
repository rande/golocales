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

type Currency struct {
	Code string
	Name string
}

func AttachCurrencies(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var currencies map[string]Currency = make(map[string]Currency)

	list := cldr.GetValidity("currency", "regular")

	if list == nil {
		fmt.Printf("No currencies found\n")
		return
	}

	for _, t := range ldml.Numbers.Currencies.Currency {
		// if _, ok := TerritoriesDenyList[t.Type]; ok {
		// 	continue
		// }

		if !slices.Contains(list.List, strings.ToUpper(t.Type)) {
			continue
		}

		name := ""

		for _, displayName := range t.DisplayName {
			if displayName.Count != "" {
				continue
			}

			name = displayName.Text
		}

		currencies[t.Type] = Currency{
			Code: t.Type,
			Name: name,
		}
	}

	locale.Currencies = currencies
}
