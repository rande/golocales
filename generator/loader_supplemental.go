// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

type MetaZone struct {
	Type string
	Zone string
}

func AttachSupplemental(cldr *CLDR, supplemental *SupplementalData) {
	AttachValidity(cldr, supplemental)
	AttachMetaZones(cldr, supplemental)
}

func AttachMetaZones(cldr *CLDR, supplemental *SupplementalData) {

	for _, t := range supplemental.MetaZones.MetazoneInfo.Timezone {
		last := t.UsesMetazone[len(t.UsesMetazone)-1]

		meta := &MetaZone{
			Type: t.Type,
			Zone: last.Mzone,
		}

		cldr.MetaZones = append(cldr.MetaZones, meta)
	}
}
