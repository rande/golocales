// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttachPattern(t *testing.T) {

	// 	<currencyFormat type="standard">
	// 	<pattern>¤#,##0.00</pattern>
	// 	<pattern alt="alphaNextToNumber">¤ #,##0.00</pattern>
	// 	<pattern alt="noCurrency">#,##0.00</pattern>
	// </currencyFormat>
	// <currencyFormat type="accounting">
	// 	<pattern>¤#,##0.00;(¤#,##0.00)</pattern>
	// 	<pattern alt="alphaNextToNumber">¤ #,##0.00;(¤ #,##0.00)</pattern>
	// 	<pattern alt="noCurrency">#,##0.00;(#,##0.00)</pattern>
	// </currencyFormat>
	tests := []struct {
		pattern        string
		expected       string
		primaryGroup   int
		secondaryGroup int
	}{
		{"¤#,##0.00", "¤0.00", 3, 3},                         // currency format
		{"¤ #,##0.00;(¤ #,##0.00)", "¤ 0.00;(¤ 0.00)", 3, 3}, // accounting format
		{"#,##0.###", "0.000", 3, 3},                         // decimal format
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			format := &NumberFormat{
				Pattern: tt.pattern,
			}

			AttachPattern(format)

			assert.Equal(t, tt.expected, format.StandardPattern)
			assert.Equal(t, tt.primaryGroup, format.PrimaryGroupingSize)
			assert.Equal(t, tt.secondaryGroup, format.SecondaryGroupingSize)
		})
	}
}
