// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MetaZone struct {
	Type string
	Zone string
}

type DayPeriodRule struct {
	Type   string
	From   int
	Before int
	At     int
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

func AttachSupplementalData(cldr *CLDR, supplemental *SupplementalData) {
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

	for _, t := range supplemental.CodeMappings.TerritoryCodes {
		cldr.Territories[t.Type] = &Territory{
			Code:    t.Type,
			Numeric: t.Numeric,
			Alpha3:  t.Alpha3,
		}
	}

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

func AttachDayPeriodRules(cldr *CLDR, supplemental *SupplementalData) {
	for _, s := range supplemental.DayPeriodRuleSet {
		if s.Type == "selection" {
			fmt.Printf("Skip selection rule: %s\n", s.Type)
			continue
		}

		for _, r := range s.DayPeriodRules {
			locales := strings.Split(r.Locales, " ")
			for _, locale := range locales {
				if _, ok := cldr.DayPeriods[locale]; !ok {
					cldr.DayPeriods[locale] = []*DayPeriodRule{}
				}

				addAm := true
				addPm := true
				for _, rule := range r.DayPeriodRule {
					at := -1
					from, _ := strconv.Atoi(strings.Replace(rule.From, ":", "", 1))
					before, _ := strconv.Atoi(strings.Replace(rule.Before, ":", "", 1))
					if len(rule.At) > 0 {
						at, _ = strconv.Atoi(strings.Replace(rule.At, ":", "", 1))
					}

					if rule.Type == "am" {
						addAm = false
					}

					if rule.Type == "pm" {
						addPm = false
					}

					// fmt.Printf("Add rule: %s %v\n", rule.Type, addPm)
					cldr.DayPeriods[locale] = append(cldr.DayPeriods[locale], &DayPeriodRule{
						Type:   rule.Type,
						From:   from,
						Before: before,
						At:     at,
					})
				}

				// am/pm periods should be always defined, but some locales are missing
				// those information, so we need to add them manually
				if addAm {
					cldr.DayPeriods[locale] = append(cldr.DayPeriods[locale], &DayPeriodRule{
						Type:   "am",
						From:   0,
						Before: 1200,
						At:     -1,
					})
				}

				if addPm {
					cldr.DayPeriods[locale] = append(cldr.DayPeriods[locale], &DayPeriodRule{
						Type:   "pm",
						From:   1200,
						Before: 2400,
						At:     -1,
					})
				}
			}
		}
	}
}
