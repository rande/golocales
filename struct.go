package golocales

type Territory string

type Locale struct {
	Name        string
	Territories map[string]Territory
}
