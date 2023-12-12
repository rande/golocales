// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales_test

import (
	"fmt"
	"testing"

	"github.com/rande/golocales"
	"github.com/rande/golocales/dto"
	"github.com/rande/golocales/locales/de_AT"
	"github.com/rande/golocales/locales/de_CH"
	"github.com/rande/golocales/locales/en"
	"github.com/rande/golocales/locales/en_NL"
	"github.com/rande/golocales/locales/en_US"
	"github.com/rande/golocales/locales/es"
	"github.com/rande/golocales/locales/fr"
	"github.com/rande/golocales/locales/fr_FR"
	"github.com/rande/golocales/locales/hi"
	"github.com/rande/golocales/locales/sr"
	"github.com/stretchr/testify/assert"
)

func TestAmountFormatter_Locale(t *testing.T) {
	locale := fr.GetLocale()
	formatter := golocales.NewAmountFormatter(locale)
	got := formatter.GetLocale().String()

	if got != "fr" {
		t.Errorf("got %v, want fr", got)
	}
}

func TestAmountFormatter_Format(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		want         string
		locale       *dto.Locale
	}{
		{"1234.59", "USD", "$1,234.59", en_US.GetLocale()},
		// {"1234.59", "USD", "en-CA", "US$1,234.59"},
		// {"1234.59", "USD", "de-CH", "$\u00a01’234.59"},
		{"2234.59", "USD", "2.234,59\u00a0US$", sr.GetLocale()},

		// {"-1234.59", "USD", "en-US", "-$1,234.59"},
		// {"-1234.59", "USD", "en-CA", "-US$1,234.59"},
		// {"-1234.59", "USD", "de-CH", "$-1’234.59"},
		{"-3234.59", "USD", "-3.234,59\u00a0US$", sr.GetLocale()},

		{"4234.00", "EUR", "€4,234.00", en.GetLocale()},
		{"-5234.00", "EUR", "-€5,234.00", en.GetLocale()},
		// {"1234.00", "EUR", "de-CH", "€\u00a01’234.00"},
		// {"1234.00", "EUR", "sr", "1.234,00\u00a0€"},

		{"6234.00", "CHF", "CHF\u00a06,234.00", en.GetLocale()},
		// {"1234.00", "CHF", "de-CH", "CHF\u00a01’234.00"},
		{"7234.00", "CHF", "7.234,00\u00a0CHF", sr.GetLocale()},

		// Arabic digits.
		// {"12345678.90", "USD", "ar", "\u200f١٢٬٣٤٥٬٦٧٨٫٩٠\u00a0US$"},
		// Arabic extended (Persian) digits.
		// {"12345678.90", "USD", "fa", "\u200e$۱۲٬۳۴۵٬۶۷۸٫۹۰"},
		// Bengali digits.
		// {"12345678.90", "USD", "bn", "১,২৩,৪৫,৬৭৮.৯০\u00a0US$"},
		// Devanagari digits.
		// {"12345678.90", "USD", "ne", "US$\u00a0१,२३,४५,६७८.९०"},
		// Myanmar (Burmese) digits.
		// {"12345678.90", "USD", "my", "၁၂,၃၄၅,၆၇၈.၉၀\u00a0US$"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, err := golocales.NewCurrency(tt.number, tt.currencyCode)
			assert.NoError(t, err)

			formatter := golocales.NewAmountFormatter(tt.locale)
			assert.Equal(t, tt.want, formatter.Format(amount))
		})
	}
}

