package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func iterateEvents(records [][]string, calId string, yearsAhead int) {
	for _, element := range records {
		b := birthdayFromStringList(element)
		writeEvents(yearsAhead, b, calId)
	}
}

func hebTypeToEventType(typeEvent string) string {
	switch typeEvent {
	case "נ":
		return EventTypes.Wedding
	case "פ":
		return EventTypes.Death
	case "ל":
		return EventTypes.Birth
	default:
		return EventTypes.Birth
	}
}

func birthdayFromStringList(lis []string) birthday {

	for i, element := range lis {
		lis[i] = strings.TrimSpace(element)
	}

	b := birthday{}

	typeEvent := lis[0]
	hebYear := lis[1]
	hebDay := lis[2]
	engMonth := lis[3]
	lastName := lis[4]
	firstName := lis[5]

	b.year = convertHebNumYear(hebYear)
	b.day = convertHebNumMonth(hebDay)
	b.name = fmt.Sprintf("%s %s", firstName, lastName)
	b.month = engMonth
	b.typeEvent = hebTypeToEventType(typeEvent)

	fmt.Println(b)
	return b
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
