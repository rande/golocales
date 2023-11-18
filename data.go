// Copyright © 2023 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file is an adaptation of https://github.com/bojanz/currency
// all credits goes to Bojan Zivanovic and contributors

package golocales

// TODO: Generate this file
type currencyInfo struct {
	numericCode string
	digits      uint8
}

var currencies = map[string]currencyInfo{
	"AED": {"784", 2}, "AFN": {"971", 0}, "ALL": {"008", 0},
	"AMD": {"051", 2}, "ANG": {"532", 2}, "AOA": {"973", 2},
	"ARS": {"032", 2}, "AUD": {"036", 2}, "AWG": {"533", 2},
	"AZN": {"944", 2}, "BAM": {"977", 2}, "BBD": {"052", 2},
	"BDT": {"050", 2}, "BGN": {"975", 2}, "BHD": {"048", 3},
	"BIF": {"108", 0}, "BMD": {"060", 2}, "BND": {"096", 2},
	"BOB": {"068", 2}, "BOV": {"984", 2}, "BRL": {"986", 2},
	"BSD": {"044", 2}, "BTN": {"064", 2}, "BWP": {"072", 2},
	"BYN": {"933", 2}, "BZD": {"084", 2}, "CAD": {"124", 2},
	"CDF": {"976", 2}, "CHE": {"947", 2}, "CHF": {"756", 2},
	"CHW": {"948", 2}, "CLF": {"990", 4}, "CLP": {"152", 0},
	"CNY": {"156", 2}, "COP": {"170", 2}, "COU": {"970", 2},
	"CRC": {"188", 2}, "CUC": {"931", 2}, "CUP": {"192", 2},
	"CVE": {"132", 2}, "CZK": {"203", 2}, "DJF": {"262", 0},
	"DKK": {"208", 2}, "DOP": {"214", 2}, "DZD": {"012", 2},
	"EGP": {"818", 2}, "ERN": {"232", 2}, "ETB": {"230", 2},
	"EUR": {"978", 2}, "FJD": {"242", 2}, "FKP": {"238", 2},
	"GBP": {"826", 2}, "GEL": {"981", 2}, "GHS": {"936", 2},
	"GIP": {"292", 2}, "GMD": {"270", 2}, "GNF": {"324", 0},
	"GTQ": {"320", 2}, "GYD": {"328", 2}, "HKD": {"344", 2},
	"HNL": {"340", 2}, "HTG": {"332", 2}, "HUF": {"348", 2},
	"IDR": {"360", 2}, "ILS": {"376", 2}, "INR": {"356", 2},
	"IQD": {"368", 0}, "IRR": {"364", 0}, "ISK": {"352", 0},
	"JMD": {"388", 2}, "JOD": {"400", 3}, "JPY": {"392", 0},
	"KES": {"404", 2}, "KGS": {"417", 2}, "KHR": {"116", 2},
	"KMF": {"174", 0}, "KPW": {"408", 0}, "KRW": {"410", 0},
	"KWD": {"414", 3}, "KYD": {"136", 2}, "KZT": {"398", 2},
	"LAK": {"418", 0}, "LBP": {"422", 0}, "LKR": {"144", 2},
	"LRD": {"430", 2}, "LSL": {"426", 2}, "LYD": {"434", 3},
	"MAD": {"504", 2}, "MDL": {"498", 2}, "MGA": {"969", 0},
	"MKD": {"807", 2}, "MMK": {"104", 0}, "MNT": {"496", 2},
	"MOP": {"446", 2}, "MRU": {"929", 2}, "MUR": {"480", 2},
	"MVR": {"462", 2}, "MWK": {"454", 2}, "MXN": {"484", 2},
	"MXV": {"979", 2}, "MYR": {"458", 2}, "MZN": {"943", 2},
	"NAD": {"516", 2}, "NGN": {"566", 2}, "NIO": {"558", 2},
	"NOK": {"578", 2}, "NPR": {"524", 2}, "NZD": {"554", 2},
	"OMR": {"512", 3}, "PAB": {"590", 2}, "PEN": {"604", 2},
	"PGK": {"598", 2}, "PHP": {"608", 2}, "PKR": {"586", 2},
	"PLN": {"985", 2}, "PYG": {"600", 0}, "QAR": {"634", 2},
	"RON": {"946", 2}, "RSD": {"941", 0}, "RUB": {"643", 2},
	"RWF": {"646", 0}, "SAR": {"682", 2}, "SBD": {"090", 2},
	"SCR": {"690", 2}, "SDG": {"938", 2}, "SEK": {"752", 2},
	"SGD": {"702", 2}, "SHP": {"654", 2}, "SLE": {"925", 2},
	"SLL": {"694", 0}, "SOS": {"706", 0}, "SRD": {"968", 2},
	"SSP": {"728", 2}, "STN": {"930", 2}, "SVC": {"222", 2},
	"SYP": {"760", 0}, "SZL": {"748", 2}, "THB": {"764", 2},
	"TJS": {"972", 2}, "TMT": {"934", 2}, "TND": {"788", 3},
	"TOP": {"776", 2}, "TRY": {"949", 2}, "TTD": {"780", 2},
	"TWD": {"901", 2}, "TZS": {"834", 2}, "UAH": {"980", 2},
	"UGX": {"800", 0}, "USD": {"840", 2}, "USN": {"997", 2},
	"UYI": {"940", 0}, "UYU": {"858", 2}, "UYW": {"927", 4},
	"UZS": {"860", 2}, "VED": {"926", 2}, "VES": {"928", 2},
	"VND": {"704", 0}, "VUV": {"548", 0}, "WST": {"882", 2},
	"XAF": {"950", 0}, "XCD": {"951", 2}, "XOF": {"952", 0},
	"XPF": {"953", 0}, "YER": {"886", 0}, "ZAR": {"710", 2},
	"ZMW": {"967", 2}, "ZWL": {"932", 2},
}
