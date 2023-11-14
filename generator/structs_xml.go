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
				ExemplarCity string `xml:"exemplarCity"`
				Long         struct {
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
		Text    string `xml:",chardata"`
		Symbols []struct {
			Text         string `xml:",chardata"`
			NumberSystem string `xml:"numberSystem,attr"`
			PercentSign  struct {
				Text  string `xml:",chardata"`
				Draft string `xml:"draft,attr"`
			} `xml:"percentSign"`
			PlusSign struct {
				Text  string `xml:",chardata"`
				Draft string `xml:"draft,attr"`
			} `xml:"plusSign"`
			MinusSign struct {
				Text  string `xml:",chardata"`
				Draft string `xml:"draft,attr"`
			} `xml:"minusSign"`
			ApproximatelySign struct {
				Text  string `xml:",chardata"`
				Draft string `xml:"draft,attr"`
			} `xml:"approximatelySign"`
			Decimal string `xml:"decimal"`
			Group   string `xml:"group"`
		} `xml:"symbols"`
		DecimalFormats struct {
			Text                string `xml:",chardata"`
			NumberSystem        string `xml:"numberSystem,attr"`
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
		CurrencyFormats struct {
			Text                 string `xml:",chardata"`
			NumberSystem         string `xml:"numberSystem,attr"`
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
