// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"slices"
	"strings"
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

		// only support abbreviated and wide, ie: Sun and Sunday
		validTypes := []string{"abbreviated", "wide" /*, "narrow" */}
		validFormats := []string{"format" /*, "stand-alone" */}

		for _, month := range calendar.Months.MonthContext {
			if !slices.Contains(validFormats, month.Type) {
				continue
			}

			for _, m := range month.MonthWidth {
				if !slices.Contains(validTypes, m.Type) {
					continue
				}

				key := fmt.Sprintf("m_%s_%s", month.Type, m.Type)
				locale.Calendars[calendar.Type].Labels[key] = []string{}
				for _, l := range m.Month {
					locale.Calendars[calendar.Type].Labels[key] = append(locale.Calendars[calendar.Type].Labels[key], l.Text)
				}
			}
		}

		for _, day := range calendar.Days.DayContext {
			if !slices.Contains(validFormats, day.Type) {
				continue
			}

			for _, d := range day.DayWidth {
				if !slices.Contains(validTypes, d.Type) {
					continue
				}

				key := fmt.Sprintf("d_%s_%s", day.Type, d.Type)
				locale.Calendars[calendar.Type].Labels[key] = []string{}
				for _, l := range d.Day {
					locale.Calendars[calendar.Type].Labels[key] = append(locale.Calendars[calendar.Type].Labels[key], l.Text)
				}
			}
		}

		for _, date := range calendar.DateFormats.DateFormatLength {
			key := fmt.Sprintf("date_%s", date.Type)
			locale.Calendars[calendar.Type].Labels[key] = []string{date.DateFormat.Pattern}
		}

		for _, time := range calendar.TimeFormats.TimeFormatLength {
			key := fmt.Sprintf("time_%s", time.Type)
			locale.Calendars[calendar.Type].Labels[key] = []string{time.TimeFormat.Pattern}
		}
	}
}

func SplitDatePattern(pattern string) []string {
	letter := rune(-1)
	groups := []string{}

	inc := 0
	for _, c := range pattern {
		if letter == rune(-1) {
			letter = c
			groups = append(groups, "")
		}

		if letter != c || (c < 65 && c > 90 && c < 97 && c > 122) { // new group
			inc++
			groups = append(groups, "")
		}

		letter = c
		groups[inc] += string(c)
	}

	return groups
}

func ParseDatePattern(pattern string) string {
	groups := SplitDatePattern(pattern)

	var str = ""
	var params = []string{}
	for _, group := range groups {
		s, p := GetPattern(string(group))

		str += s
		if len(p) > 0 {
			params = append(params, p)
		}
	}

	return "fmt.Sprintf(\"" + str + "\", " + strings.Join(params, ", ") + ")"
}

func GetPattern(pattern string) (string, string) {

	// http://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
	// http://www.unicode.org/reports/tr35/tr35-dates.html#Date_Format_Patterns

	switch pattern {
	case "yyyy":
		return "%s", "string(time.Year())"
	case "YYYY":
		return "%s", "string(time.Year())"
	case "yyy":
		return "%03d", "time.Year()"
	case "YYY":
		return "%03d", "time.Year()"
	case "yy":
		return "%02d", "time.Year()"
	case "YY":
		return "%02d", "time.Year()"
	case "y":
		return "%d", "time.Year()"
	case "Y":
		return "%d", "time.Year()"
	case "MMMM":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"m_format_wide\"][time.Month()-1]"
	case "MMM":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"m_format_abbreviated\"][time.Month()-1]"
	case "MM":
		return "%02d", "time.Month()"
	case "M":
		return "%d", "time.Month()"
	case "d":
		return "%d", "time.Day()"
	case "dd":
		return "%02d", "time.Day()"
	case "D":
		return "%d", "time.YearDay()"
	case "DD":
		return "%02d", "time.YearDay()"
	case "DDD":
		return "%03d", "time.YearDay()"
	case "EEEE":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_wide\"][time.Weekday()]"
	case "eeee":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_wide\"][time.Weekday()]"
	case "E":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_abbreviated\"][time.Weekday()]"
	case "EE":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_abbreviated\"][time.Weekday()]"
	case "EEE":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_abbreviated\"][time.Weekday()]"
	case "eee":
		return "%s", "locale.Calendars[\"gregorian\"].Labels[\"d_format_abbreviated\"][time.Weekday()]"
	case "e":
		return "%d", "time.Weekday()"
	case "ee":
		return "%02d", "time.Weekday()"
	case "h":
		return "%d", "time.Hour()%12"
	case "hh":
		return "%02d", "time.Hour()%12"
	case "H":
		return "%d", "time.Hour()"
	case "HH":
		return "%02d", "time.Hour()"
	case "m":
		return "%d", "time.Minute()"
	case "mm":
		return "%02d", "time.Minute()"
	case "s":
		return "%d", "time.Second()"
	case "ss":
		return "%02d", "time.Second()"
	}

	return pattern, ""
}