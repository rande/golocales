// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/cockroachdb/apd/v3"
)

// RoundingMode determines how the amount will be rounded.
type RoundingMode uint8

const (
	// RoundHalfUp rounds up if the next digit is >= 5.
	RoundHalfUp RoundingMode = iota
	// RoundHalfDown rounds up if the next digit is > 5.
	RoundHalfDown
	// RoundUp rounds away from 0.
	RoundUp
	// RoundDown rounds towards 0, truncating extra digits.
	RoundDown
	// RoundHalfEven rounds up if the next digit is > 5. If the next digit is equal
	// to 5, it rounds to the nearest even decimal. Also called bankers' rounding.
	RoundHalfEven
)

// InvalidNumberError is returned when a numeric string can't be converted to a decimal.
type InvalidNumberError struct {
	Number string
}

func (e InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number %q", e.Number)
}

// InvalidCurrencyCodeError is returned when a currency code is invalid or unrecognized.
type InvalidCurrencyCodeError struct {
	CurrencyCode string
}

func (e InvalidCurrencyCodeError) Error() string {
	return fmt.Sprintf("invalid currency code %q", e.CurrencyCode)
}

// MismatchError is returned when two amounts have mismatched currency codes.
type MismatchError struct {
	A Amount
	B Amount
}

func (e MismatchError) Error() string {
	return fmt.Sprintf("amounts %q and %q have mismatched currency codes", e.A, e.B)
}

// Amount stores a decimal number with its currency code.
type Amount struct {
	number apd.Decimal
	code   string
	unit   unitSystem
}

type unitSystem uint8

const (
	unitEmpty    unitSystem = 0
	unitCurrency            = 1
	unitPercent             = 2
)

func NewCurrency(n, code string) (Amount, error) {
	number := apd.Decimal{}
	if _, _, err := number.SetString(n); err != nil {
		return Amount{}, InvalidNumberError{n}
	}

	if code == "" || !IsValid(code) {
		return Amount{}, InvalidCurrencyCodeError{code}
	}

	return Amount{number, code, unitCurrency}, nil
}

// NewAmountFromBigInt creates a new Amount from a big.Int and a currency code.
func NewCurrencyFromBigInt(n *big.Int, currencyCode string) (Amount, error) {
	if n == nil {
		return Amount{}, InvalidNumberError{"nil"}
	}
	d, ok := GetCurrencyDigits(currencyCode)
	if !ok {
		return Amount{}, InvalidCurrencyCodeError{currencyCode}
	}
	coeff := new(apd.BigInt).SetMathBigInt(n)
	number := apd.NewWithBigInt(coeff, -int32(d))

	return Amount{*number, currencyCode, unitCurrency}, nil
}

// NewAmountFromInt64 creates a new Amount from an int64 and a currency code.
func NewCurrencyFromInt64(n int64, currencyCode string) (Amount, error) {
	d, ok := GetCurrencyDigits(currencyCode)
	if !ok {
		return Amount{}, InvalidCurrencyCodeError{currencyCode}
	}
	number := apd.Decimal{}
	number.SetFinite(n, -int32(d))

	return Amount{number, currencyCode, unitCurrency}, nil
}

func NewAmount(n string) (Amount, error) {
	number := apd.Decimal{}
	if _, _, err := number.SetString(n); err != nil {
		return Amount{}, InvalidNumberError{n}
	}

	return Amount{number, "", unitEmpty}, nil
}

// NewAmountFromBigInt creates a new Amount from a big.Int and a currency code.
func NewAmountFromBigInt(n *big.Int) (Amount, error) {
	if n == nil {
		return Amount{}, InvalidNumberError{"nil"}
	}
	// TODO: what is the default digit, ie: read the documentation
	coeff := new(apd.BigInt).SetMathBigInt(n)
	number := apd.NewWithBigInt(coeff, -int32(DefaultDigits))

	return Amount{*number, "", unitEmpty}, nil
}

// NewAmountFromInt64 creates a new Amount from an int64 and a currency code.
func NewAmountFromInt64(n int64) (Amount, error) {
	number := apd.Decimal{}
	number.SetFinite(n, -int32(DefaultDigits))

	return Amount{number, "", unitEmpty}, nil
}

// Number returns the number as a numeric string.
func (a Amount) Number() string {
	return a.number.String()
}

func (a Amount) Code() string {
	return a.code
}

func (a Amount) Unit() unitSystem {
	return a.unit
}

