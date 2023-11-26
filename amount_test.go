// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/rande/golocales"
	"github.com/stretchr/testify/assert"
)

func TestNewCurrency(t *testing.T) {
	_, err := golocales.NewCurrency("INVALID", "USD")
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != "INVALID" {
			t.Errorf("got %v, want INVALID", e.Number)
		}
		wantError := `invalid number "INVALID"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	_, err = golocales.NewCurrency("10.99", "usd")
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "usd" {
			t.Errorf("got %v, want usd", e.CurrencyCode)
		}
		wantError := `invalid currency code "usd"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	a, err := golocales.NewCurrency("10.99", "USD")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if a.Number() != "10.99" {
		t.Errorf("got %v, want 10.99", a.Number())
	}
	if a.Code() != "USD" {
		t.Errorf("got %v, want USD", a.Code())
	}
	if a.String() != "10.99 USD" {
		t.Errorf("got %v, want 10.99 USD", a.String())
	}
}

func TestNewCurrencyFromBigInt(t *testing.T) {
	_, err := golocales.NewCurrencyFromBigInt(nil, "USD")
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != "nil" {
			t.Errorf("got %v, want nil", e.Number)
		}
		wantError := `invalid number "nil"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	_, err = golocales.NewCurrencyFromBigInt(big.NewInt(1099), "usd")
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "usd" {
			t.Errorf("got %v, want usd", e.CurrencyCode)
		}
		wantError := `invalid currency code "usd"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	// An integer larger than math.MaxInt64.
	hugeInt, _ := big.NewInt(0).SetString("922337203685477598799", 10)
	tests := []struct {
		n            *big.Int
		currencyCode string
		wantNumber   string
	}{
		{big.NewInt(2099), "USD", "20.99"},
		{big.NewInt(5000), "USD", "50.00"},
		{big.NewInt(50), "JPY", "50"},
		{hugeInt, "USD", "9223372036854775987.99"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, err := golocales.NewCurrencyFromBigInt(tt.n, tt.currencyCode)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if a.Number() != tt.wantNumber {
				t.Errorf("got %v, want %v", a.Number(), tt.wantNumber)
			}
			if a.Code() != tt.currencyCode {
				t.Errorf("got %v, want %v", a.Code(), tt.currencyCode)
			}
		})
	}
}

