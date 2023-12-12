// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import "encoding/xml"

type SupplementalData struct {
	XMLName xml.Name `xml:"supplementalData"`
	Text    string   `xml:",chardata"`
	Version struct {
		Text   string `xml:",chardata"`
		Number string `xml:"number,attr"`
	} `xml:"version"`
	IdValidity struct {
		Text string `xml:",chardata"`
		ID   []struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			IdStatus string `xml:"idStatus,attr"`
		} `xml:"id"`
	} `xml:"idValidity"`
	MetaZones struct {
		Text         string `xml:",chardata"`
		MetazoneInfo struct {
			Text     string `xml:",chardata"`
			Timezone []struct {
				Text         string `xml:",chardata"`
				Type         string `xml:"type,attr"`
				UsesMetazone []struct {
					Text  string `xml:",chardata"`
					Mzone string `xml:"mzone,attr"`
					To    string `xml:"to,attr"`
					From  string `xml:"from,attr"`
				} `xml:"usesMetazone"`
			} `xml:"timezone"`
		} `xml:"metazoneInfo"`
		MapTimezones struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			TypeVersion string `xml:"typeVersion,attr"`
			MapZone     []struct {
				Text      string `xml:",chardata"`
				Other     string `xml:"other,attr"`
				Territory string `xml:"territory,attr"`
				Type      string `xml:"type,attr"`
			} `xml:"mapZone"`
		} `xml:"mapTimezones"`
		MetazoneIds struct {
			Text       string `xml:",chardata"`
			MetazoneId []struct {
				Text    string `xml:",chardata"`
				ShortId string `xml:"shortId,attr"`
				LongId  string `xml:"longId,attr"`
			} `xml:"metazoneId"`
		} `xml:"metazoneIds"`
	} `xml:"metaZones"`
	PrimaryZones struct {
		Text        string `xml:",chardata"`
		PrimaryZone []struct {
			Text    string `xml:",chardata"`
			Iso3166 string `xml:"iso3166,attr"`
		} `xml:"primaryZone"`
	} `xml:"primaryZones"`

	// -- supplememtalData.xml
	CurrencyData struct {
		Text      string `xml:",chardata"`
		Fractions struct {
			Text string `xml:",chardata"`
			Info []struct {
				Text         string `xml:",chardata"`
				Iso4217      string `xml:"iso4217,attr"`
				Digits       string `xml:"digits,attr"`
				Rounding     string `xml:"rounding,attr"`
				CashDigits   string `xml:"cashDigits,attr"`
				CashRounding string `xml:"cashRounding,attr"`
			} `xml:"info"`
		} `xml:"fractions"`
		Region []struct {
			Text     string `xml:",chardata"`
			Iso3166  string `xml:"iso3166,attr"`
			Currency []struct {
				Text    string `xml:",chardata"`
				Iso4217 string `xml:"iso4217,attr"`
				From    string `xml:"from,attr"`
				To      string `xml:"to,attr"`
				Tender  string `xml:"tender,attr"`
			} `xml:"currency"`
		} `xml:"region"`
	} `xml:"currencyData"`
	TerritoryContainment struct {
		Text  string `xml:",chardata"`
		Group []struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			Contains string `xml:"contains,attr"`
			Status   string `xml:"status,attr"`
			Grouping string `xml:"grouping,attr"`
		} `xml:"group"`
	} `xml:"territoryContainment"`
	LanguageData struct {
		Text     string `xml:",chardata"`
		Language []struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Scripts     string `xml:"scripts,attr"`
			Territories string `xml:"territories,attr"`
			Alt         string `xml:"alt,attr"`
		} `xml:"language"`
	} `xml:"languageData"`
	TerritoryInfo struct {
		Text      string `xml:",chardata"`
		Territory []struct {
			Text               string `xml:",chardata"`
			Type               string `xml:"type,attr"`
			Gdp                string `xml:"gdp,attr"`
			LiteracyPercent    string `xml:"literacyPercent,attr"`
			Population         string `xml:"population,attr"`
			LanguagePopulation []struct {
				Text              string `xml:",chardata"`
				Type              string `xml:"type,attr"`
				PopulationPercent string `xml:"populationPercent,attr"`
				References        string `xml:"references,attr"`
				OfficialStatus    string `xml:"officialStatus,attr"`
				WritingPercent    string `xml:"writingPercent,attr"`
				LiteracyPercent   string `xml:"literacyPercent,attr"`
			} `xml:"languagePopulation"`
		} `xml:"territory"`
	} `xml:"territoryInfo"`
	CalendarData struct {
		Text     string `xml:",chardata"`
		Calendar []struct {
			Text           string `xml:",chardata"`
			Type           string `xml:"type,attr"`
			CalendarSystem struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"calendarSystem"`
			Eras struct {
				Text string `xml:",chardata"`
				Era  []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					End     string `xml:"end,attr"`
					Code    string `xml:"code,attr"`
					Aliases string `xml:"aliases,attr"`
					Start   string `xml:"start,attr"`
				} `xml:"era"`
			} `xml:"eras"`
			InheritEras struct {
				Text     string `xml:",chardata"`
				Calendar string `xml:"calendar,attr"`
			} `xml:"inheritEras"`
		} `xml:"calendar"`
	} `xml:"calendarData"`
	CalendarPreferenceData struct {
		Text               string `xml:",chardata"`
		CalendarPreference []struct {
			Text        string `xml:",chardata"`
			Territories string `xml:"territories,attr"`
			Ordering    string `xml:"ordering,attr"`
		} `xml:"calendarPreference"`
	} `xml:"calendarPreferenceData"`
	WeekData struct {
		Text    string `xml:",chardata"`
		MinDays []struct {
			Text        string `xml:",chardata"`
			Count       string `xml:"count,attr"`
			Territories string `xml:"territories,attr"`
		} `xml:"minDays"`
		FirstDay []struct {
			Text        string `xml:",chardata"`
			Day         string `xml:"day,attr"`
			Territories string `xml:"territories,attr"`
			Alt         string `xml:"alt,attr"`
			References  string `xml:"references,attr"`
		} `xml:"firstDay"`
		WeekendStart []struct {
			Text        string `xml:",chardata"`
			Day         string `xml:"day,attr"`
			Territories string `xml:"territories,attr"`
		} `xml:"weekendStart"`
		WeekendEnd []struct {
			Text        string `xml:",chardata"`
			Day         string `xml:"day,attr"`
			Territories string `xml:"territories,attr"`
		} `xml:"weekendEnd"`
		WeekOfPreference []struct {
			Text     string `xml:",chardata"`
			Ordering string `xml:"ordering,attr"`
			Locales  string `xml:"locales,attr"`
		} `xml:"weekOfPreference"`
	} `xml:"weekData"`
	TimeData struct {
		Text  string `xml:",chardata"`
		Hours []struct {
			Text      string `xml:",chardata"`
			Preferred string `xml:"preferred,attr"`
			Allowed   string `xml:"allowed,attr"`
			Regions   string `xml:"regions,attr"`
		} `xml:"hours"`
	} `xml:"timeData"`
	MeasurementData struct {
		Text              string `xml:",chardata"`
		MeasurementSystem []struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Territories string `xml:"territories,attr"`
			Category    string `xml:"category,attr"`
		} `xml:"measurementSystem"`
		PaperSize []struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Territories string `xml:"territories,attr"`
		} `xml:"paperSize"`
	} `xml:"measurementData"`
	CodeMappings struct {
		Text           string `xml:",chardata"`
		TerritoryCodes []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			Numeric string `xml:"numeric,attr"`
			Alpha3  string `xml:"alpha3,attr"`
			Fips10  string `xml:"fips10,attr"`
		} `xml:"territoryCodes"`
		CurrencyCodes []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			Numeric string `xml:"numeric,attr"`
		} `xml:"currencyCodes"`
	} `xml:"codeMappings"`
	ParentLocales []struct {
		Text         string `xml:",chardata"`
		Component    string `xml:"component,attr"`
		ParentLocale []struct {
			Text    string `xml:",chardata"`
			Parent  string `xml:"parent,attr"`
			Locales string `xml:"locales,attr"`
		} `xml:"parentLocale"`
	} `xml:"parentLocales"`
	PersonNamesDefaults struct {
		Text                    string `xml:",chardata"`
		NameOrderLocalesDefault []struct {
			Text  string `xml:",chardata"`
			Order string `xml:"order,attr"`
		} `xml:"nameOrderLocalesDefault"`
	} `xml:"personNamesDefaults"`
	References struct {
		Text      string `xml:",chardata"`
		Reference []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			URI  string `xml:"uri,attr"`
		} `xml:"reference"`
	} `xml:"references"`
}

type XmlAnnotation struct {
	XMLName  xml.Name `xml:"ldml"`
	Text     string   `xml:",chardata"`
	Identity struct {
		Text    string `xml:",chardata"`
		Version struct {
			Text   string `xml:",chardata"`
			Number string `xml:"number,attr"`
		} `xml:"version"`
		Language struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"language"`
	} `xml:"identity"`
	Annotations struct {
		Text       string `xml:",chardata"`
		Annotation []struct {
			Text  string `xml:",chardata"`
			Cp    string `xml:"cp,attr"`
			Type  string `xml:"type,attr"`
			Draft string `xml:"draft,attr"`
		} `xml:"annotation"`
	} `xml:"annotations"`
}