func TestAmountFormatter_AccountingStyle(t *testing.T) {

	// It is possible to check result with the repl
	// from https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/NumberFormat/NumberFormat
	// or from you browser console:
	//
	// const number = -123456.789;
	// const options = {
	// 	currencySign: 'accounting',
	// 	style: 'currency',
	// 	currency: 'EUR'
	// };

	// console.log(new Intl.NumberFormat('en', options).format(number,),);
	tests := []struct {
		number       string
		currencyCode string
		AddPlusSign  bool
		want         string
		locale       *dto.Locale
	}{
		// Locale with an accounting pattern.
		{"11234.59", "USD", false, "$11,234.59", en.GetLocale()},
		{"-21234.59", "USD", false, "($21,234.59)", en.GetLocale()},
		{"31234.59", "USD", true, "+$31,234.59", en.GetLocale()},

		// Locale without an accounting pattern.
		{"41234.59", "EUR", false, "41.234,59 €", es.GetLocale()},
		{"-51234.59", "EUR", false, "-51.234,59 €", es.GetLocale()},
		{"61234.59", "EUR", true, "+61.234,59 €", es.GetLocale()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.AddPlusSign = tt.AddPlusSign
			options.Style = "accounting"

			got := formatter.Format(amount, options)

			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

func TestAmountFormatter_PlusSign(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		AddPlusSign  bool
		want         string
		locale       *dto.Locale
	}{
		{"123.99", "USD", false, "$123.99", en.GetLocale()},
		{"223.99", "USD", true, "+$223.99", en.GetLocale()},

		{"323.99", "USD", false, "$\u00a0323.99", de_CH.GetLocale()},
		// original value was $+423.99 but does not match the xml file: ¤\u00a0#,##0.00;¤-#,##0.00
		{"423.99", "USD", true, "+$\u00a0423.99", de_CH.GetLocale()},

		{"523.99", "USD", false, "523,99\u00a0$US", fr_FR.GetLocale()},
		{"623.99", "USD", true, "+623,99\u00a0$US", fr_FR.GetLocale()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.AddPlusSign = tt.AddPlusSign

			got := formatter.Format(amount, options)

			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

func TestAmountFormatter_Grouping(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		NoGrouping   bool
		want         string
		locale       *dto.Locale
	}{
		{"123.99", "USD", false, "$123.99", en.GetLocale()},
		{"1234.99", "USD", false, "$1,234.99", en.GetLocale()},
		{"1234567.99", "USD", false, "$1,234,567.99", en.GetLocale()},

		{"123.99", "USD", true, "$123.99", en.GetLocale()},
		{"1234.99", "USD", true, "$1234.99", en.GetLocale()},
		{"1234567.99", "USD", true, "$1234567.99", en.GetLocale()},

		// The "es" locale has a different minGroupingSize.
		{"123.99", "USD", false, "123,99\u00a0US$", es.GetLocale()},
		{"1234.99", "USD", false, "1234,99\u00a0US$", es.GetLocale()},
		{"12345.99", "USD", false, "12.345,99\u00a0US$", es.GetLocale()},
		{"1234567.99", "USD", false, "1.234.567,99\u00a0US$", es.GetLocale()},

		// The "hi" locale has a different secondaryGroupingSize.
		{"123.99", "USD", false, "$123.99", hi.GetLocale()},
		{"1234.99", "USD", false, "$1,234.99", hi.GetLocale()},
		{"1234567.99", "USD", false, "$12,34,567.99", hi.GetLocale()},
		{"12345678.99", "USD", false, "$1,23,45,678.99", hi.GetLocale()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.NoGrouping = tt.NoGrouping

			got := formatter.Format(amount, options)

			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

func TestAmountFormatter_Digits(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		minDigits    uint8
		maxDigits    uint8
		want         string
		locale       *dto.Locale
	}{
		{"59", "KRW", golocales.DefaultDigits, 6, "₩59", en.GetLocale()},
		{"59", "USD", golocales.DefaultDigits, 6, "$59.00", en.GetLocale()},
		{"59", "OMR", golocales.DefaultDigits, 6, "OMR\u00a059.000", en.GetLocale()},

		{"59.6789", "KRW", 0, golocales.DefaultDigits, "₩60", en.GetLocale()},
		{"59.6789", "USD", 0, golocales.DefaultDigits, "$59.68", en.GetLocale()},
		{"59.6789", "OMR", 0, golocales.DefaultDigits, "OMR\u00a059.679", en.GetLocale()},

		// minDigits:0 strips all trailing zeroes.
		{"59", "USD", 0, 6, "$59", en.GetLocale()},
		{"59.5", "USD", 0, 6, "$59.5", en.GetLocale()},
		{"59.56", "USD", 0, 6, "$59.56", en.GetLocale()},

		// minDigits can't override maxDigits.
		{"59.5", "USD", 3, 2, "$59.50", en.GetLocale()},
		{"59.567", "USD", 3, 2, "$59.57", en.GetLocale()},

		// maxDigits rounds the number.
		{"59.5", "USD", 2, 3, "$59.50", en.GetLocale()},
		{"59.567", "USD", 2, 3, "$59.567", en.GetLocale()},
		{"59.5678", "USD", 2, 3, "$59.568", en.GetLocale()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)

			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.MinDigits = tt.minDigits
			options.MaxDigits = tt.maxDigits

			got := formatter.Format(amount, options)
			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

func TestAmountFormatter_RoundingMode(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		roundingMode golocales.RoundingMode
		want         string
		locale       *dto.Locale
	}{
		{"1234.453", "USD", golocales.RoundHalfUp, "$1,234.45", en.GetLocale()},
		{"1234.455", "USD", golocales.RoundHalfUp, "$1,234.46", en.GetLocale()},
		{"1234.456", "USD", golocales.RoundHalfUp, "$1,234.46", en.GetLocale()},

		{"1234.453", "USD", golocales.RoundHalfDown, "$1,234.45", en.GetLocale()},
		{"1234.455", "USD", golocales.RoundHalfDown, "$1,234.45", en.GetLocale()},
		{"1234.457", "USD", golocales.RoundHalfDown, "$1,234.46", en.GetLocale()},

		{"1234.453", "USD", golocales.RoundUp, "$1,234.46", en.GetLocale()},
		{"1234.455", "USD", golocales.RoundUp, "$1,234.46", en.GetLocale()},
		{"1234.457", "USD", golocales.RoundUp, "$1,234.46", en.GetLocale()},

		{"1234.453", "USD", golocales.RoundDown, "$1,234.45", en.GetLocale()},
		{"1234.455", "USD", golocales.RoundDown, "$1,234.45", en.GetLocale()},
		{"1234.457", "USD", golocales.RoundDown, "$1,234.45", en.GetLocale()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.MaxDigits = golocales.DefaultDigits
			options.RoundingMode = tt.roundingMode

			got := formatter.Format(amount, options)

			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

func TestAmountFormatter_CurrencyDisplay(t *testing.T) {
	tests := []struct {
		number          string
		currencyCode    string
		currencyDisplay golocales.Display
		want            string
		locale          *dto.Locale
	}{
		{"1234.59", "USD", golocales.DisplaySymbol, "$1,234.59", en.GetLocale()},
		{"1234.59", "USD", golocales.DisplayCode, "USD\u00a01,234.59", en.GetLocale()},
		{"1234.59", "USD", golocales.DisplayNone, "1,234.59", en.GetLocale()},

		{"1234.59", "USD", golocales.DisplaySymbol, "$\u00a01.234,59", de_AT.GetLocale()},
		{"1234.59", "USD", golocales.DisplayCode, "USD\u00a01.234,59", de_AT.GetLocale()},
		{"1234.59", "USD", golocales.DisplayNone, "1.234,59", de_AT.GetLocale()},

		// {"1234.59", "USD", "sr-Latn", golocales.DisplaySymbol, "1.234,59\u00a0US$"},
		// {"1234.59", "USD", "sr-Latn", golocales.DisplayCode, "1.234,59\u00a0USD"},
		// {"1234.59", "USD", "sr-Latn", golocales.DisplayNone, "1.234,59"},

		// Confirm that any extra spacing around the currency is stripped
		// even when the negative amount is formatted with the accounting style.
		{"-1234.59", "USD", golocales.DisplayNone, "(1,234.59)", en.GetLocale()},
		{"-1234.59", "USD", golocales.DisplayNone, "(1.234,59)", en_NL.GetLocale()},
		// {"-1234.59", "USD", "sr-Latn", golocales.DisplayNone, "(1.234,59)"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			amount, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			formatter := golocales.NewAmountFormatter(tt.locale)

			options := golocales.CreateFormattingOptions()
			options.CurrencyDisplay = tt.currencyDisplay
			options.Style = "accounting"

			got := formatter.Format(amount, options)

			assert.Equal(t, tt.want, got, fmt.Sprintf("got %v, want %v", got, tt.want))
		})
	}
}

// func TestAmountFormatter_Parse(t *testing.T) {
// 	tests := []struct {
// 		s            string
// 		currencyCode string
// 		localeID     string
// 		want         string
// 	}{
// 		{"$1,234.59", "USD", "en", "1234.59"},
// 		{"USD\u00a01,234.59", "USD", "en", "1234.59"},
// 		{"1,234.59", "USD", "en", "1234.59"},
// 		{"1234.59", "USD", "en", "1234.59"},
// 		{"+1234.59", "USD", "en", "1234.59"},
// 		{"1234", "USD", "en", "1234"},

// 		{"-$1,234.59", "USD", "en", "-1234.59"},
// 		{"-USD\u00a01,234.59", "USD", "en", "-1234.59"},
// 		{"-1,234.59", "USD", "en", "-1234.59"},
// 		{"-1234.59", "USD", "en", "-1234.59"},
// 		{"(1234.59)", "USD", "en", "-1234.59"},

// 		{"€\u00a01.234,00", "EUR", "de-AT", "1234.00"},
// 		{"EUR\u00a01.234,00", "EUR", "de-AT", "1234.00"},
// 		{"1.234,00", "EUR", "de-AT", "1234.00"},
// 		{"1234,00", "EUR", "de-AT", "1234.00"},

// 		// Arabic digits.
// 		{"١٢٬٣٤٥٬٦٧٨٫٩٠\u00a0US$", "USD", "ar", "12345678.90"},
// 		// Arabic extended (Persian) digits.
// 		{"\u200e$۱۲٬۳۴۵٬۶۷۸٫۹۰", "USD", "fa", "12345678.90"},
// 		// Bengali digits.
// 		{"১,২৩,৪৫,৬৭৮.৯০\u00a0US$", "USD", "bn", "12345678.90"},
// 		// Devanagari digits.
// 		{"US$\u00a0१,२३,४५,६७८.९०", "USD", "ne", "12345678.90"},
// 		// Myanmar (Burmese) digits.
// 		{"၁၂,၃၄၅,၆၇၈.၉၀\u00a0US$", "USD", "my", "12345678.90"},
// 	}

// 	for _, tt := range tests {
// 		t.Run("", func(t *testing.T) {
// 			locale := currency.NewLocale(tt.localeID)
// 			formatter := currency.NewAmountFormatter(locale)
// 			// Allow parsing negative amounts formatted using parenthesis.
// 			formatter.AccountingStyle = true
// 			got, err := formatter.Parse(tt.s, tt.currencyCode)
// 			if err != nil {
// 				t.Errorf("unexpected error: %v", err)
// 			}
// 			if got.Number() != tt.want {
// 				t.Errorf("got %v, want %v", got, tt.want)
// 			}
// 			if got.CurrencyCode() != tt.currencyCode {
// 				t.Errorf("got %v, want %v", got.CurrencyCode(), tt.currencyCode)
// 			}
// 		})
// 	}
// }