func TestNewCurrencyFromInt64(t *testing.T) {
	_, err := golocales.NewCurrencyFromInt64(1099, "usd")
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "usd" {
			t.Errorf("got %v, want usd", e.CurrencyCode)
		}
		wantError := `invalid currency code "usd"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	tests := []struct {
		n            int64
		currencyCode string
		wantNumber   string
	}{
		{2099, "USD", "20.99"},
		{5000, "USD", "50.00"},
		{50, "JPY", "50"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, err := golocales.NewCurrencyFromInt64(tt.n, tt.currencyCode)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if a.Number() != tt.wantNumber {
				t.Errorf("got %v, want %v", a.Number(), tt.wantNumber)
			}
			if a.Code() != tt.currencyCode {
				t.Errorf("got %v, want %v", a.Code(), tt.currencyCode)
			}
		})
	}
}

func TestAmount_BigInt(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		want         *big.Int
	}{
		{"20.99", "USD", big.NewInt(2099)},
		// Number with additional decimals.
		{"12.3564", "USD", big.NewInt(1236)},
		// Number with no decimals.
		{"50", "USD", big.NewInt(5000)},
		{"50", "JPY", big.NewInt(50)},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			got := a.BigInt()
			if got.Cmp(tt.want) != 0 {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			// Confirm that a is unchanged.
			if a.Number() != tt.number {
				t.Errorf("got %v, want %v", a.Number(), tt.number)
			}
		})
	}
}

func TestAmount_Int64(t *testing.T) {
	// Number that can't be represented as an int64.
	a, _ := golocales.NewCurrency("922337203685477598799", "USD")
	n, err := a.Int64()
	if n != 0 {
		t.Error("expected a.Int64() to be 0")
	}
	if err == nil {
		t.Error("expected a.Int64() to return an error")
	}

	tests := []struct {
		number       string
		currencyCode string
		want         int64
	}{
		{"20.99", "USD", 2099},
		// Number with additional decimals.
		{"12.3564", "USD", 1236},
		// Number with no decimals.
		{"50", "USD", 5000},
		{"50", "JPY", 50},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			got, _ := a.Int64()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			// Confirm that a is unchanged.
			if a.Number() != tt.number {
				t.Errorf("got %v, want %v", a.Number(), tt.number)
			}
		})
	}
}

// func TestAmount_Convert(t *testing.T) {
// 	a, _ := NewCurrency("20.99", "USD")

// 	_, err := a.Convert("eur", "0.91")
// 	if e, ok := err.(InvalidCurrencyCodeError); ok {
// 		if e.CurrencyCode != "eur" {
// 			t.Errorf("got %v, want eur", e.CurrencyCode)
// 		}
// 		wantError := `invalid currency code "eur"`
// 		if e.Error() != wantError {
// 			t.Errorf("got %v, want %v", e.Error(), wantError)
// 		}
// 	} else {
// 		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
// 	}

// 	_, err = a.Convert("EUR", "INVALID")
// 	if e, ok := err.(InvalidNumberError); ok {
// 		if e.Number != "INVALID" {
// 			t.Errorf("got %v, want INVALID", e.Number)
// 		}
// 		wantError := `invalid number "INVALID"`
// 		if e.Error() != wantError {
// 			t.Errorf("got %v, want %v", e.Error(), wantError)
// 		}
// 	} else {
// 		t.Errorf("got %T, want InvalidNumberError", err)
// 	}

// 	b, err := a.Convert("EUR", "0.91")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if b.String() != "19.1009 EUR" {
// 		t.Errorf("got %v, want 19.1009 EUR", b.String())
// 	}
// 	// Confirm that a is unchanged.
// 	if a.String() != "20.99 USD" {
// 		t.Errorf("got %v, want 20.99 USD", a.String())
// 	}

// 	// An amount larger than math.MaxInt64.
// 	c, _ := NewCurrency("922337203685477598799", "USD")
// 	d, err := c.Convert("RSD", "100")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if d.String() != "92233720368547759879900 RSD" {
// 		t.Errorf("got %v, want 92233720368547759879900 RSD", d.String())
// 	}
// }

func TestAmount_Add(t *testing.T) {
	a, _ := golocales.NewCurrency("20.99", "USD")
	b, _ := golocales.NewCurrency("3.50", "USD")
	x, _ := golocales.NewCurrency("99.99", "EUR")
	var z golocales.Amount

	_, err := a.Add(x)
	if e, ok := err.(golocales.MismatchError); ok {
		if e.A != a {
			t.Errorf("got %v, want %v", e.A, a)
		}
		if e.B != x {
			t.Errorf("got %v, want %v", e.B, x)
		}
		wantError := `amounts "20.99 USD" and "99.99 EUR" have mismatched currency codes`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want MismatchError", err)
	}

	c, err := a.Add(b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if c.String() != "24.49 USD" {
		t.Errorf("got %v, want 24.49 USD", c.String())
	}
	// Confirm that a and b are unchanged.
	if a.String() != "20.99 USD" {
		t.Errorf("got %v, want 20.99 USD", a.String())
	}
	if b.String() != "3.50 USD" {
		t.Errorf("got %v, want 3.50 USD", b.String())
	}

	// An amount equal to math.MaxInt64.
	d, _ := golocales.NewCurrency("9223372036854775807", "USD")
	e, err := d.Add(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if e.String() != "9223372036854775827.99 USD" {
		t.Errorf("got %v, want 9223372036854775827.99 USD", e.String())
	}

	// Test that addition with the zero value works and yields the other operand.
	f, err := a.Add(z)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !f.Equal(a) {
		t.Errorf("%v + zero = %v, want %v", a, f, a)
	}

	g, err := z.Add(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !g.Equal(a) {
		t.Errorf("%v + zero = %v, want %v", a, g, a)
	}
}

func TestAmount_Sub(t *testing.T) {
	a, _ := golocales.NewCurrency("20.99", "USD")
	b, _ := golocales.NewCurrency("3.50", "USD")
	x, _ := golocales.NewCurrency("99.99", "EUR")
	var z golocales.Amount

	_, err := a.Sub(x)
	if e, ok := err.(golocales.MismatchError); ok {
		if e.A != a {
			t.Errorf("got %v, want %v", e.A, a)
		}
		if e.B != x {
			t.Errorf("got %v, want %v", e.B, x)
		}
		wantError := `amounts "20.99 USD" and "99.99 EUR" have mismatched currency codes`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want MismatchError", err)
	}

	c, err := a.Sub(b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if c.String() != "17.49 USD" {
		t.Errorf("got %v, want 17.49 USD", c.String())
	}
	// Confirm that a and b are unchanged.
	if a.String() != "20.99 USD" {
		t.Errorf("got %v, want 20.99 USD", a.String())
	}
	if b.String() != "3.50 USD" {
		t.Errorf("got %v, want 3.50 USD", b.String())
	}

	// An amount larger than math.MaxInt64.
	d, _ := golocales.NewCurrency("922337203685477598799", "USD")
	e, err := d.Sub(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if e.String() != "922337203685477598778.01 USD" {
		t.Errorf("got %v, want 922337203685477598778.01 USD", e.String())
	}

	// Test that subtraction with the zero value works.
	f, err := a.Sub(z)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !f.Equal(a) {
		t.Errorf("%v - zero = %v, want %v", a, f, a)
	}

	g, err := z.Sub(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	negA, err := a.Mul("-1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !g.Equal(negA) {
		t.Errorf("zero - %v = %v, want %v", a, g, negA)
	}
}

func TestAmount_Mul(t *testing.T) {
	a, _ := golocales.NewCurrency("20.99", "USD")

	_, err := a.Mul("INVALID")
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != "INVALID" {
			t.Errorf("got %v, want INVALID", e.Number)
		}
		wantError := `invalid number "INVALID"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	b, err := a.Mul("0.20")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if b.String() != "4.1980 USD" {
		t.Errorf("got %v, want 4.1980 USD", b.String())
	}
	// Confirm that a is unchanged.
	if a.String() != "20.99 USD" {
		t.Errorf("got %v, want 20.99 USD", a.String())
	}

	// An amount equal to math.MaxInt64.
	d, _ := golocales.NewCurrency("9223372036854775807", "USD")
	e, err := d.Mul("10")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if e.String() != "92233720368547758070 USD" {
		t.Errorf("got %v, want 92233720368547758070 USD", e.String())
	}
}

