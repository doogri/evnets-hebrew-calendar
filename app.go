// todo - features:
// add age (by birth year)
// set reminders a week before (make configurable by user input)
// gui
// additional kinds of events (wedding, death, free text)
// input list from file
// write to specific secondary calendar - get by args

package main

import (
	"flag"
	"time"
)

type birthday struct {
	name  string
	year  int
	month string
	day   int
}

func main() {
	//config
	flag.Parse()
	initViper()
	yearsAhead := viperEnvVariableInt("YEARS_AHEAD")
	calId := viperEnvVariableStr("CALENDAR_ID")

	hebBirthday := getDetailsFromUser()
	currHebYear := calcCurrHebYear()

	for hebYear := currHebYear; hebYear < currHebYear+yearsAhead; hebYear++ {
		gregDate := convertDateFromHebrew(hebYear, hebBirthday.month, hebBirthday.day)
		setEventGoogleCal(hebBirthday, gregDate, calId)
	}
}

func calcCurrHebYear() int {
	currYear := time.Now().Year()
	GAP_CALS := 5782 - 2022
	currHebYear := currYear + GAP_CALS
	return currHebYear
}

var (
	birthYear = flag.Int("birthYear", -1, "hebrew birth year")
	month     = flag.String("month", "", "Nisan, Iyyar, Sivan, Tamuz, Av, Elul, Tishrei, Cheshvan, Kislev, Tevet, Shvat, Adar, Adar1, Adar2")
	day       = flag.Int("day", -1, "day")
	name      = flag.String("name", "", "the name")
)

func getDetailsFromUser() birthday {
	// listMonths := [14]string{
	// 	"Nisan",
	// 	"Iyyar",
	// 	"Sivan",
	// 	"Tamuz",
	// 	"Av",
	// 	"Elul",
	// 	"Tishrei",
	// 	"Cheshvan",
	// 	"Kislev",
	// 	"Tevet",
	// 	"Shvat",
	// 	"Adar",
	// 	"Adar1",
	// 	"Adar2"}

	b := birthday{}
	b.day = *day
	b.month = *month
	b.year = *birthYear
	b.name = *name

	return b
}

func setEventGoogleCal(b birthday, t time.Time, calId string) {
	setEvent(b.name, t, calId)
}