type Ldml struct {
	XMLName  xml.Name `xml:"ldml"`
	Text     string   `xml:",chardata"`
	Identity struct {
		Text    string `xml:",chardata"`
		Version struct {
			Text   string `xml:",chardata"`
			Number string `xml:"number,attr"`
		} `xml:"version"`
		Territory struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"territory"`
		Language struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"language"`
	} `xml:"identity"`
	LocaleDisplayNames struct {
		LocaleDisplayPattern struct {
			Text                 string `xml:",chardata"`
			LocaleKeyTypePattern string `xml:"localeKeyTypePattern"`
		} `xml:"localeDisplayPattern"`
		Languages struct {
			Text     string `xml:",chardata"`
			Language []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"language"`
		} `xml:"languages"`
		Scripts struct {
			Text   string `xml:",chardata"`
			Script []struct {
				Text  string `xml:",chardata"`
				Type  string `xml:"type,attr"`
				Draft string `xml:"draft,attr"`
				Alt   string `xml:"alt,attr"`
			} `xml:"script"`
		} `xml:"scripts"`
		Territories struct {
			Territory []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Alt  string `xml:"alt,attr"`
			} `xml:"territory"`
		} `xml:"territories"`
		Keys struct {
			Text string `xml:",chardata"`
			Key  []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"key"`
		} `xml:"keys"`
		Subdivisions struct {
			Text        string `xml:",chardata"`
			Subdivision []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"subdivision"`
		} `xml:"subdivisions"`
	} `xml:"localeDisplayNames"`
	ContextTransforms struct {
		Text                  string `xml:",chardata"`
		ContextTransformUsage []struct {
			Text             string `xml:",chardata"`
			Type             string `xml:"type,attr"`
			ContextTransform []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"contextTransform"`
		} `xml:"contextTransformUsage"`
	} `xml:"contextTransforms"`
	Characters struct {
		Text               string `xml:",chardata"`
		ExemplarCharacters []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"exemplarCharacters"`
		Ellipsis []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"ellipsis"`
		ParseLenients []struct {
			Text         string `xml:",chardata"`
			Scope        string `xml:"scope,attr"`
			Level        string `xml:"level,attr"`
			ParseLenient []struct {
				Text   string `xml:",chardata"`
				Sample string `xml:"sample,attr"`
			} `xml:"parseLenient"`
		} `xml:"parseLenients"`
	} `xml:"characters"`
	Delimiters struct {
		Text                    string `xml:",chardata"`
		QuotationStart          string `xml:"quotationStart"`
		QuotationEnd            string `xml:"quotationEnd"`
		AlternateQuotationStart string `xml:"alternateQuotationStart"`
		AlternateQuotationEnd   string `xml:"alternateQuotationEnd"`
	} `xml:"delimiters"`
	Dates struct {
		Text      string `xml:",chardata"`
		Calendars struct {
			Text     string `xml:",chardata"`
			Calendar []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Eras struct {
					Text     string `xml:",chardata"`
					EraNames struct {
						Text string `xml:",chardata"`
						Era  []struct {
							Text  string `xml:",chardata"`
							Type  string `xml:"type,attr"`
							Draft string `xml:"draft,attr"`
							Alt   string `xml:"alt,attr"`
						} `xml:"era"`
					} `xml:"eraNames"`
					EraAbbr struct {
						Text string `xml:",chardata"`
						Era  []struct {
							Text  string `xml:",chardata"`
							Type  string `xml:"type,attr"`
							Draft string `xml:"draft,attr"`
							Alt   string `xml:"alt,attr"`
						} `xml:"era"`
					} `xml:"eraAbbr"`
					EraNarrow struct {
						Text string `xml:",chardata"`
						Era  []struct {
							Text  string `xml:",chardata"`
							Type  string `xml:"type,attr"`
							Draft string `xml:"draft,attr"`
						} `xml:"era"`
					} `xml:"eraNarrow"`
				} `xml:"eras"`
				DateTimeFormats struct {
					Text             string `xml:",chardata"`
					AvailableFormats struct {
						Text           string `xml:",chardata"`
						DateFormatItem []struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Count string `xml:"count,attr"`
						} `xml:"dateFormatItem"`
					} `xml:"availableFormats"`
					IntervalFormats struct {
						Text               string `xml:",chardata"`
						IntervalFormatItem []struct {
							Text               string `xml:",chardata"`
							ID                 string `xml:"id,attr"`
							GreatestDifference []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Draft string `xml:"draft,attr"`
							} `xml:"greatestDifference"`
						} `xml:"intervalFormatItem"`
					} `xml:"intervalFormats"`
					DateTimeFormatLength []struct {
						Text           string `xml:",chardata"`
						Type           string `xml:"type,attr"`
						DateTimeFormat []struct {
							Text    string `xml:",chardata"`
							Type    string `xml:"type,attr"`
							Pattern string `xml:"pattern"`
						} `xml:"dateTimeFormat"`
					} `xml:"dateTimeFormatLength"`
				} `xml:"dateTimeFormats"`
				Months struct {
					Text         string `xml:",chardata"`
					MonthContext []struct {
						Text       string `xml:",chardata"`
						Type       string `xml:"type,attr"`
						MonthWidth []struct {
							Text  string `xml:",chardata"`
							Type  string `xml:"type,attr"`
							Month []struct {
								Text     string `xml:",chardata"`
								Type     string `xml:"type,attr"`
								Draft    string `xml:"draft,attr"`
								Yeartype string `xml:"yeartype,attr"`
							} `xml:"month"`
						} `xml:"monthWidth"`
					} `xml:"monthContext"`
				} `xml:"months"`
				DateFormats struct {
					Text             string `xml:",chardata"`
					DateFormatLength []struct {
						Text       string `xml:",chardata"`
						Type       string `xml:"type,attr"`
						DateFormat struct {
							Text             string `xml:",chardata"`
							Pattern          string `xml:"pattern"`
							DatetimeSkeleton string `xml:"datetimeSkeleton"`
						} `xml:"dateFormat"`
					} `xml:"dateFormatLength"`
				} `xml:"dateFormats"`
				Days struct {
					Text       string `xml:",chardata"`
					DayContext []struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						DayWidth []struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Day  []struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"day"`
						} `xml:"dayWidth"`
					} `xml:"dayContext"`
				} `xml:"days"`
				Quarters struct {
					Text           string `xml:",chardata"`
					QuarterContext struct {
						Text         string `xml:",chardata"`
						Type         string `xml:"type,attr"`
						QuarterWidth []struct {
							Text    string `xml:",chardata"`
							Type    string `xml:"type,attr"`
							Quarter []struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"quarter"`
						} `xml:"quarterWidth"`
					} `xml:"quarterContext"`
				} `xml:"quarters"`
				DayPeriods struct {
					Text             string `xml:",chardata"`
					DayPeriodContext []struct {
						Text           string `xml:",chardata"`
						Type           string `xml:"type,attr"`
						DayPeriodWidth []struct {
							Text      string `xml:",chardata"`
							Type      string `xml:"type,attr"`
							DayPeriod []struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"dayPeriod"`
						} `xml:"dayPeriodWidth"`
					} `xml:"dayPeriodContext"`
				} `xml:"dayPeriods"`
				TimeFormats struct {
					Text             string `xml:",chardata"`
					TimeFormatLength []struct {
						Text       string `xml:",chardata"`
						Type       string `xml:"type,attr"`
						TimeFormat struct {
							Text    string `xml:",chardata"`
							Pattern string `xml:"pattern"`
						} `xml:"timeFormat"`
					} `xml:"timeFormatLength"`
				} `xml:"timeFormats"`
			} `xml:"calendar"`
		} `xml:"calendars"`
		Fields struct {
			Text  string `xml:",chardata"`
			Field []struct {
				Text        string `xml:",chardata"`
				Type        string `xml:"type,attr"`
				DisplayName string `xml:"displayName"`
				Relative    []struct {
					Text  string `xml:",chardata"`
					Type  string `xml:"type,attr"`
					Draft string `xml:"draft,attr"`
				} `xml:"relative"`
				RelativeTime []struct {
					Text                string `xml:",chardata"`
					Type                string `xml:"type,attr"`
					RelativeTimePattern []struct {
						Text  string `xml:",chardata"`
						Count string `xml:"count,attr"`
					} `xml:"relativeTimePattern"`
				} `xml:"relativeTime"`
				RelativePeriod string `xml:"relativePeriod"`
			} `xml:"field"`
		} `xml:"fields"`
		TimeZoneNames struct {
			Text          string `xml:",chardata"`
			HourFormat    string `xml:"hourFormat"`
			GmtFormat     string `xml:"gmtFormat"`
			GmtZeroFormat string `xml:"gmtZeroFormat"`
			RegionFormat  []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"regionFormat"`
			Zone []struct {
				Text         string `xml:",chardata"`
				Type         string `xml:"type,attr"`
				ExemplarCity struct {
					Text string `xml:",chardata"`
				} `xml:"exemplarCity"`
				Long struct {
					Text     string `xml:",chardata"`
					Standard string `xml:"standard"`
					Daylight string `xml:"daylight"`
				} `xml:"long"`
				Short struct {
					Text     string `xml:",chardata"`
					Standard struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"standard"`
					Generic struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"generic"`
					Daylight struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"daylight"`
				} `xml:"short"`
			} `xml:"zone"`
			Metazone []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Long struct {
					Text    string `xml:",chardata"`
					Generic struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"generic"`
					Standard struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"standard"`
					Daylight struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"daylight"`
				} `xml:"long"`
				Short struct {
					Text    string `xml:",chardata"`
					Generic struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"generic"`
					Standard struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"standard"`
					Daylight struct {
						Text  string `xml:",chardata"`
						Draft string `xml:"draft,attr"`
					} `xml:"daylight"`
				} `xml:"short"`
			} `xml:"metazone"`
		} `xml:"timeZoneNames"`
	} `xml:"dates"`
	Numbers struct {
		Text                   string `xml:",chardata"`
		MinimumGroupingDigits  string `xml:"minimumGroupingDigits"`
		DefaultNumberingSystem string `xml:"defaultNumberingSystem"`
		Symbols                []struct {
			Text         string `xml:",chardata"`
			NumberSystem string `xml:"numberSystem,attr"`
			Alias        struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Path   string `xml:"path,attr"`
			} `xml:"alias"`
			Decimal                string `xml:"decimal"`
			Group                  string `xml:"group"`
			CurrencyGroup          string `xml:"currencyGroup"`
			List                   string `xml:"list"`
			PercentSign            string `xml:"percentSign"`
			PlusSign               string `xml:"plusSign"`
			MinusSign              string `xml:"minusSign"`
			ApproximatelySign      string `xml:"approximatelySign"`
			Exponential            string `xml:"exponential"`
			SuperscriptingExponent string `xml:"superscriptingExponent"`
			PerMille               string `xml:"perMille"`
			Infinity               string `xml:"infinity"`
			Nan                    string `xml:"nan"`
			TimeSeparator          string `xml:"timeSeparator"`
		} `xml:"symbols"`
		DecimalFormats []struct {
			Text         string `xml:",chardata"`
			NumberSystem string `xml:"numberSystem,attr"`
			Alias        struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Path   string `xml:"path,attr"`
			} `xml:"alias"`
			DecimalFormatLength []struct {
				Text          string `xml:",chardata"`
				Type          string `xml:"type,attr"`
				DecimalFormat struct {
					Text    string `xml:",chardata"`
					Pattern []struct {
						Text  string `xml:",chardata"`
						Type  string `xml:"type,attr"`
						Count string `xml:"count,attr"`
					} `xml:"pattern"`
				} `xml:"decimalFormat"`
				Alias struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Path   string `xml:"path,attr"`
				} `xml:"alias"`
			} `xml:"decimalFormatLength"`
		} `xml:"decimalFormats"`
		PercentFormats struct {
			Text                string `xml:",chardata"`
			NumberSystem        string `xml:"numberSystem,attr"`
			PercentFormatLength struct {
				Text          string `xml:",chardata"`
				PercentFormat struct {
					Text    string `xml:",chardata"`
					Pattern string `xml:"pattern"`
				} `xml:"percentFormat"`
			} `xml:"percentFormatLength"`
		} `xml:"percentFormats"`
		CurrencyFormats []struct {
			Text         string `xml:",chardata"`
			NumberSystem string `xml:"numberSystem,attr"`
			Alias        struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Path   string `xml:"path,attr"`
			} `xml:"alias"`
			CurrencyFormatLength []struct {
				Text           string `xml:",chardata"`
				Type           string `xml:"type,attr"`
				CurrencyFormat []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Pattern []struct {
						Text  string `xml:",chardata"`
						Alt   string `xml:"alt,attr"`
						Type  string `xml:"type,attr"`
						Count string `xml:"count,attr"`
					} `xml:"pattern"`
				} `xml:"currencyFormat"`
			} `xml:"currencyFormatLength"`
			UnitPattern []struct {
				Text  string `xml:",chardata"`
				Count string `xml:"count,attr"`
			} `xml:"unitPattern"`
		} `xml:"currencyFormats"`
		Currencies struct {
			Text     string `xml:",chardata"`
			Currency []struct {
				Text        string `xml:",chardata"`
				Type        string `xml:"type,attr"`
				DisplayName []struct {
					Text  string `xml:",chardata"`
					Count string `xml:"count,attr"`
					Draft string `xml:"draft,attr"`
				} `xml:"displayName"`
				Symbol []struct {
					Text  string `xml:",chardata"`
					Alt   string `xml:"alt,attr"`
					Draft string `xml:"draft,attr"`
				} `xml:"symbol"`
			} `xml:"currency"`
		} `xml:"currencies"`
		MiscPatterns struct {
			Text         string `xml:",chardata"`
			NumberSystem string `xml:"numberSystem,attr"`
			Pattern      struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"pattern"`
		} `xml:"miscPatterns"`
		MinimalPairs struct {
			Text               string `xml:",chardata"`
			PluralMinimalPairs []struct {
				Text  string `xml:",chardata"`
				Count string `xml:"count,attr"`
			} `xml:"pluralMinimalPairs"`
			OrdinalMinimalPairs []struct {
				Text    string `xml:",chardata"`
				Ordinal string `xml:"ordinal,attr"`
			} `xml:"ordinalMinimalPairs"`
			GenderMinimalPairs []struct {
				Text   string `xml:",chardata"`
				Gender string `xml:"gender,attr"`
			} `xml:"genderMinimalPairs"`
		} `xml:"minimalPairs"`
	} `xml:"numbers"`
	Units struct {
		Text       string `xml:",chardata"`
		UnitLength []struct {
			Text         string `xml:",chardata"`
			Type         string `xml:"type,attr"`
			CompoundUnit []struct {
				Text                 string `xml:",chardata"`
				Type                 string `xml:"type,attr"`
				UnitPrefixPattern    string `xml:"unitPrefixPattern"`
				CompoundUnitPattern  string `xml:"compoundUnitPattern"`
				CompoundUnitPattern1 []struct {
					Text   string `xml:",chardata"`
					Count  string `xml:"count,attr"`
					Gender string `xml:"gender,attr"`
				} `xml:"compoundUnitPattern1"`
			} `xml:"compoundUnit"`
			Unit []struct {
				Text        string `xml:",chardata"`
				Type        string `xml:"type,attr"`
				Gender      string `xml:"gender"`
				DisplayName string `xml:"displayName"`
				UnitPattern []struct {
					Text  string `xml:",chardata"`
					Count string `xml:"count,attr"`
				} `xml:"unitPattern"`
				PerUnitPattern string `xml:"perUnitPattern"`
			} `xml:"unit"`
			CoordinateUnit struct {
				Text                  string `xml:",chardata"`
				CoordinateUnitPattern []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"coordinateUnitPattern"`
			} `xml:"coordinateUnit"`
		} `xml:"unitLength"`
	} `xml:"units"`
	ListPatterns struct {
		Text        string `xml:",chardata"`
		ListPattern []struct {
			Text            string `xml:",chardata"`
			Type            string `xml:"type,attr"`
			ListPatternPart []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"listPatternPart"`
		} `xml:"listPattern"`
	} `xml:"listPatterns"`
	Posix struct {
		Text     string `xml:",chardata"`
		Messages struct {
			Text   string `xml:",chardata"`
			Yesstr string `xml:"yesstr"`
			Nostr  string `xml:"nostr"`
		} `xml:"messages"`
	} `xml:"posix"`
	CharacterLabels struct {
		Text                  string `xml:",chardata"`
		CharacterLabelPattern []struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Count string `xml:"count,attr"`
		} `xml:"characterLabelPattern"`
		CharacterLabel []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"characterLabel"`
	} `xml:"characterLabels"`
	TypographicNames struct {
		Text     string `xml:",chardata"`
		AxisName []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"axisName"`
		StyleName []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			Subtype string `xml:"subtype,attr"`
			Alt     string `xml:"alt,attr"`
		} `xml:"styleName"`
		FeatureName []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"featureName"`
	} `xml:"typographicNames"`
	PersonNames struct {
		Text             string `xml:",chardata"`
		NameOrderLocales []struct {
			Text  string `xml:",chardata"`
			Order string `xml:"order,attr"`
		} `xml:"nameOrderLocales"`
		PersonName []struct {
			Text        string `xml:",chardata"`
			Order       string `xml:"order,attr"`
			Length      string `xml:"length,attr"`
			Usage       string `xml:"usage,attr"`
			Formality   string `xml:"formality,attr"`
			NamePattern string `xml:"namePattern"`
		} `xml:"personName"`
		SampleName []struct {
			Text      string `xml:",chardata"`
			Item      string `xml:"item,attr"`
			NameField []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"nameField"`
		} `xml:"sampleName"`
	} `xml:"personNames"`
}
