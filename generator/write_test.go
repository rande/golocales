package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Write_Ldml(t *testing.T) {

	ldml, err := LoadLdml("fixtures/ldml_main.xml")

	assert.Nil(t, err)

	locale := LoadLocale(GetCLDR(), ldml)

	buffer := bytes.NewBuffer([]byte{})

	err = WriteGo(locale, buffer)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Printf("%s\n", errors.Unwrap(err))
	}

	assert.Nil(t, err)

	fmt.Printf("%s\n", buffer.String())
}
