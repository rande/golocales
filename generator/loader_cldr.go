package main

type CLDR struct {
	Validities []*Validity
	RootLocale *Locale
}

func (cldr *CLDR) GetValidity(code, status string) *Validity {
	for _, v := range cldr.Validities {
		if v.From == code && v.Status == status {
			return v
		}
	}

	return nil
}
