// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

// Package currency handles currency amounts, provides currency information and formatting.
package golocales

import (
	"sort"

	"github.com/rande/golocales/dto"
	"github.com/rande/golocales/locales/root"
)

// DefaultDigits is a placeholder for each currency's number of fraction digits.
const DefaultDigits uint8 = 255

// ForCountryCode returns the currency code for a country code.
// func ForCountryCode(countryCode string) (currencyCode string, ok bool) {
// 	currencyCode, ok = countryCurrencies[countryCode]

// 	return currencyCode, ok
// }

// GetCurrencyCodes returns all known currency codes.
// func GetCurrencyCodes() []string {
// 	return currencyCodes
// }

// IsValid checks whether a currency code is valid.
//
// An empty currency code is considered valid.
func IsValid(currencyCode string) bool {
	if currencyCode == "" {
		return true
	}

	_, ok := root.Locale().Currencies[currencyCode]

	return ok
}

// GetNumericCode returns the numeric code for a currency code.
func GetNumericCode(currencyCode string) (numericCode string, ok bool) {
	if currencyCode == "" || !IsValid(currencyCode) {
		return "000", false
	}
	return root.Locale().Currencies[currencyCode].Numeric, true
}

// GetDigits returns the number of fraction digits for a currency code.
func GetCurrencyDigits(currencyCode string) (digits uint8, ok bool) {
	if currencyCode == "" || !IsValid(currencyCode) {
		return 0, false
	}
	return root.Locale().Currencies[currencyCode].Digits, true
}

func GetDigits(amount Amount) (digits uint8, ok bool) {
	if amount.unit == unitCurrency {
		return GetCurrencyDigits(amount.code)
	}

	// TODO: find out for other unit type
	return 2, true
}

// GetSymbol returns the symbol for a currency code.
func GetSymbol(currencyCode string, locale *dto.Locale) (symbol string, ok bool) {
	if currencyCode == "" || !IsValid(currencyCode) {
		return currencyCode, false
	}

	l := locale

	for {
		if currency, ok := l.Currencies[currencyCode]; ok {
			if currency.Symbol != "" {
				return currency.Symbol, true
			}
		}

		if l.Parent != nil {
			l = l.Parent
		} else {
			return currencyCode, true
		}
	}
}

// // getFormat returns the format for a locale.
// func getFormat(locale Locale) currencyFormat {
// 	var format currencyFormat
// 	// CLDR considers "en" and "en-US" to be equivalent.
// 	// Fall back immediately for better performance
// 	enUSLocale := Locale{Language: "en", Territory: "US"}
// 	if locale == enUSLocale {
// 		locale = Locale{Language: "en"}
// 	}
// 	for {
// 		localeID := locale.String()
// 		if cf, ok := currencyFormats[localeID]; ok {
// 			format = cf
// 			break
// 		}
// 		locale = locale.GetParent()
// 		if locale.IsEmpty() {
// 			break
// 		}
// 	}

// 	return format
// }

// contains returns whether the sorted slice a contains x.
// The slice must be sorted in ascending order.
func contains(a []string, x string) bool {
	n := len(a)
	if n < 10 {
		// Linear search is faster with a small number of elements.
		for _, v := range a {
			if v == x {
				return true
			}
		}
	} else {
		// Binary search is faster with a large number of elements.
		i := sort.SearchStrings(a, x)
		if i < n && a[i] == x {
			return true
		}
	}
	return false
}