func TestAmount_Div(t *testing.T) {
	a, _ := golocales.NewCurrency("99.99", "USD")

	for _, n := range []string{"INVALID", "0"} {
		_, err := a.Div(n)
		if e, ok := err.(golocales.InvalidNumberError); ok {
			if e.Number != n {
				t.Errorf("got %v, want %v", e.Number, n)
			}
		} else {
			t.Errorf("got %T, want InvalidNumberError", err)
		}
	}

	b, err := a.Div("3")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if b.String() != "33.33 USD" {
		t.Errorf("got %v, want 33.33 USD", b.String())
	}
	// Confirm that a is unchanged.
	if a.String() != "99.99 USD" {
		t.Errorf("got %v, want 99.99 USD", a.String())
	}

	// An amount equal to math.MaxInt64.
	d, _ := golocales.NewCurrency("9223372036854775807", "USD")
	e, err := d.Div("0.5")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if e.String() != "18446744073709551614 USD" {
		t.Errorf("got %v, want 18446744073709551614 USD", e.String())
	}
}

func TestAmount_Round(t *testing.T) {
	tests := []struct {
		number       string
		currencyCode string
		want         string
	}{
		{"12.345", "USD", "12.35"},
		{"12.345", "JPY", "12"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.number, tt.currencyCode)
			b := a.Round()
			if b.Number() != tt.want {
				t.Errorf("got %v, want %v", b.Number(), tt.want)
			}
			// Confirm that a is unchanged.
			if a.Number() != tt.number {
				t.Errorf("got %v, want %v", a.Number(), tt.number)
			}
		})
	}
}

