// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/rande/golocales/dto"
)

// Display represents the currency display type.
type Display uint8

const (
	// DisplaySymbol shows the currency symbol.
	DisplaySymbol Display = iota
	// DisplayCode shows the currency code.
	DisplayCode
	// DisplayNone shows nothing, hiding the currency.
	DisplayNone
)

var localDigits = map[string]string{
	"arab":    "٠١٢٣٤٥٦٧٨٩",
	"arabext": "۰۱۲۳۴۵۶۷۸۹",
	"beng":    "০১২৩৪৫৬৭৮৯",
	"deva":    "०१२३४५६७८९",
	"mymr":    "၀၁၂၃၄၅၆၇၈၉",
}

// Formatter formats and parses currency amounts.
type AmountFormatter struct {
	locale  *dto.Locale
	formats map[string]*dto.NumberFormat
	symbol  *dto.Symbol

	// NoGrouping turns off grouping of major digits.
	// Defaults to false.
	noGrouping bool
	// MinDigits specifies the minimum number of fraction digits.
	// All zeroes past the minimum will be removed (0 => no trailing zeroes).
	// Defaults to currency.DefaultDigits (e.g. 2 for USD, 0 for RSD).
	minDigits uint8
	// MaxDigits specifies the maximum number of fraction digits.
	// Formatted amounts will be rounded to this number of digits.
	// Defaults to 6, so that most amounts are shown as-is (without rounding).
	maxDigits uint8
	// RoundingMode specifies how the formatted amount will be rounded.
	// Defaults to currency.RoundHalfUp.
	roundingMode RoundingMode
	// CurrencyDisplay specifies how the currency will be displayed (symbol/code/none).
	// Defaults to currency.DisplaySymbol.
	currencyDisplay Display
	// SymbolMap specifies custom symbols for individual currency codes.
	// For example, "USD": "$" means that the $ symbol will be used even if
	// the current locale's symbol is different ("US$", "$US", etc).
	SymbolMap map[string]string
}

// NewAmountFormatter creates a new AmountFormatter for the given locale.
func NewAmountFormatter(locale *dto.Locale) *AmountFormatter {
	// load correct AmountFormatter

	cFormats := locale.GetCurrencyFormats(locale.Number.DefaultNumberSystem, "default_standard")
	dFormats := locale.GetDecimalFormats(locale.Number.DefaultNumberSystem, "default")

	if cFormats != nil && len(cFormats) == 0 {
		panic(fmt.Sprintf("Unable to find default currency formats: %s", locale.Name))
	}

	if dFormats != nil && len(dFormats) == 0 {
		panic(fmt.Sprintf("Unable to find default decimal formats: %s", locale.Name))
	}

	f := &AmountFormatter{
		locale: locale,
		symbol: locale.GetSymbol(locale.Number.DefaultNumberSystem),
		formats: map[string]*dto.NumberFormat{
			"currency": cFormats[0],
			"decimal":  dFormats[0],
		},
		minDigits:       DefaultDigits,
		maxDigits:       6,
		roundingMode:    RoundHalfUp,
		currencyDisplay: DisplaySymbol,
		SymbolMap:       make(map[string]string),
	}

	return f
}

// Locale returns the locale.
func (f *AmountFormatter) GetLocale() *dto.Locale {
	return f.locale
}

// Format formats a currency amount.
func (f *AmountFormatter) Format(amount Amount) string {
	pattern := ""
	if amount.IsCurrency() {
		pattern = f.formats["currency"].StandardPattern
	}

	if amount.IsNumber() {
		pattern = f.formats["decimal"].StandardPattern
	}

	if pattern == "" {
		panic(fmt.Sprintf("Unable to find pattern for %s", amount.Code()))
	}

	// Thomas: I have commented this operation as it seems to be a bug
	//         the output of amount is 1234.00 for negative amount
	//         and the replace cannot find the minus sign anymore
	// if amount.IsNegative() {
	// 	// The minus sign will be provided by the pattern.
	// 	amount, _ = amount.Mul("-1")
	// }

	formattedNumber := f.formatNumber(amount)
	formattedCurrency := f.formatCurrency(amount.Code())

	if formattedCurrency != "" {
		// CLDR requires having a space between the letters
		// in a currency symbol and adjacent numbers.
		if strings.Contains(pattern, "0¤") {
			r, _ := utf8.DecodeRuneInString(formattedCurrency)
			if unicode.IsLetter(r) {
				formattedCurrency = "\u00a0" + formattedCurrency
			}
		} else if strings.Contains(pattern, "¤0") {
			r, _ := utf8.DecodeLastRuneInString(formattedCurrency)
			if unicode.IsLetter(r) {
				formattedCurrency = formattedCurrency + "\u00a0"
			}
		}
	}

	replacements := []string{
		"0.00", formattedNumber,
		"+", f.symbol.PlusSign,
		"-", f.symbol.MinusSign,
	}

	if formattedCurrency == "" {
		// Many patterns have a non-breaking space between
		// the number and currency, not needed in this case.
		replacements = append(replacements, "\u00a0¤", "", "¤\u00a0", "", "¤", "")
	} else {
		replacements = append(replacements, "¤", formattedCurrency)
	}

	r := strings.NewReplacer(replacements...)

	return r.Replace(pattern)
}

