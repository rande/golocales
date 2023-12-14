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

type FormattingOptions struct {
	// AddPlusSign inserts the plus sign in front of positive amounts.
	// Defaults to false.
	AddPlusSign bool
	// Style=accounting formats the amount using the accounting style.
	// For example, "-3.00 USD" in the "en" locale is formatted as "($3.00)" instead of "-$3.00".
	// Defaults to "currency".
	Style string
	// CurrencyDisplay specifies how the currency will be displayed (symbol/code/none).
	// Defaults to currency.DisplaySymbol.
	CurrencyDisplay Display
	// RoundingMode specifies how the formatted amount will be rounded.
	// Defaults to currency.RoundHalfUp.
	RoundingMode RoundingMode
	// NoGrouping turns off grouping of major digits.
	// Defaults to false.
	NoGrouping bool
	// MinDigits specifies the minimum number of fraction digits.
	// All zeroes past the minimum will be removed (0 => no trailing zeroes).
	// Defaults to currency.DefaultDigits (e.g. 2 for USD, 0 for RSD).
	MinDigits uint8
	// MaxDigits specifies the maximum number of fraction digits.
	// Formatted amounts will be rounded to this number of digits.
	// Defaults to 6, so that most amounts are shown as-is (without rounding).
	MaxDigits uint8
}

func CreateFormattingOptions() *FormattingOptions {
	return &FormattingOptions{
		AddPlusSign:     false,
		Style:           "currency",
		NoGrouping:      false,
		MinDigits:       DefaultDigits,
		MaxDigits:       6,
		CurrencyDisplay: DisplaySymbol,
	}
}

// Formatter formats and parses currency amounts.
type AmountFormatter struct {
	locale  *dto.Locale
	formats map[string]*dto.NumberFormat
	symbol  *dto.Symbol

	// SymbolMap specifies custom symbols for individual currency codes.
	// For example, "USD": "$" means that the $ symbol will be used even if
	// the current locale's symbol is different ("US$", "$US", etc).
	SymbolMap map[string]string
}

// NewAmountFormatter creates a new AmountFormatter for the given locale.
func NewAmountFormatter(locale *dto.Locale) *AmountFormatter {
	// load correct AmountFormatter
	cFormats := locale.GetCurrencyFormats(locale.Number.DefaultNumberSystem, "default_standard")
	aFormats := locale.GetCurrencyFormats(locale.Number.DefaultNumberSystem, "default_accounting")
	dFormats := locale.GetDecimalFormats(locale.Number.DefaultNumberSystem, "default")
	pFormats := locale.GetPercentFormats(locale.Number.DefaultNumberSystem, "default")

	if cFormats == nil || len(cFormats) == 0 {
		panic(fmt.Sprintf("Unable to find default currency formats: %s", locale.Name))
	}

	if dFormats == nil || len(dFormats) == 0 {
		panic(fmt.Sprintf("Unable to find default decimal formats: %s", locale.Name))
	}

	if pFormats == nil || len(pFormats) == 0 {
		panic(fmt.Sprintf("Unable to find default decimal formats: %s", locale.Name))
	}

	// if the accounting format is not defined, use the default currency format
	if aFormats == nil || len(aFormats) == 0 {
		aFormats = cFormats
	}

	f := &AmountFormatter{
		locale: locale,
		symbol: locale.GetSymbol(locale.Number.DefaultNumberSystem),
		formats: map[string]*dto.NumberFormat{
			"currency":   cFormats[0],
			"decimal":    dFormats[0],
			"accounting": aFormats[0],
			"percent":    pFormats[0],
		},
		SymbolMap: make(map[string]string),
	}

	return f
}

// Locale returns the locale.
func (f *AmountFormatter) GetLocale() *dto.Locale {
	return f.locale
}

