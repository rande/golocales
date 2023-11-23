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
	Code         string
	Name         string
	Symbol       string
	Digits       string
	Rounding     string
	CashDigits   string
	CashRounding string
	Numeric      string
}

func AttachCurrencies(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var currencies map[string]*Currency = map[string]*Currency{}

	list := cldr.GetValidity("currency", "regular")

	if list == nil {
		fmt.Printf("[%s] No currencies found\n", locale.Code)
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

		symbol := ""
		if len(t.Symbol) > 0 {
			symbol = t.Symbol[0].Text
		}

		currencies[t.Type] = &Currency{
			Code:         t.Type,
			Name:         name,
			Symbol:       symbol,
			Digits:       "2",
			Rounding:     "0",
			CashDigits:   "0",
			CashRounding: "0",
			Numeric:      "",
		}
	}

	locale.Currencies = currencies
}