// String returns the string representation of a.
func (a Amount) String() string {
	if a.unit == unitEmpty {
		return a.Number()
	}

	if a.unit == unitCurrency {
		return a.Number() + " " + a.code
	}

	if a.unit == unitPercent {
		return a.Number() + " %"
	}

	panic("Invalid type")
}

// BigInt returns a in minor units, as a big.Int.
func (a Amount) BigInt() *big.Int {
	r := a.Round()
	return r.number.Coeff.MathBigInt()
}

// Int64 returns a in minor units, as an int64.
// If a cannot be represented in an int64, an error is returned.
func (a Amount) Int64() (int64, error) {
	n := a.Round().number
	n.Exponent = 0
	return n.Int64()
}

// Add adds a and b together and returns the result.
func (a Amount) Add(b Amount) (Amount, error) {
	if a.unit != b.unit || a.code != b.code {
		if a.Equal(Amount{}) {
			return b, nil
		}
		if b.Equal(Amount{}) {
			return a, nil
		}
		return Amount{}, MismatchError{a, b}
	}
	result := apd.Decimal{}
	ctx := decimalContext(&a.number, &b.number)
	ctx.Add(&result, &a.number, &b.number)

	return Amount{result, a.code, a.unit}, nil
}

// Sub subtracts b from a and returns the result.
func (a Amount) Sub(b Amount) (Amount, error) {
	if a.unit != b.unit || a.code != b.code {
		if a.Equal(Amount{}) {
			// 0-b == -b
			var result apd.Decimal
			result.Neg(&b.number)
			return Amount{result, b.code, b.unit}, nil
		}
		if b.Equal(Amount{}) {
			return a, nil
		}
		return Amount{}, MismatchError{a, b}
	}
	result := apd.Decimal{}
	ctx := decimalContext(&a.number, &b.number)
	ctx.Sub(&result, &a.number, &b.number)

	return Amount{result, a.code, a.unit}, nil
}

// Mul multiplies a by n and returns the result.
func (a Amount) Mul(n string) (Amount, error) {
	result := apd.Decimal{}
	if _, _, err := result.SetString(n); err != nil {
		return Amount{}, InvalidNumberError{n}
	}
	ctx := decimalContext(&a.number, &result)
	ctx.Mul(&result, &a.number, &result)

	return Amount{result, a.code, a.unit}, nil
}

// Div divides a by n and returns the result.
func (a Amount) Div(n string) (Amount, error) {
	result := apd.Decimal{}
	if _, _, err := result.SetString(n); err != nil {
		return Amount{}, InvalidNumberError{n}
	}
	if result.IsZero() {
		return Amount{}, InvalidNumberError{n}
	}
	ctx := decimalContext(&a.number, &result)
	ctx.Quo(&result, &a.number, &result)
	result.Reduce(&result)

	return Amount{result, a.code, a.unit}, nil
}

// Round is a shortcut for RoundTo(currency.DefaultDigits, currency.RoundHalfUp).
func (a Amount) Round() Amount {
	return a.RoundTo(DefaultDigits, RoundHalfUp)
}

// RoundTo rounds a to the given number of fraction digits.
func (a Amount) RoundTo(digits uint8, mode RoundingMode) Amount {
	if digits == DefaultDigits {
		digits, _ = GetCurrencyDigits(a.code)
	}

	result := apd.Decimal{}
	ctx := roundingContext(&a.number, mode)
	ctx.Quantize(&result, &a.number, -int32(digits))

	return Amount{result, a.code, a.unit}
}

// Cmp compares a and b and returns:
//
//	-1 if a <  b
//	0 if a == b
//	+1 if a >  b
func (a Amount) Cmp(b Amount) (int, error) {
	if a.unit != b.unit || a.code != b.code {
		return -1, MismatchError{a, b}
	}
	return a.number.Cmp(&b.number), nil
}

// Equal returns whether a and b are equal.
func (a Amount) Equal(b Amount) bool {
	if a.unit != b.unit || a.code != b.code {
		return false
	}
	return a.number.Cmp(&b.number) == 0
}

// IsPositive returns whether a is positive.
func (a Amount) IsPositive() bool {
	zero := apd.New(0, 0)
	return a.number.Cmp(zero) == 1
}

// IsNegative returns whether a is negative.
func (a Amount) IsNegative() bool {
	zero := apd.New(0, 0)
	return a.number.Cmp(zero) == -1
}

