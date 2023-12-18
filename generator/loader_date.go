// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type CalendarFormatter struct {
	Comment string
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
		validTypes := []string{"abbreviated", "wide", "narrow"}
		validFormats := []string{"format", "stand-alone"}

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
			Func := ParseDatePattern(date.DateFormat.Pattern)
			Func = strings.Replace(Func, "calendarSystem", calendar.Type, -1)

			locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
				Comment: "Pattern: " + date.DateFormat.Pattern,
				Func:    Func,
			}
		}

		for _, time := range calendar.TimeFormats.TimeFormatLength {
			key := fmt.Sprintf("time_%s", time.Type)
			Func := ParseDatePattern(time.TimeFormat.Pattern)
			Func = strings.Replace(Func, "calendarSystem", calendar.Type, -1)

			locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
				Comment: "Pattern: " + time.TimeFormat.Pattern,
				Func:    Func,
			}
		}

		// -- Load the period labels for the current locale, and then use them
		// to generate the function to get the period name based on the time
		// There are multiple period group: narrow, wide, etc ...
		if periods, ok := cldr.DayPeriods[locale.Code]; ok {
			for _, period := range calendar.DayPeriods.DayPeriodContext {
				if !slices.Contains(validFormats, period.Type) {
					continue
				}

				// iterate over defined set, and then the aliases
				for _, p := range period.DayPeriodWidth {
					baseKey := fmt.Sprintf("p_%s_%s", period.Type, p.Type)

					for name, includes := range map[string][]string{
						"a": {"am", "pm"},
						"b": {"am", "pm", "midnight", "noon"},
						"B": {"morning1", "morning2", "afternoon1", "afternoon2", "evening1", "evening2", "night1", "night2"},
					} {
						key := fmt.Sprintf("%s_%s", baseKey, name)
						// Yes, this part should resolve the alias to avoid
						// nested call, but it's not implemented yet
						if p.Alias.Path != "" {
							Func := GenAliasDatePeriodFunc(GetKeyAlias(p.Alias.Path, period.Type, p.Type, name))
							Func = strings.Replace(Func, "calendarSystem", calendar.Type, -1)

							locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
								Comment: "This is an alias to another configuration",
								Func:    Func,
							}

							continue
						}

						labels := map[string]string{}

						for _, l := range p.DayPeriod {
							labels[l.Type] = l.Text
						}

						Func := GenDatePeriodFunc(locale, periods, labels, includes)
						Func = strings.Replace(Func, "calendarSystem", calendar.Type, -1)
						Func = strings.Replace(Func, "formatSystem", key, -1)

						locale.Calendars[calendar.Type].Formatters[key] = CalendarFormatter{
							Comment: "No pattern defined, read the periods configuration",
							Func:    Func,
						}
					}
				}
			}
		}
	}
}

func SplitDatePattern(pattern string) (string, []string) {
	subset := []rune("")

	isLiteral := false
	var str = ""
	var params = []string{}
	for _, c := range pattern {
		// start a group literal, 39 is a quote
		if c == 39 && !isLiteral {
			isLiteral = true
			continue
		}

		// close a group literal
		if c == 39 && isLiteral {
			isLiteral = false
			continue
		}

		if isLiteral {
			str += string(c)
			continue
		}

		if c > 64 && c < 91 || c > 96 && c < 123 {
			subset = append(subset, c)
		} else { // new group
			s, p := GetPattern(string(subset))

			str += s
			if len(p) > 0 {
				params = append(params, p)
			}

			str += string(c)

			subset = []rune("")
		}
	}

	if len(subset) > 0 {
		s, p := GetPattern(string(subset))

		str += s
		if len(p) > 0 {
			params = append(params, p)
		}
	}

	return str, params
}

