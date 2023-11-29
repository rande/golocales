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

type CalendarFormatter struct {
	Pattern string
	Func    string
}

type Calendar struct {
	// The calendar type
	System     string
	Labels     map[string][]string
	Formatters map[string]CalendarFormatter
}

func AttachCalendars(locale *Locale, cldr *CLDR, ldml *Ldml) {
	for _, calendar := range ldml.Dates.Calendars.Calendar {
		// only support gregorian calendar
		if calendar.Type != "gregorian" {
			continue
		}

		locale.Calendars[calendar.Type] = &Calendar{
			System:     calendar.Type,
			Labels:     map[string][]string{},
			Formatters: map[string]CalendarFormatter{},
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
			locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
				Pattern: date.DateFormat.Pattern,
				Func:    ParseDatePattern(date.DateFormat.Pattern),
			}
		}

		for _, time := range calendar.TimeFormats.TimeFormatLength {
			key := fmt.Sprintf("time_%s", time.Type)
			locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
				Pattern: time.TimeFormat.Pattern,
				Func:    ParseDatePattern(time.TimeFormat.Pattern),
			}
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
		return "%s", "string(t.Year())"
	case "YYYY":
		return "%s", "string(t.Year())"
	case "yyy":
		return "%03d", "t.Year()"
	case "YYY":
		return "%03d", "t.Year()"
	case "yy":
		return "%02d", "t.Year()"
	case "YY":
		return "%02d", "t.Year()"
	case "y":
		return "%d", "t.Year()"
	case "Y":
		return "%d", "t.Year()"
	case "MMMM":
		return "%s", "l.GetCalendarLabels(calendarSystem, \"m_format_wide\")[t.Month()-1]"
	case "MMM":
		return "%s", "l.GetCalendarLabels(calendarSystem, \"m_format_abbreviated\")[t.Month()-1]"
	case "MM":
		return "%02d", "t.Month()"
	case "M":
		return "%d", "t.Month()"
	case "d":
		return "%d", "t.Day()"
	case "dd":
		return "%02d", "t.Day()"
	case "D":
		return "%d", "t.YearDay()"
	case "DD":
		return "%02d", "t.YearDay()"
	case "DDD":
		return "%03d", "t.YearDay()"
	case "EEEE":
		fallthrough
	case "eeee":
		return "%s", "l.GetCalendarLabels(calendarSystem, \"d_format_wide\")[t.Weekday()]"
	case "E":
		fallthrough
	case "EE":
		fallthrough
	case "EEE":
		fallthrough
	case "eee":
		return "%s", "l.GetCalendarLabels(calendarSystem, \"d_format_abbreviated\")[t.Weekday()]"
	case "e":
		return "%d", "t.Weekday()"
	case "ee":
		return "%02d", "t.Weekday()"
	case "h":
		return "%d", "t.Hour()%12"
	case "hh":
		return "%02d", "t.Hour()%12"
	case "H":
		return "%d", "t.Hour()"
	case "HH":
		return "%02d", "t.Hour()"
	case "m":
		return "%d", "t.Minute()"
	case "mm":
		return "%02d", "t.Minute()"
	case "s":
		return "%d", "t.Second()"
	case "ss":
		return "%02d", "t.Second()"
	case "z":
		fallthrough
	case "zz":
		fallthrough
	case "zzz":
		fallthrough
	case "Z":
		fallthrough
	case "ZZ":
		fallthrough
	case "ZZZ":
		return "%s", "t.Format(\"-07:00\")"
	case "ZZZZ":
		return "%s", "t.Format(\"PST-07:00\")"
	case "zzzz":
		fallthrough
	case "ZZZZZ":
		return "%s", "t.Format(\"Z07:00:00\")"
	}

	return pattern, ""
}
