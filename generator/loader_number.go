// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

const DefaultNumberSystem = "latn"

func AttachNumber(locale *Locale, cldr *CLDR, ldml *Ldml) {

	locale.MinimumGroupingDigits = ifEmptyInt(ldml.Numbers.MinimumGroupingDigits, "1")
	locale.DefaultNumberSystem = ldml.Numbers.DefaultNumberingSystem

	// Parent must have a valid configuration
	if locale.Parent != nil {
		locale.DefaultNumberSystem = locale.Parent.DefaultNumberSystem
	}

	AttachNumberSymbols(locale, cldr, ldml)
	AttachNumberDecimals(locale, cldr, ldml)
}