// // Parse parses a formatted amount.
// func (f *AmountFormatter) Parse(s, currencyCode string) (Amount, error) {
// 	symbol, _ := GetSymbol(currencyCode, f.locale)
// 	replacements := []string{
// 		f.format.decimalSeparator, ".",
// 		f.format.groupingSeparator, "",
// 		f.format.plusSign, "+",
// 		f.format.minusSign, "-",
// 		symbol, "",
// 		currencyCode, "",
// 		"\u200e", "",
// 		"\u200f", "",
// 		"\u00a0", "",
// 		" ", "",
// 	}
// 	if f.format.numberingSystem != numLatn {
// 		digits := localDigits[f.format.numberingSystem]
// 		for i, v := range strings.Split(digits, "") {
// 			replacements = append(replacements, v, strconv.Itoa(i))
// 		}
// 	}
// 	if f.AccountingStyle {
// 		replacements = append(replacements, "(", "-", ")", "")
// 	}
// 	r := strings.NewReplacer(replacements...)
// 	n := r.Replace(s)

// 	return NewAmount(n, currencyCode)
// }

// // getPattern returns a positive or negative pattern for a currency amount.
// func (f *AmountFormatter) getPattern(amount Amount) string {
// 	var patterns []string
// 	if f.usesAccountingPattern() {
// 		patterns = strings.Split(f.format.accountingPattern, ";")
// 	} else {
// 		patterns = strings.Split(f.format.standardPattern, ";")
// 	}

// 	switch {
// 	case amount.IsNegative():
// 		if len(patterns) == 1 {
// 			return "-" + patterns[0]
// 		}
// 		return patterns[1]
// 	case f.AddPlusSign:
// 		if len(patterns) == 1 || f.usesAccountingPattern() {
// 			return "+" + patterns[0]
// 		}
// 		return strings.Replace(patterns[1], "-", "+", 1)
// 	default:
// 		return patterns[0]
// 	}
// }

// // usesAccountingPattern returns whether the AmountFormatter needs to use the accounting pattern.
// func (f *AmountFormatter) usesAccountingPattern() bool {
// 	return f.AccountingStyle && f.format.accountingPattern != ""
// }

// formatNumber formats the number for display.
func (f *AmountFormatter) formatNumber(amount Amount) string {
	minDigits := f.minDigits
	if minDigits == DefaultDigits {
		minDigits, _ = GetDigits(amount)
	}
	maxDigits := f.maxDigits
	if maxDigits == DefaultDigits {
		maxDigits, _ = GetDigits(amount)
	}
	amount = amount.RoundTo(maxDigits, f.roundingMode)
	numberParts := strings.Split(amount.Number(), ".")
	majorDigits := f.groupMajorDigits(numberParts[0], amount.unit)
	minorDigits := ""
	if len(numberParts) == 2 {
		minorDigits = numberParts[1]
	}
	if minDigits < maxDigits {
		// Strip any trailing zeroes.
		minorDigits = strings.TrimRight(minorDigits, "0")
		if len(minorDigits) < int(minDigits) {
			// Now there are too few digits, re-add trailing zeroes
			// until minDigits is reached.
			minorDigits += strings.Repeat("0", int(minDigits)-len(minorDigits))
		}
	}
	b := strings.Builder{}
	b.WriteString(majorDigits)
	if minorDigits != "" {
		b.WriteString(f.symbol.Decimal)
		b.WriteString(minorDigits)
	}

	formatted := f.localizeDigits(b.String())

	return formatted
}

// formatCurrency formats the currency for display.
func (f *AmountFormatter) formatCurrency(currencyCode string) string {
	var formatted string
	switch f.currencyDisplay {
	case DisplaySymbol:
		if symbol, ok := f.SymbolMap[currencyCode]; ok {
			formatted = symbol
		} else {
			formatted, _ = GetSymbol(currencyCode, f.locale)
		}
	case DisplayCode:
		formatted = currencyCode
	default:
		formatted = ""
	}

	return formatted
}

// groupMajorDigits groups major digits according to the currency format.
func (f *AmountFormatter) groupMajorDigits(majorDigits string, unit unitSystem) string {

	var format *dto.NumberFormat
	if unit == unitEmpty {
		format = f.formats["decimal"]
	}

	if unit == unitCurrency {
		format = f.formats["decimal"]
	}

	if unit == unitPercent {
		format = f.formats["percent"]
	}

	if format == nil {
		panic("No format found")
	}

	if f.noGrouping || format.PrimaryGroupingSize == 0 {
		return majorDigits
	}
	numDigits := len(majorDigits)
	minDigits := int(f.locale.Number.MinimumGroupingDigits)
	primarySize := int(format.PrimaryGroupingSize)
	secondarySize := int(format.SecondaryGroupingSize)
	if numDigits < (minDigits + primarySize) {
		return majorDigits
	}

	// Digits are grouped from right to left.
	// First the primary group, then the secondary groups.
	var groups []string
	groups = append(groups, majorDigits[numDigits-primarySize:numDigits])
	for i := numDigits - primarySize; i > 0; i = i - secondarySize {
		low := i - secondarySize
		if low < 0 {
			low = 0
		}
		groups = append(groups, majorDigits[low:i])
	}
	// Reverse the groups and reconstruct the digits.
	for i, j := 0, len(groups)-1; i < j; i, j = i+1, j-1 {
		groups[i], groups[j] = groups[j], groups[i]
	}
	majorDigits = strings.Join(groups, f.symbol.Group)

	return majorDigits
}

// localizeDigits replaces digits with their localized equivalents.
func (f *AmountFormatter) localizeDigits(number string) string {
	if f.locale.Number.DefaultNumberSystem == "latn" {
		return number
	}
	digits := localDigits[f.locale.Number.DefaultNumberSystem]
	replacements := make([]string, 0, 20)
	for i, v := range strings.Split(digits, "") {
		replacements = append(replacements, strconv.Itoa(i), v)
	}
	r := strings.NewReplacer(replacements...)
	number = r.Replace(number)

	return number
}