func TestAmount_RoundTo(t *testing.T) {
	tests := []struct {
		number string
		digits uint8
		mode   golocales.RoundingMode
		want   string
	}{
		{"12.343", 2, golocales.RoundHalfUp, "12.34"},
		{"12.345", 2, golocales.RoundHalfUp, "12.35"},
		{"12.347", 2, golocales.RoundHalfUp, "12.35"},

		{"12.343", 2, golocales.RoundHalfDown, "12.34"},
		{"12.345", 2, golocales.RoundHalfDown, "12.34"},
		{"12.347", 2, golocales.RoundHalfDown, "12.35"},

		{"12.343", 2, golocales.RoundUp, "12.35"},
		{"12.345", 2, golocales.RoundUp, "12.35"},
		{"12.347", 2, golocales.RoundUp, "12.35"},

		{"12.343", 2, golocales.RoundDown, "12.34"},
		{"12.345", 2, golocales.RoundDown, "12.34"},
		{"12.347", 2, golocales.RoundDown, "12.34"},

		{"12.344", 2, golocales.RoundHalfEven, "12.34"},
		{"12.345", 2, golocales.RoundHalfEven, "12.34"},
		{"12.346", 2, golocales.RoundHalfEven, "12.35"},

		{"12.334", 2, golocales.RoundHalfEven, "12.33"},
		{"12.335", 2, golocales.RoundHalfEven, "12.34"},
		{"12.336", 2, golocales.RoundHalfEven, "12.34"},

		// Negative amounts.
		{"-12.345", 2, golocales.RoundHalfUp, "-12.35"},
		{"-12.345", 2, golocales.RoundHalfDown, "-12.34"},
		{"-12.345", 2, golocales.RoundUp, "-12.35"},
		{"-12.345", 2, golocales.RoundDown, "-12.34"},
		{"-12.345", 2, golocales.RoundHalfEven, "-12.34"},
		{"-12.335", 2, golocales.RoundHalfEven, "-12.34"},

		// More digits that the amount has.
		{"12.345", 4, golocales.RoundHalfUp, "12.3450"},
		{"12.345", 4, golocales.RoundHalfDown, "12.3450"},

		// Same number of digits that the amount has.
		{"12.345", 3, golocales.RoundHalfUp, "12.345"},
		{"12.345", 3, golocales.RoundHalfDown, "12.345"},
		{"12.345", 3, golocales.RoundUp, "12.345"},
		{"12.345", 3, golocales.RoundDown, "12.345"},

		// 0 digits.
		{"12.345", 0, golocales.RoundHalfUp, "12"},
		{"12.345", 0, golocales.RoundHalfDown, "12"},
		{"12.345", 0, golocales.RoundUp, "13"},
		{"12.345", 0, golocales.RoundDown, "12"},

		// Amounts larger than math.MaxInt64.
		{"12345678901234567890.0345", 3, golocales.RoundHalfUp, "12345678901234567890.035"},
		{"12345678901234567890.0345", 3, golocales.RoundHalfDown, "12345678901234567890.034"},
		{"12345678901234567890.0345", 3, golocales.RoundUp, "12345678901234567890.035"},
		{"12345678901234567890.0345", 3, golocales.RoundDown, "12345678901234567890.034"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.number, "USD")
			b := a.RoundTo(tt.digits, tt.mode)
			if b.Number() != tt.want {
				t.Errorf("got %v, want %v", b.Number(), tt.want)
			}
			// Confirm that a is unchanged.
			if a.Number() != tt.number {
				t.Errorf("got %v, want %v", a.Number(), tt.number)
			}
		})
	}
}

