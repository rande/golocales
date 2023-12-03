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
	"github.com/stretchr/testify/assert"
)

func Test_Time_Format(t *testing.T) {
	l_fr := fr.Locale()
	l_en := en.Locale()
	tm := time.Date(2015, 1, 30, 0, 0, 0, 0, time.UTC)

	formatter_fr := l_fr.Calendars["gregorian"].Formatters["date_short"]
	formatter_en := l_en.Calendars["gregorian"].Formatters["date_short"]

	assert.Equal(t, "30/01/2015", formatter_fr(tm, "Europe/Paris"))
	assert.Equal(t, "1/30/2015", formatter_en(tm, "Europe/Paris"))
}
