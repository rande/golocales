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
	Const        string
}

func AttachCurrencies(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var currencies map[string]*Currency = map[string]*Currency{}

	list := cldr.GetValidity("currency", "regular")

	// fmt.Printf("[%s] Loading currencies %#v %d\n", locale.Code, list, len(list.List))
	// fmt.Printf("[%s] Loading currencies %#v %d\n", locale.Code, cldr.Currencies, len(cldr.Currencies))

	if list == nil {
		fmt.Printf("[%s] No currencies found\n", locale.Code)
		return
	}

	if locale.IsRoot {
		// root get all currencies
		for code, c := range cldr.Currencies {
			if !slices.Contains(list.List, strings.ToUpper(code)) {
				continue
			}

			currencies[code] = c
			currencies[code].Const = fmt.Sprintf("Currency_%s", strings.ToUpper(code))
		}
	}

	for _, t := range ldml.Numbers.Currencies.Currency {
		// if _, ok := TerritoriesDenyList[t.Type]; ok {
		// 	continue
		// }

		if _, ok := cldr.Currencies[t.Type]; !ok {
			continue
		}

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
			Digits:       ifEmptyString(cldr.Currencies[t.Type].Digits, "2"),
			Rounding:     ifEmptyString(cldr.Currencies[t.Type].Rounding, "0"),
			CashDigits:   ifEmptyString(cldr.Currencies[t.Type].CashDigits, "0"),
			CashRounding: ifEmptyString(cldr.Currencies[t.Type].CashRounding, "0"),
			Numeric:      ifEmptyString(cldr.Currencies[t.Type].Numeric, "000"),
			Const:        fmt.Sprintf("Currency_%s", strings.ToUpper(t.Type)),
		}
	}

	locale.Currencies = currencies
}
