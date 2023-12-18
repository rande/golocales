// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

type CLDR struct {
	Path        string
	Validities  []*Validity
	RootLocale  *Locale
	MetaZones   []*MetaZone
	Locales     map[string]*Locale
	Territories map[string]*Territory
	Currencies  map[string]*Currency
	DayPeriods  map[string][]*DayPeriodRule
}

func LoadCLDR(CldrPath string) *CLDR {
	cldr := &CLDR{}
	cldr.Path = CldrPath
	cldr.Locales = map[string]*Locale{}
	cldr.Territories = map[string]*Territory{}
	cldr.Currencies = map[string]*Currency{}
	cldr.DayPeriods = map[string][]*DayPeriodRule{}

	// load validity files
	validityFiles := map[string]func(cldr *CLDR, supplemental *SupplementalData){
		"currency.xml": AttachValidity,
		"language.xml": AttachValidity,
		"region.xml":   AttachValidity,
		// "script.xml",
		// "subdivision.xml",
		// "unit.xml",
		// "variant.xml",
	}

	// validities are required to load root module
	for file, handler := range validityFiles {
		fmt.Printf(" > Loading validity file: %s\n", file)

		supplemental := &SupplementalData{}
		if err := LoadXml(CldrPath+"/validity/"+file, supplemental); err != nil {
			log.Panic(err.Error())
		}

		handler(cldr, supplemental)
	}

	// load supplemental files
	supplementalFiles := map[string]func(cldr *CLDR, supplemental *SupplementalData){
		// "attributeValueValidity.xml",
		// "characters.xml",
		// "coverageLevels.xml",
		"dayPeriods.xml": AttachDayPeriodRules,
		// "genderList.xml",
		// "grammaticalFeatures.xml",
		// "languageGroup.xml",
		// "languageInfo.xml",
		// "likelySubtags.xml",
		"metaZones.xml": AttachMetaZones,
		// "numberingSystems.xml",
		// "ordinals.xml",
		// "pluralRanges.xml",
		// "plurals.xml",
		// "rgScope.xml",
		// "subdivisions.xml",
		"supplementalData.xml": AttachSupplementalData,
		// "supplementalMetadata.xml",
		// "units.xml",
		// "windowsZones.xml",
	}

	for file, handler := range supplementalFiles {
		fmt.Printf(" > Loading supplemental file: %s\n", file)

		supplemental := &SupplementalData{}
		if err := LoadXml(CldrPath+"/supplemental/"+file, supplemental); err != nil {
			log.Panic(err.Error())
		}

		handler(cldr, supplemental)
	}

	return cldr
}

func (cldr *CLDR) GetValidity(code, status string) *Validity {
	for _, v := range cldr.Validities {
		if v.From == code && v.Status == status {
			return v
		}
	}

	return nil
}
