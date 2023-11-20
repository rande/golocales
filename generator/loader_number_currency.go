// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// <currencyFormats numberSystem="latn">
//   <currencyFormatLength> == default if not defined
//       <currencyFormat type="standard">
//           <pattern>#,##0.00 ¤</pattern>
//       </currencyFormat>
//       <currencyFormat type="accounting">
//           <pattern>#,##0.00 ¤;(#,##0.00 ¤)</pattern>
//           <pattern alt="noCurrency">#,##0.00;(#,##0.00)</pattern>
//       </currencyFormat>
//       <! -- Alternative with an alias --> Skipping to use the standard format
//       <currencyFormat type="accounting">
//          <alias source="locale" path="../currencyFormat[@type='standard']"/>
//       </currencyFormat>
//   </currencyFormatLength>
//   <currencyFormatLength type="short">
//       <currencyFormat type="standard">
//           <pattern type="1000" count="one">0 k ¤</pattern>
//           <pattern type="1000" count="other">0 k ¤</pattern>
//           <pattern type="10000" count="one">00 k ¤</pattern>
//           <pattern type="10000" count="other">00 k ¤</pattern>
//           <pattern type="100000" count="one">000 k ¤</pattern>
//          ...
//           <pattern type="10000000000000" count="other">00 Bn ¤</pattern>
//           <pattern type="100000000000000" count="one">000 Bn ¤</pattern>
//           <pattern type="100000000000000" count="other">000 Bn ¤</pattern>
//       </currencyFormat>
//   </currencyFormatLength>
//   <unitPattern count="one">{0} {1}</unitPattern>
//   <unitPattern count="other">{0} {1}</unitPattern>
// </currencyFormats>

func AttachNumberCurrencies(locale *Locale, cldr *CLDR, ldml *Ldml) {

	// <currencyFormats numberSystem="latn">
	for _, cfs := range ldml.Numbers.CurrencyFormats {
		// no symbol is defined, so we skip
		if cfs.NumberSystem == "" {
			continue
		}

		if cfs.Alias.Source != "" {
			continue
		}

		// <currencyFormatLength>
		for _, cfl := range cfs.CurrencyFormatLength {
			code := ifEmptyString(cfl.Type, "default")

			fmt.Printf("AttachNumberCurrencies: %s, %s\n", cfs.NumberSystem, code)

			// then we iterate over the patterns
			// <currencyFormat type="standard">
			for _, cf := range cfl.CurrencyFormat {

				if locale.Number.Currencies[cfs.NumberSystem] == nil {
					locale.Number.Currencies[cfs.NumberSystem] = FormatGroup{}
				}

				fmt.Printf("\nCurrencyFormat: %s \n", cf.Type)

				key := code + "_" + cf.Type

				for _, p := range cf.Pattern {
					format := &NumberFormat{
						Type:    p.Type,
						Count:   p.Count,
						Pattern: p.Text,
						Alt:     p.Alt,
					}

					AttachPattern(format)

					locale.Number.Currencies[cfs.NumberSystem][key] = append(locale.Number.Currencies[cfs.NumberSystem][key], format)
				}
			}
		}
	}
}
