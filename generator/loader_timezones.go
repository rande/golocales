// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

type TimeZone struct {
	Code string
	Name string
}

func AttachTimeZones(locale *Locale, cldr *CLDR, ldml *Ldml) {
	var timezones map[string]TimeZone = make(map[string]TimeZone)

	for _, t := range ldml.Dates.TimeZoneNames.Zone {
		if t.Type == "Etc/Unknown" || t.Type == "Etc/UTC" {
			continue
		}

		timezones[t.Type] = TimeZone{
			Code: t.Type,
			Name: t.ExemplarCity,
		}
	}

	locale.TimeZones = timezones
}
