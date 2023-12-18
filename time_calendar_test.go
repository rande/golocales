// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package golocales_test

import (
	"testing"
	"time"

	"github.com/rande/golocales/locales/en"
	"github.com/rande/golocales/locales/fr"
	"github.com/rande/golocales/locales/hi"
	"github.com/stretchr/testify/assert"
)

func Test_Calendar_DateFormat(t *testing.T) {
	l_fr := fr.GetLocale()
	l_en := en.GetLocale()
	tm := time.Date(2015, 1, 30, 0, 0, 0, 0, time.UTC)

	formatter_fr := l_fr.Calendars["gregorian"].Formatters["date_short"]
	formatter_en := l_en.Calendars["gregorian"].Formatters["date_short"]

	assert.Equal(t, "30/01/2015", formatter_fr(tm, "Europe/Paris"))
	assert.Equal(t, "1/30/2015", formatter_en(tm, "Europe/Paris"))
}

func Test_Calendar_TimeFormat(t *testing.T) {
	l_fr := fr.GetLocale()
	l_en := en.GetLocale()
	tm := time.Date(2015, 1, 30, 1, 1, 1, 0, time.UTC)

	formatter_fr := l_fr.GetCalendarFormatter("gregorian", "time_short")
	formatter_en := l_en.GetCalendarFormatter("gregorian", "time_short")

	assert.Equal(t, "02:01", formatter_fr(tm, "Europe/Paris"))
	assert.Equal(t, "2:01 am", formatter_en(tm, "Europe/Paris"))
}

func Test_Calendar_TimeShortFormat(t *testing.T) {
	l_hi := hi.GetLocale()

	tm := time.Date(2015, 1, 30, 1, 1, 1, 0, time.UTC)

	formatter := l_hi.GetCalendarFormatter("gregorian", "time_short")

	assert.Equal(t, "2:01 am", formatter(tm, "Europe/Paris"))
}
