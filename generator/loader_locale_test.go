// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetCLDR() *CLDR {
	cldr := &CLDR{
		Validities: []*Validity{
			{"region", []string{"AC", "AD", "CD", "CF", "CG", "HK", "HM", "HN", "IO"}, "regular"},
		},
	}

	return cldr
}

func Test_Load_Simple_Ldml(t *testing.T) {
	ldml, err := LoadLdml("fixtures/ldml_inherited.xml")

	assert.Nil(t, err)

	assert.Equal(t, "$Revision$", ldml.Identity.Version.Number)
	assert.Equal(t, "fr", ldml.Identity.Language.Type)
	assert.Equal(t, "YT", ldml.Identity.Territory.Type)
}

func Test_Load_Full_Ldml(t *testing.T) {
	ldml, err := LoadLdml("fixtures/ldml_main.xml")

	locale := LoadLocale(GetCLDR(), ldml)

	assert.Nil(t, err)

	assert.Equal(t, "$Revision$", ldml.Identity.Version.Number)
	// assert.Equal(t, "fr", ldml.Identity.Language.Type)
	// assert.Equal(t, "YT", ldml.Identity.Territory.Type)
	assert.NotNil(t, ldml.LocaleDisplayNames.Territories)
	assert.Len(t, ldml.LocaleDisplayNames.Territories.Territory, 14)

	AttachTerritories(locale, GetCLDR(), ldml)

	assert.Len(t, locale.Territories, 8)
}