func ParseDatePattern(pattern string) string {
	str, params := SplitDatePattern(pattern)

	Func := `
	if tz, err := time.LoadLocation(timeZone); err != nil {
		panic(err)
	} else {
		t := tm.In(tz)
		`

	Func += "return fmt.Sprintf(\"" + str + "\", " + strings.Join(params, ", ") + ")"
	Func += "\n}"

	return Func
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
		return "%s", "l.GetCalendarLabels(\"calendarSystem\", \"m_format_wide\")[t.Month()-1]"
	case "MMM":
		return "%s", "l.GetCalendarLabels(\"calendarSystem\", \"m_format_abbreviated\")[t.Month()-1]"
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
		return "%s", "l.GetCalendarLabels(\"calendarSystem\", \"d_format_wide\")[t.Weekday()]"
	case "E":
		fallthrough
	case "EE":
		fallthrough
	case "EEE":
		fallthrough
	case "eee":
		return "%s", "l.GetCalendarLabels(\"calendarSystem\", \"d_format_abbreviated\")[t.Weekday()]"
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
	case "a":
		fallthrough
	case "aa":
		fallthrough
	case "aaa":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_abbreviated_a\")(tm, timeZone)"
	case "aaaa":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_wide_a\")(tm, timeZone)"
	case "aaaaa":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_narrow_a\")(tm, timeZone)"
	case "b":
		fallthrough
	case "bb":
		fallthrough
	case "bbb":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_abbreviated_b\")(tm, timeZone)"
	case "bbbb":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_wide_b\")(tm, timeZone)"
	case "bbbbb":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_narrow_b\")(tm, timeZone)"
	case "B":
		fallthrough
	case "BB":
		fallthrough
	case "BBB":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_abbreviated_B\")(tm, timeZone)"
	case "BBBB":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_wide_B\")(tm, timeZone)"
	case "BBBBB":
		return "%s", "l.GetCalendarFormatter(\"calendarSystem\", \"p_format_narrow_B\")(tm, timeZone)"
	}

	return pattern, ""
}

func GenDatePeriodFunc(locale *Locale, periods []*DayPeriodRule, labels map[string]string, includes []string) string {
	Func := ""
	once := false
	// first generate the equal rule
	for _, v := range periods {
		if !slices.Contains(includes, v.Type) {
			continue
		}

		if label, ok := labels[v.Type]; ok {
			once = true
			if v.At != -1 {
				Func += `if hour == ` + strconv.Itoa(v.At) + ` {
					return "` + label + `"
				}
			`
			} else if v.Before > v.From {
				Func += `if hour >= ` + strconv.Itoa(v.From) + ` && hour < ` + strconv.Itoa(v.Before) + ` {
					return "` + label + `"
				}
			`
			} else {
				Func += `if (hour >= ` + strconv.Itoa(v.From) + ` && hour < 2400) || (hour >= 0 && hour < ` + strconv.Itoa(v.Before) + `) {
					return "` + label + `"
				}
			`
			}
		}
	}

	if once {
		Func = `
		if tz, err := time.LoadLocation(timeZone); err != nil {
			panic(err)
		} else {
			// periods exist, use them
			t := tm.In(tz)
			hour := t.Hour()*100 + t.Minute()

		` + Func + `

		}
		`
	}

	if locale.Code == "root" {
		Func += `return tm.Format("PM")`
	} else {
		Func += `return l.Parent.GetCalendarFormatter("calendarSystem", "formatSystem")(tm, timeZone)`
	}

	return Func
}

func GetKeyAlias(alias, context, ptype, name string) string {
	aliases := ReadAlias(alias)

	if len(aliases) == 0 {
		panic("Unable to read/find alias: " + alias)
	}

	if len(aliases) == 1 {
		return fmt.Sprintf("p_%s_%s_%s", context, aliases[0], name)
	}

	if len(aliases) == 2 {
		return fmt.Sprintf("p_%s_%s_%s", aliases[0], aliases[1], name)
	}

	panic("Unsupported chain, please check alis: " + alias)
}

func GenAliasDatePeriodFunc(alias string) string {
	return "return l.Calendars[\"calendarSystem\"].Formatters[\"" + alias + "\"](tm, timeZone)"
}