// Format formats a currency amount.
func (f *AmountFormatter) Format(amount Amount, options ...*FormattingOptions) string {
	// default value
	formattingOptions := CreateFormattingOptions()
	if len(options) > 0 {
		formattingOptions = options[0]
	}

	pattern := f.getPattern(amount, formattingOptions)

	if amount.IsNegative() {
		// The minus sign will be provided by the pattern.
		amount, _ = amount.Mul("-1")
	}

	formattedNumber := f.formatNumber(amount, formattingOptions)
	formattedCurrency := f.formatCurrency(amount.Code(), formattingOptions)

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
		"+", f.symbol.PlusSign,
		"-", f.symbol.MinusSign,
	}

	if amount.IsPercent() {
		replacements = append(replacements, "0", formattedNumber, "%", f.symbol.PercentSign)
	}

	if amount.IsNumber() {
		replacements = append(replacements, "0.00", formattedNumber)
	}

	if amount.IsCurrency() {
		replacements = append(replacements, "0.00", formattedNumber)
		if formattedCurrency == "" {
			// Many patterns have a non-breaking space between
			// the number and currency, not needed in this case.
			replacements = append(replacements, "\u00a0¤", "", "¤\u00a0", "", "¤", "")
		} else {
			replacements = append(replacements, "¤", formattedCurrency)
		}
	}

	r := strings.NewReplacer(replacements...)

	fmt.Printf("pattern: %#v %s %s\n", replacements, pattern, r.Replace(pattern))

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

// getPattern returns a positive or negative pattern for a currency amount.
func (f *AmountFormatter) getPattern(amount Amount, options *FormattingOptions) string {
	pattern := ""
	// -- deal with currency pattern
	if amount.IsCurrency() {
		pattern = f.formats["currency"].StandardPattern

		if options.Style == "accounting" {
			pattern = f.formats["accounting"].StandardPattern
		}
	}

	// -- deal with number pattern
	if amount.IsNumber() {
		pattern = f.formats["decimal"].StandardPattern
	}

	// -- deal with percent pattern
	if amount.IsPercent() {
		pattern = f.formats["percent"].StandardPattern
	}

	// the accounting format is `#,##0.00 ¤;(#,##0.00 ¤)`
	// the first section is for positive number, and the second part is the negative
	// representation (ie: without the minus sign).
	patterns := strings.Split(pattern, ";")

	if amount.IsNegative() && len(patterns) > 1 {
		return patterns[1]
	} else if amount.IsNegative() {
		return "-" + patterns[0]
	} else {
		pattern = patterns[0]
	}

	if pattern == "" {
		panic(fmt.Sprintf("Unable to find pattern for %s", amount.Code()))
	}

	if options.AddPlusSign {
		// this is not really part of the accounting format,
		// but was part of the original Currency library
		pattern = "+" + pattern
	}

	return pattern
}

// formatNumber formats the number for display.
func (f *AmountFormatter) formatNumber(amount Amount, options *FormattingOptions) string {
	if amount.IsPercent() {
		amount, _ = amount.Mul("100")
	}

	minDigits := options.MinDigits
	if minDigits == DefaultDigits {
		minDigits, _ = GetDigits(amount)
	}
	maxDigits := options.MaxDigits
	if maxDigits == DefaultDigits {
		maxDigits, _ = GetDigits(amount)
	}
	amount = amount.RoundTo(maxDigits, options.RoundingMode)
	numberParts := strings.Split(amount.Number(), ".")
	majorDigits := f.groupMajorDigits(numberParts[0], amount.unit, options)
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
func (f *AmountFormatter) formatCurrency(currencyCode string, options *FormattingOptions) string {
	var formatted string
	switch options.CurrencyDisplay {
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
func (f *AmountFormatter) groupMajorDigits(majorDigits string, unit unitSystem, options *FormattingOptions) string {

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

	if options.NoGrouping || format.PrimaryGroupingSize == 0 {
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
	majorDigits = strings.Join(groups, f.symbol.CurrencyGroup)

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
