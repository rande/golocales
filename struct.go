package golocales

type Territory string
type Currency string
type TimeZone string
type NumberSystem string
type Locale struct {
	Name        string
	Territories map[string]Territory
	Currencies  map[string]Currency
	TimeZones   map[string]TimeZone
	Numbers     map[NumberSystem]*Number
	Decimal     map[NumberSystem]FormatGroup
	Parent      *Locale
}

type Number struct {
	System                 string
	MinusSign              string
	PlusSign               string
	Exponential            string
	SuperscriptingExponent string
	Decimal                string
	Group                  string
	PercentSign            string
	ApproximatelySign      string
	Infinity               string
	TimeSeparator          string
}

type FormatGroup struct {
	Default []*NumberFormat
	Long    []*NumberFormat
	Short   []*NumberFormat
}

type NumberFormat struct {
	Type    string
	Count   string
	Pattern string
}
