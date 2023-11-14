package golocales

type Territory string
type Currency string
type TimeZone string

type Locale struct {
	Name        string
	Territories map[string]Territory
	Currencies  map[string]Currency
	TimeZones   map[string]TimeZone
}
