// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestForCountryCode(t *testing.T) {
// 	tests := []struct {
// 		countryCode      string
// 		wantCurrencyCode string
// 		wantOK           bool
// 	}{
// 		{"FR", "EUR", true},
// 		{"RS", "RSD", true},
// 		{"XX", "", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run("", func(t *testing.T) {
// 			gotCurrencyCode, gotOK := ForCountryCode(tt.countryCode)
// 			if gotOK != tt.wantOK {
// 				t.Errorf("got %v, want %v", gotOK, tt.wantOK)
// 			}
// 			if gotCurrencyCode != tt.wantCurrencyCode {
// 				t.Errorf("got %q, want %q", gotCurrencyCode, tt.wantCurrencyCode)
// 			}
// 		})
// 	}
// }

// func TestGetCurrencyCodes(t *testing.T) {
// 	currencyCodes := GetCurrencyCodes()
// 	var got [10]string
// 	copy(got[:], currencyCodes[0:10])
// 	want := [10]string{"AUD", "CAD", "CHF", "EUR", "GBP", "JPY", "NOK", "NZD", "SEK", "USD"}
// 	// Confirm that the first 10 currency codes are the "G10" ones.
// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

func TestIsValid(t *testing.T) {
	tests := []struct {
		currencyCode string
		want         bool
	}{
		{"", true},
		{"INVALID", false},
		{"XXX", false},
		{"usd", false},
		{"USD", true},
		{"EUR", true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsValid(tt.currencyCode)
			assert.Equal(t, tt.want, got, "got %v, want %v", got, tt.want)
		})
	}
}

func TestGetNumericCode(t *testing.T) {
	numericCode, ok := GetNumericCode("USD")
	if !ok {
		t.Errorf("got %v, want true", ok)
	}
	if numericCode != "840" {
		t.Errorf("got %v, want 840", numericCode)
	}

	// Non-existent currency code.
	numericCode, ok = GetNumericCode("XXX")
	if ok {
		t.Errorf("got %v, want false", ok)
	}
	if numericCode != "000" {
		t.Errorf("got %v, want 000", numericCode)
	}
}

func TestGetDigits(t *testing.T) {
	digits, ok := GetCurrencyDigits("USD")
	if !ok {
		t.Errorf("got %v, want true", ok)
	}
	if digits != 2 {
		t.Errorf("got %v, want 2", digits)
	}

	// Non-existent currency code.
	digits, ok = GetCurrencyDigits("XXX")
	if ok {
		t.Errorf("got %v, want false", ok)
	}
	if digits != 0 {
		t.Errorf("got %v, want 0", digits)
	}
}

// func TestGetSymbol(t *testing.T) {
// 	tests := []struct {
// 		currencyCode string
// 		locale       currency.Locale
// 		wantSymbol   string
// 		wantOk       bool
// 	}{
// 		{"XXX", currency.NewLocale("en"), "XXX", false},
// 		{"usd", currency.NewLocale("en"), "usd", false},
// 		{"CHF", currency.NewLocale("en"), "CHF", true},
// 		{"USD", currency.NewLocale("en"), "$", true},
// 		{"USD", currency.NewLocale("en-US"), "$", true},
// 		{"USD", currency.NewLocale("en-AU"), "US$", true},
// 		{"USD", currency.NewLocale("es"), "US$", true},
// 		{"USD", currency.NewLocale("es-ES"), "US$", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run("", func(t *testing.T) {
// 			gotSymbol, gotOk := currency.GetSymbol(tt.currencyCode, tt.locale)
// 			if gotSymbol != tt.wantSymbol {
// 				t.Errorf("got %v, want %v", gotSymbol, tt.wantSymbol)
// 			}
// 			if gotOk != tt.wantOk {
// 				t.Errorf("got %v, want %v", gotOk, tt.wantOk)
// 			}
// 		})
// 	}
// }
