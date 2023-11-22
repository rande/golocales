// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

type CLDR struct {
	Path       string
	Validities []*Validity
	RootLocale *Locale
	MetaZones  []*MetaZone
}

func (cldr *CLDR) GetValidity(code, status string) *Validity {
	for _, v := range cldr.Validities {
		if v.From == code && v.Status == status {
			return v
		}
	}

	return nil
}
