package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse_Region(t *testing.T) {

	message := "AC~G AI AZ\n\t   BA~B"

	countries := ParseValidityValues(message)

	assert.Equal(t, 9, len(countries))
	assert.Equal(t, []string{"AC", "CD", "CE", "CF", "CG", "AI", "AZ", "BA", "AB"}, countries)
}

func Test_Parse_Currency(t *testing.T) {

	message := `			AED AFN ALL AMD ANG AOA ARS AUD AWG AZN
	BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYN BZD
`

	countries := ParseValidityValues(message)

	assert.Equal(t, 25, len(countries))
	assert.Equal(t, []string{"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD"}, countries)
}
