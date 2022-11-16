package main

import (
	"strings"
	"unicode/utf8"
)

func convertHebNumMonth(hebMonth string) int {
	return convertHebNumber(hebMonth)
}

func convertHebNumYear(hebYear string) int {
	cleanedYear := removeTheThousandLetter(hebYear)
	onlyVal := convertHebNumber(cleanedYear)
	prefix := 5000
	return onlyVal + prefix
}

func removeTheThousandLetter(hebYear string) string {
	if strings.Contains(hebYear, "הת") {
		return strings.Replace(hebYear, "הת", "ת", 1)
	}
	return hebYear
}

func convertHebNumber(hebn string) int {
	finalValue := 0

	// https://siongui.github.io/2016/02/03/go-iterate-over-utf8-non-ascii-string/
	for i, w := 0, 0; i < len(hebn); i += w {
		runeValue, width := utf8.DecodeRuneInString(hebn[i:])
		w = width
		val := valueOfHebLetter(string(runeValue))
		finalValue += val
	}
	return finalValue
}

func valueOfHebLetter(hebl string) int {
	// todo - change to hash map - letters and values

	switch hebl {
	case "א":
		return 1
	case "ב":
		return 2
	case "ג":
		return 3
	case "ד":
		return 4
	case "ה":
		return 5
	case "ו":
		return 6
	case "ז":
		return 7
	case "ח":
		return 8
	case "ט":
		return 9
	case "י":
		return 10
	case "כ":
		return 20
	case "ל":
		return 30
	case "מ":
		return 40
	case "נ":
		return 50
	case "ס":
		return 60
	case "ע":
		return 70
	case "פ":
		return 80
	case "צ":
		return 90
	case "ק":
		return 100
	case "ר":
		return 200
	case "ש":
		return 300
	case "ת":
		return 400
	default:
		return 0
	}
}