// IsZero returns whether a is zero.
func (a Amount) IsZero() bool {
	zero := apd.New(0, 0)
	return a.number.Cmp(zero) == 0
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (a Amount) MarshalBinary() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(string(a.unit))
	// need to be 3 char long unit
	if len(a.code) != 3 {
		panic("Invalid code lenght for amount!")
	}
	buf.WriteString(a.code)
	buf.WriteString(a.Number())

	return buf.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (a *Amount) UnmarshalBinary(data []byte) error {
	if len(data) < 3 {
		return InvalidCurrencyCodeError{string(data)}
	}

	// TODO: ADD THE CORRECT TYPE HERE
	unit := unitEmpty
	n := string(data[4:])
	code := string(data[1:4])
	number := apd.Decimal{}
	if _, _, err := number.SetString(n); err != nil {
		return InvalidNumberError{n}
	}

	a.unit = unit
	a.number = number
	a.code = code

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Unit   unitSystem `json:"unit"`
		Number string     `json:"number"`
		Code   string     `json:"code"`
	}{
		Number: a.Number(),
		Code:   a.code,
		Unit:   a.unit,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Amount) UnmarshalJSON(data []byte) error {
	aux := struct {
		Number json.RawMessage `json:"number"`
		Code   string          `json:"code"`
		Unit   int             `json:"unit"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	var auxNumber string
	if err = json.Unmarshal(aux.Number, &auxNumber); err != nil {
		auxNumber = string(aux.Number)
	}

	number := apd.Decimal{}
	if _, _, err := number.SetString(auxNumber); err != nil {
		return InvalidNumberError{auxNumber}
	}

	var unit unitSystem

	if aux.Unit == 0 {
		unit = unitEmpty
	}

	if aux.Unit == 1 {
		unit = unitCurrency
	}

	if aux.Unit == 2 {
		unit = unitPercent
	}

	if unit == unitCurrency && (aux.Code == "" || !IsValid(aux.Code)) {
		return InvalidCurrencyCodeError{aux.Code}
	}

	a.number = number
	a.code = aux.Code
	a.unit = unit

	return nil
}

// Value implements the database/driver.Valuer interface.
//
// Allows storing amounts in a PostgreSQL composite type.
func (a Amount) Value() (driver.Value, error) {
	return fmt.Sprintf("(%v,%v,%v)", a.Number(), a.unit, a.code), nil
}

// Scan implements the database/sql.Scanner interface.
//
// Allows scanning amounts from a PostgreSQL composite type.
func (a *Amount) Scan(src interface{}) error {
	// Wire format: "(9.99,1,USD)".
	input := src.(string)
	if len(input) == 0 {
		return nil
	}
	input = strings.Trim(input, "()")
	values := strings.Split(input, ",")
	n := values[0]

	var unit unitSystem

	if values[1] == "0" {
		unit = unitEmpty
	}

	if values[1] == "1" {
		unit = unitCurrency
	}

	if values[1] == "2" {
		unit = unitPercent
	}

	code := values[2]
	number := apd.Decimal{}
	if _, _, err := number.SetString(n); err != nil {
		return InvalidNumberError{n}
	}
	// Allow the zero value (number=0, currencyCode is empty).
	// An empty currencyCode consists of 3 spaces when stored in a char(3).
	if (code == "" || code == "   ") && number.IsZero() {
		a.number = number
		a.code = ""
		return nil
	}

	if unit == unitCurrency && (code == "" || !IsValid(code)) {
		return InvalidCurrencyCodeError{code}
	}

	a.number = number
	a.code = code
	a.unit = unit

	return nil
}

var (
	decimalContextPrecision19 = apd.BaseContext.WithPrecision(19)
	decimalContextPrecision39 = apd.BaseContext.WithPrecision(39)
)

// decimalContext returns the decimal context to use for a calculation.
// The returned context is not safe for concurrent modification.
func decimalContext(decimals ...*apd.Decimal) *apd.Context {
	// Choose between decimal64 (19 digits) and decimal128 (39 digits)
	// based on operand size (> int32), for increased performance.
	for _, d := range decimals {
		if d.Coeff.BitLen() > 31 {
			return decimalContextPrecision39
		}
	}
	return decimalContextPrecision19
}

// roundingContext returns the decimal context to use for rounding.
// It optimizes for the most common RoundHalfUp mode by returning a preallocated global context for it.
func roundingContext(decimal *apd.Decimal, mode RoundingMode) *apd.Context {
	if mode == RoundHalfUp {
		return decimalContext(decimal)
	}

	extModes := map[RoundingMode]apd.Rounder{
		RoundHalfDown: apd.RoundHalfDown,
		RoundUp:       apd.RoundUp,
		RoundDown:     apd.RoundDown,
		RoundHalfEven: apd.RoundHalfEven,
	}
	ctx := *decimalContext(decimal)
	ctx.Rounding = extModes[mode]

	return &ctx
}
