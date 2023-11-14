package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"slices"
	"strings"
	"text/template"
)

func LoadXml(filename string, strct interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err = xml.Unmarshal(data, strct); err != nil {
		return err
	}

	return nil
}

func LoadLdml(filename string) (*Ldml, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ldml Ldml
	err = xml.Unmarshal(data, &ldml)
	if err != nil {
		return nil, err
	}

	return &ldml, nil
}

func BuildLocale(cldr *CLDR, ldml *Ldml) *Locale {

	locale := &Locale{
		Code:        ldml.Identity.Language.Type,
		Name:        ldml.Identity.Language.Type,
		Territory:   ldml.Identity.Territory.Type,
		Territories: GetTerritories(cldr, ldml),
		Currencies:  GetCurrencies(cldr, ldml),
		TimeZones:   GetTimeZones(cldr, ldml),
	}

	if locale.IsTerritory() {
		locale.Code = locale.Code + "_" + locale.Territory
	}

	return locale
}

func AttachValidity(cldr *CLDR, supplemental *SupplementalData) {
	for _, id := range supplemental.IdValidity.ID {
		v := &Validity{
			From:   id.Type,
			Status: id.IdStatus,
			List:   ParseValidityValues(id.Text),
		}

		cldr.Validities = append(cldr.Validities, v)
	}
}

func WriteGo(locale *Locale, w io.Writer) error {
	fs := GetEmbedFS()

	tpl := template.Must(template.ParseFS(fs, "templates/data.tmpl"))

	ctx := map[string]interface{}{
		"Locale":      locale.Code,
		"Territories": locale.Territories,
		"Currencies":  locale.Currencies,
		"TimeZones":   locale.TimeZones,
	}

	return tpl.Execute(w, ctx)
}

type CLDR struct {
	Validities []*Validity
}

func (cldr *CLDR) GetValidity(code, status string) *Validity {
	for _, v := range cldr.Validities {
		if v.From == code && v.Status == status {
			return v
		}
	}

	return nil
}

type Locale struct {
	Code        string
	Name        string
	Territory   string
	Territories map[string]Territory
	Currencies  map[string]Currency
	TimeZones   map[string]TimeZone
}

func (locale *Locale) IsTerritory() bool {
	return locale.Territory != ""
}

func (locale *Locale) IsRoot() bool {
	return !locale.IsTerritory()
}

type Territory struct {
	Code string
	Name string
	Alt  string
}

type Currency struct {
	Code string
	Name string
}

type TimeZone struct {
	Code string
	Name string
}

type Validity struct {
	From   string
	List   []string
	Status string
}

var TerritoriesDenyList = map[string]bool{
	// Exceptional reservations
	"AC": true, // Ascension Island
	"CP": true, // Clipperton Island
	"CQ": true, // Island of Sark
	"DG": true, // Diego Garcia
	"EA": true, // Ceuta & Melilla
	"EU": true, // European Union
	"EZ": true, // Eurozone
	"IC": true, // Canary Islands
	"TA": true, // Tristan da Cunha
	"UN": true, // United Nations
	// User-assigned
	"QO": true, // Outlying Oceania
	"XA": true, // Pseudo-Accents
	"XB": true, // Pseudo-Bidi
	"XK": true, // Kosovo
	// Misc
	"ZZ": true, // Unknown Region
}

// @see https://en.wikipedia.org/wiki/ISO_3166-1_numeric#Withdrawn_codes
var TerritoriesWithdrawnCodes = []string{
	"128", //	Canton and Enderbury Islands
	"200", //	Czechoslovakia
	"216", //	Dronning Maud Land
	"230", //	Ethiopia
	"249", //	France, Metropolitan
	"278", //	German Democratic Republic
	"280", //	Germany, Federal Republic of
	"396", //	Johnston Island
	"488", //	Midway Islands
	"530", //	Netherlands Antilles
	"532", //	Netherlands Antilles
	"536", //	Neutral Zone
	"582", //	Pacific Islands (Trust Territory)
	"590", //	Panama
	"658", //	Saint Kitts-Nevis-Anguilla
	"720", //	Yemen, Democratic
	"736", //	Sudan
	"810", //	USSR
	"849", //	United States Miscellaneous Pacific Islands
	"872", //	Wake Island
	"886", //	Yemen Arab Republic
	"890", //	Yugoslavia, Socialist Federal Republic of
	"891", //	Serbia and Montenegro
}

func GetTerritories(cldr *CLDR, ldml *Ldml) map[string]Territory {
	var territories map[string]Territory = make(map[string]Territory)

	list := cldr.GetValidity("region", "regular")

	if list == nil {
		fmt.Printf("No regions found\n")
		return territories
	}

	for _, t := range ldml.LocaleDisplayNames.Territories.Territory {
		// if _, ok := TerritoriesDenyList[t.Type]; ok {
		// 	continue
		// }

		if !slices.Contains(list.List, strings.ToUpper(t.Type)) {
			continue
		}

		// if slices.Contains(TerritoriesWithdrawnCodes, t.Type) {
		// 	continue
		// }

		// we don't keep variants or short names.
		if t.Alt != "" {
			continue
		}

		// if _, err := strconv.Atoi(t.Type); err == nil {
		// 	continue
		// }

		territories[t.Type] = Territory{
			Code: t.Type,
			Name: t.Text,
			Alt:  t.Alt,
		}
	}

	return territories
}

func GetCurrencies(cldr *CLDR, ldml *Ldml) map[string]Currency {
	var currencies map[string]Currency = make(map[string]Currency)

	list := cldr.GetValidity("currency", "regular")

	if list == nil {
		fmt.Printf("No currencies found\n")
		return currencies
	}

	for _, t := range ldml.Numbers.Currencies.Currency {
		// if _, ok := TerritoriesDenyList[t.Type]; ok {
		// 	continue
		// }

		if !slices.Contains(list.List, strings.ToUpper(t.Type)) {
			continue
		}

		name := ""

		for _, displayName := range t.DisplayName {
			if displayName.Count != "" {
				continue
			}

			name = displayName.Text
		}

		currencies[t.Type] = Currency{
			Code: t.Type,
			Name: name,
		}
	}

	return currencies
}

func GetTimeZones(cldr *CLDR, ldml *Ldml) map[string]TimeZone {
	var timezones map[string]TimeZone = make(map[string]TimeZone)

	for _, t := range ldml.Dates.TimeZoneNames.Zone {
		if t.Type == "Etc/Unknown" || t.Type == "Etc/UTC" {
			continue
		}

		timezones[t.Type] = TimeZone{
			Code: t.Type,
			Name: t.ExemplarCity,
		}
	}

	return timezones
}
