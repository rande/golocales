package golocales

type Territory string
type Currency string

type Locale struct {
	Name        string
	Territories map[string]Territory
	Currencies  map[string]Currency
}
