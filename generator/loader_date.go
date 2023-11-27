// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

type Calendar struct {
	// The calendar type
	System string
	Labels map[string][]string
}

func AttachCalendars(locale *Locale, cldr *CLDR, ldml *Ldml) {

	for _, calendar := range ldml.Dates.Calendars.Calendar {
		// only support gregorian calendar
		if calendar.Type != "gregorian" {
			continue
		}

		locale.Calendars[calendar.Type] = &Calendar{
			System: calendar.Type,
			Labels: map[string][]string{},
		}
	}

	AttachLabels(locale, cldr, ldml)
}

func AttachLabels(locale *Locale, cldr *CLDR, ldml *Ldml) {

	for _, calendar := range ldml.Dates.Calendars.Calendar {

		// only support gregorian calendar
		if calendar.Type != "gregorian" {
			continue
		}

		for _, month := range calendar.Months.MonthContext {
			for _, m := range month.MonthWidth {
				key := fmt.Sprintf("m.%s.%s", month.Type, m.Type)
				locale.Calendars[calendar.Type].Labels[key] = []string{}
				for _, l := range m.Month {
					locale.Calendars[calendar.Type].Labels[key] = append(locale.Calendars[calendar.Type].Labels[key], l.Text)
				}
			}
		}

		for _, day := range calendar.Days.DayContext {
			for _, d := range day.DayWidth {
				key := fmt.Sprintf("d.%s.%s", day.Type, d.Type)
				locale.Calendars[calendar.Type].Labels[key] = []string{}
				for _, l := range d.Day {
					locale.Calendars[calendar.Type].Labels[key] = append(locale.Calendars[calendar.Type].Labels[key], l.Text)
				}
			}
		}
	}
}
