// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func ifEmptyString(s string, def string) string {
	if s == "" {
		return def
	}

	return s
}

func LoadXml(filename string, strct interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err = xml.Unmarshal(data, strct); err != nil {
		return err
	}

	return nil
}

func LoadLdml(filename string) (*Ldml, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ldml Ldml
	err = xml.Unmarshal(data, &ldml)
	if err != nil {
		return nil, err
	}

	return &ldml, nil
}

func LoadLocaleFromFile(path string, cldr *CLDR) *Locale {
	ldml := &Ldml{}
	if err := LoadXml(path, ldml); err != nil {
		log.Panic(err.Error())
	}

	return LoadLocale(cldr, ldml)
}
