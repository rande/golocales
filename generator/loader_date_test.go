// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_Date_Split_Pattern(t *testing.T) {
// 	pattern := "y.MMMM / EEEE-d"

// 	results := SplitDatePattern(pattern)

// 	assert.Equal(t, 9, len(results))
// 	assert.Equal(t, "y", results[0])
// 	assert.Equal(t, ".", results[1])
// 	assert.Equal(t, "MMMM", results[2])
// 	assert.Equal(t, " ", results[3])
// 	assert.Equal(t, "/", results[4])
// 	assert.Equal(t, " ", results[5])
// 	assert.Equal(t, "EEEE", results[6])
// 	assert.Equal(t, "-", results[7])
// 	assert.Equal(t, "d", results[8])
// }

func Test_Date_Parse_Pattern(t *testing.T) {

	pattern := "y.MMMM / EEEE-d"

	str, _ := SplitDatePattern(pattern)

	assert.Equal(t, "%d.%s / %s-%d", str)
}

func Test_Date_Parse_PatternWithLiteral(t *testing.T) {
	pattern := "d 'de' MMMM 'de' y"

	str, _ := SplitDatePattern(pattern)

	assert.Equal(t, "%d de %s de %d", str)
}

// func Test_Gen_DatePeriodFunc(t *testing.T) {

// 	periods := map[string]*DayPeriodRule{
// 		"am": {
// 			From:   0,
// 			Before: 1200,
// 			At:     -1,
// 		},
// 		"pm": {
// 			From:   1200,
// 			Before: 2400,
// 			At:     -1,
// 		},
// 		"midnight": {
// 			At: 0,
// 		},
// 		"noon": {
// 			At: 1200,
// 		},
// 	}

// 	periodFunc := GenDatePeriodFunc(periods)

// 	fmt.Printf("%s", periodFunc)

// }
