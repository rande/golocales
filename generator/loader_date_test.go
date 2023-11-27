// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Date_Split_Pattern(t *testing.T) {
	pattern := "y.MMMM / EEEE-d"

	results := SplitDatePattern(pattern)

	assert.Equal(t, 9, len(results))
	assert.Equal(t, "y", results[0])
	assert.Equal(t, ".", results[1])
	assert.Equal(t, "MMMM", results[2])
	assert.Equal(t, " ", results[3])
	assert.Equal(t, "/", results[4])
	assert.Equal(t, " ", results[5])
	assert.Equal(t, "EEEE", results[6])
	assert.Equal(t, "-", results[7])
	assert.Equal(t, "d", results[8])
}

func Test_Date_Parse_Pattern(t *testing.T) {

	pattern := "y.MMMM / EEEE-d"

	str := ParseDatePattern(pattern)

	fmt.Printf("%s", str)
}
