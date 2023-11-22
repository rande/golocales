// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type TimeZone struct {
	Code string
	Name string
	Zone string
	City string
}

// This implementation is wrong
// The correct implementation should be:
//  1. Load the timezones from the supplemental data: supplementalData/metaZones.xml
//  2. Match custom translation with the timezones
//  3. Also some name can be duplicated, so we need to append the city name
func AttachTimeZones(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var timezones map[string]*TimeZone = map[string]*TimeZone{}

	for _, t := range cldr.MetaZones {
		if t.Type == "Etc/Unknown" || t.Type == "Etc/UTC" {
			continue
		}

		timezones[t.Type] = &TimeZone{
			Code: t.Type,
			Zone: t.Zone,
		}

	}

	for _, t := range ldml.Dates.TimeZoneNames.Zone {
		if _, ok := timezones[t.Type]; ok {
			timezones[t.Type].City = t.ExemplarCity.Text
		}
	}

	checks := map[string]int{}
	for _, t := range ldml.Dates.TimeZoneNames.Metazone {
		for _, z := range timezones {
			if t.Type == z.Zone {
				if _, ok := checks[t.Type]; !ok {
					checks[t.Type] = 0
				}
				checks[t.Type]++
			}
		}
	}

	for _, t := range ldml.Dates.TimeZoneNames.Metazone {
		for _, z := range timezones {
			if t.Type == z.Zone {
				name := t.Long.Standard.Text
				if checks[t.Type] > 1 {
					name = fmt.Sprintf("%s (%s)", name, z.Code)
				}

				timezones[z.Code].Name = name
			}
		}
	}

	locale.TimeZones = timezones
}
