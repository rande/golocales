// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
)

type MetaZone struct {
	Type string
	Zone string
}

func AttachSupplemental(cldr *CLDR, supplemental *SupplementalData) {
	AttachValidity(cldr, supplemental)
	AttachMetaZones(cldr, supplemental)
	AttachCodeMapping(cldr, supplemental)
	AttachCurrencyFractions(cldr, supplemental)
	AttachCurrencyCode(cldr, supplemental)
}

func AttachMetaZones(cldr *CLDR, supplemental *SupplementalData) {
	for _, t := range supplemental.MetaZones.MetazoneInfo.Timezone {
		last := t.UsesMetazone[len(t.UsesMetazone)-1]

		meta := &MetaZone{
			Type: t.Type,
			Zone: last.Mzone,
		}

		cldr.MetaZones = append(cldr.MetaZones, meta)
	}
}

func AttachCodeMapping(cldr *CLDR, supplemental *SupplementalData) {
	for _, t := range supplemental.CodeMappings.TerritoryCodes {
		cldr.Territories[t.Type] = &Territory{
			Code:    t.Type,
			Numeric: t.Numeric,
			Alpha3:  t.Alpha3,
		}
	}
}

func AttachCurrencyFractions(cldr *CLDR, supplemental *SupplementalData) {
	for _, i := range supplemental.CurrencyData.Fractions.Info {
		if i.Iso4217 == "DEFAULT" {
			continue
		}

		cldr.Currencies[i.Iso4217] = &Currency{
			Code:         i.Iso4217,
			Digits:       ifEmptyString(i.Digits, "2"),
			Rounding:     ifEmptyString(i.Rounding, "0"),
			CashDigits:   ifEmptyString(i.CashDigits, "0"),
			CashRounding: ifEmptyString(i.CashRounding, "0"),
			Numeric:      "000",
		}
	}
}

func AttachCurrencyCode(cldr *CLDR, supplemental *SupplementalData) {
	for _, i := range supplemental.CodeMappings.CurrencyCodes {
		if _, ok := cldr.Currencies[i.Type]; !ok {
			cldr.Currencies[i.Type] = &Currency{
				Code:         i.Type,
				Digits:       "2",
				Rounding:     "0",
				CashDigits:   "0",
				CashRounding: "0",
			}
		}

		if v, err := strconv.Atoi(i.Numeric); err == nil {
			cldr.Currencies[i.Type].Numeric = fmt.Sprintf("%03d", v)
		}
	}
}
