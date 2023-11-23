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

var TimezoneDenyList = map[string]bool{
	// Exceptional reservations
	"Etc/UTC":     true,
	"Etc/Unknown": true,
	"Etc/GMT":     true,
	"CST6CDT":     true,
	"EST5EDT":     true,
	"MST7MDT":     true,
	"PST8PDT":     true,
}

// This implementation is wrong
// The correct implementation should be:
//  1. Load the timezones from the supplemental data: supplementalData/metaZones.xml
//  2. Match custom translation with the timezones
//  3. Also some name can be duplicated, so we need to append the city name
//  4. Some Ldml does not have any timezone configured, in this case we can
//     delete all timezones - the parent locale will be used
func AttachTimeZones(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var timezones map[string]*TimeZone = map[string]*TimeZone{}

	for _, t := range cldr.MetaZones {

		if _, ok := TimezoneDenyList[t.Type]; ok {
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

	overwrites := false

	for _, m := range ldml.Dates.TimeZoneNames.Metazone {
		for _, t := range timezones {
			if m.Type == t.Zone {
				name := m.Long.Standard.Text
				if checks[m.Type] > 1 {
					if t.City != "" {
						name = fmt.Sprintf("%s (%s)", name, t.City)
					} else {
						name = fmt.Sprintf("%s (%s)", name, t.Code)
					}

				}

				timezones[t.Code].Name = name

				overwrites = true
			}
		}
	}

	if !overwrites && !locale.IsRoot {
		return
	}

	locale.TimeZones = timezones
}
