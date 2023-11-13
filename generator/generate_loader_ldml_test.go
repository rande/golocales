package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func Test_Load_Simple_Ldml(t *testing.T) {
	ldml, err := LoadLdml("fixtures/ldml_inherited.xml")

	assert.Nil(t, err)

	assert.Equal(t, "$Revision$", ldml.Identity.Version.Number)
	assert.Equal(t, "fr", ldml.Identity.Language.Type)
	assert.Equal(t, "YT", ldml.Identity.Territory.Type)
}

func Test_Load_Full_Ldml(t *testing.T) {
	ldml, err := LoadLdml("fixtures/ldml_main.xml")

	assert.Nil(t, err)

	assert.Equal(t, "$Revision$", ldml.Identity.Version.Number)
	// assert.Equal(t, "fr", ldml.Identity.Language.Type)
	// assert.Equal(t, "YT", ldml.Identity.Territory.Type)
	assert.NotNil(t, ldml.LocaleDisplayNames.Territories)
	assert.Len(t, ldml.LocaleDisplayNames.Territories.Territory, 12)

	territories := GetTerritories(ldml)
	assert.Len(t, territories, 8)
}

func Test_Write_Ldml(t *testing.T) {

	ldml, err := LoadLdml("fixtures/ldml_main.xml")

	assert.Nil(t, err)

	locale := BuildLocale(ldml)

	buffer := bytes.NewBuffer([]byte{})

	err = WriteGo(locale, buffer)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Printf("%s\n", errors.Unwrap(err))
	}

	assert.Nil(t, err)

	fmt.Printf("%s\n", buffer.String())
}

func Test_Safe(t *testing.T) {

	var countries = map[string]string{
		"UG": "אוגאַנדע",
		"US": "פֿ\"ש",
	}

	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)

		utf8.DecodeRune([]byte(name))
	}
}