func TestAmount_RoundToWithConcurrency(t *testing.T) {
	n := 2
	roundingModes := []golocales.RoundingMode{
		golocales.RoundHalfUp,
		golocales.RoundHalfDown,
		golocales.RoundUp,
		golocales.RoundDown,
	}

	for _, roundingMode := range roundingModes {
		roundingMode := roundingMode

		t.Run(fmt.Sprintf("rounding_mode_%d", roundingMode), func(t *testing.T) {
			t.Parallel()

			var allDone sync.WaitGroup
			allDone.Add(n)

			for i := 0; i < n; i++ {
				go func() {
					defer allDone.Done()
					amount, _ := golocales.NewCurrency("10.99", "EUR")
					amount.RoundTo(1, roundingMode)
				}()
			}

			allDone.Wait()
		})
	}
}

func TestAmount_Cmp(t *testing.T) {
	a, _ := golocales.NewCurrency("3.33", "USD")
	b, _ := golocales.NewCurrency("3.33", "EUR")
	_, err := a.Cmp(b)
	if e, ok := err.(golocales.MismatchError); ok {
		if e.A != a {
			t.Errorf("got %v, want %v", e.A, a)
		}
		if e.B != b {
			t.Errorf("got %v, want %v", e.B, b)
		}
		wantError := `amounts "3.33 USD" and "3.33 EUR" have mismatched currency codes`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want MismatchError", err)
	}

	tests := []struct {
		aNumber string
		bNumber string
		want    int
	}{
		{"3.33", "6.66", -1},
		{"3.33", "3.33", 0},
		{"6.66", "3.33", 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.aNumber, "USD")
			b, _ := golocales.NewCurrency(tt.bNumber, "USD")
			got, err := a.Cmp(b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_Equal(t *testing.T) {
	tests := []struct {
		aNumber       string
		aCurrencyCode string
		bNumber       string
		bCurrencyCode string
		want          bool
	}{
		{"3.33", "USD", "6.66", "EUR", false},
		{"3.33", "USD", "3.33", "EUR", false},
		{"3.33", "USD", "3.33", "USD", true},
		{"3.33", "USD", "6.66", "USD", false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.aNumber, tt.aCurrencyCode)
			b, _ := golocales.NewCurrency(tt.bNumber, tt.bCurrencyCode)
			got := a.Equal(b)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_Checks(t *testing.T) {
	tests := []struct {
		number       string
		wantPositive bool
		wantNegative bool
		wantZero     bool
	}{
		{"9.99", true, false, false},
		{"-9.99", false, true, false},
		{"0", false, false, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, _ := golocales.NewCurrency(tt.number, "USD")
			gotPositive := a.IsPositive()
			gotNegative := a.IsNegative()
			gotZero := a.IsZero()
			if gotPositive != tt.wantPositive {
				t.Errorf("positive: got %v, want %v", gotPositive, tt.wantPositive)
			}
			if gotNegative != tt.wantNegative {
				t.Errorf("negative: got %v, want %v", gotNegative, tt.wantNegative)
			}
			if gotZero != tt.wantZero {
				t.Errorf("zero: got %v, want %v", gotZero, tt.wantZero)
			}
		})
	}
}

func TestAmount_MarshalBinary(t *testing.T) {
	a, _ := golocales.NewCurrency("3.45", "USD")
	d, err := a.MarshalBinary()

	got := string(d)
	want := "1USD3.45"

	assert.NoError(t, err)
	assert.Equal(t, "1USD3.45", got, "got %v, want %v", got, want)
}

func TestAmount_UnmarshalBinary(t *testing.T) {
	d := []byte("US")
	a := &golocales.Amount{}
	err := a.UnmarshalBinary(d)
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "US" {
			t.Errorf("got %v, want US", e.CurrencyCode)
		}
		wantError := `invalid currency code "US"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	d = []byte("1USD3,60")
	err = a.UnmarshalBinary(d)
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != "3,60" {
			t.Errorf("got %v, want 3,60", e.Number)
		}
		wantError := `invalid number "3,60"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	d = []byte("1XXX2.60")
	err = a.UnmarshalBinary(d)
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "XXX" {
			t.Errorf("got %v, want XXX", e.CurrencyCode)
		}
		wantError := `invalid currency code "XXX"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	d = []byte("1USD3.45")
	err = a.UnmarshalBinary(d)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if a.Number() != "3.45" {
		t.Errorf("got %v, want 3.45", a.Number())
	}
	if a.Code() != "USD" {
		t.Errorf("got %v, want USD", a.Code())
	}
}

func TestAmount_MarshalJSON(t *testing.T) {
	a, _ := golocales.NewCurrency("3.45", "USD")
	d, err := json.Marshal(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	got := string(d)
	want := `{"unit":1,"number":"3.45","code":"USD"}`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAmount_UnmarshalJSON(t *testing.T) {
	d := []byte(`{"unit":1,"number":"INVALID","code":"USD"}`)
	unmarshalled := &golocales.Amount{}
	err := json.Unmarshal(d, unmarshalled)
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != "INVALID" {
			t.Errorf("got %v, want INVALID", e.Number)
		}
		wantError := `invalid number "INVALID"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	d = []byte(`{"unit":1,"number": {"key": "value"}, "code": "USD"}`)
	err = json.Unmarshal(d, unmarshalled)
	if e, ok := err.(golocales.InvalidNumberError); ok {
		if e.Number != `{"key": "value"}` {
			t.Errorf(`got %v, "want {"key": "value"}"`, e.Number)
		}
		wantError := `invalid number "{\"key\": \"value\"}"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidNumberError", err)
	}

	d = []byte(`{"unit":1,"number":3.45,"code":"USD"}`)
	err = json.Unmarshal(d, unmarshalled)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if unmarshalled.Number() != "3.45" {
		t.Errorf("got %v, want 3.45", unmarshalled.Number())
	}
	if unmarshalled.Code() != "USD" {
		t.Errorf("got %v, want USD", unmarshalled.Code())
	}

	d = []byte(`{"unit":1,"number":"3.45","code":"usd"}`)
	err = json.Unmarshal(d, unmarshalled)
	if e, ok := err.(golocales.InvalidCurrencyCodeError); ok {
		if e.CurrencyCode != "usd" {
			t.Errorf("got %v, want usd", e.CurrencyCode)
		}
		wantError := `invalid currency code "usd"`
		if e.Error() != wantError {
			t.Errorf("got %v, want %v", e.Error(), wantError)
		}
	} else {
		t.Errorf("got %T, want InvalidCurrencyCodeError", err)
	}

	d = []byte(`{"unit":1,"number":"3.45","code":"USD"}`)
	err = json.Unmarshal(d, unmarshalled)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if unmarshalled.Number() != "3.45" {
		t.Errorf("got %v, want 3.45", unmarshalled.Number())
	}
	if unmarshalled.Code() != "USD" {
		t.Errorf("got %v, want USD", unmarshalled.Code())
	}
}

func TestAmount_Value(t *testing.T) {
	a, _ := golocales.NewCurrency("3.45", "USD")
	got, _ := a.Value()
	want := "(3.45,1,USD)"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	var b golocales.Amount
	got, _ = b.Value()
	want = "(0,0,)"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAmount_Scan(t *testing.T) {
	tests := []struct {
		src              string
		wantNumber       string
		wantCurrencyCode string
		wantError        string
	}{
		{"", "0", "", ""},
		{"(3.45,1,USD)", "3.45", "USD", ""},
		{"(3.45,1,)", "0", "", `invalid currency code ""`},
		{"(,1,USD)", "0", "", `invalid number ""`},
		{"(0,1,)", "0", "", ""},
		{"(0,1,   )", "0", "", ""},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var a golocales.Amount
			err := a.Scan(tt.src)
			if a.Number() != tt.wantNumber {
				t.Errorf("number: got %v, want %v", a.Number(), tt.wantNumber)
			}
			if a.Code() != tt.wantCurrencyCode {
				t.Errorf("currency code: got %v, want %v", a.Code(), tt.wantCurrencyCode)
			}
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}
			if errStr != tt.wantError {
				t.Errorf("error: got %v, want %v", errStr, tt.wantError)
			}
		})
	}
}
