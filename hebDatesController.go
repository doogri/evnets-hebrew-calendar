package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// example output:
// {
// 	"gy":1988,
// 	"gm":12,
// 	"gd":4,
// 	"afterSunset":false,
// 	"hy":5749,
// 	"hm":"Kislev",
// 	"hd":25,
// 	"hebrew":"כ״ה בְּכִסְלֵו תשמ״ט",
// 	"heDateParts":{"y":"תשמ״ט","m":"כסלו","d":"כ״ה"},
// 	"events":["Chanukah day 1","Parashat Miketz"]
// }

// https://www.hebcal.com/home/219/hebrew-date-converter-rest-api
// https://www.hebcal.com/converter?cfg=json&hy=5749&hm=Kislev&hd=25&h2g=1&strict=1
func convertDateFromHebrew(year int, month string, day int) time.Time {
	url := fmt.Sprintf("https://www.hebcal.com/converter?cfg=json&hy=%d&hm=%s&hd=%d&h2g=1&strict=1", year, month, day)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	m := make(map[string]interface{})
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		log.Fatal(err)
	}

	yearG := int(m["gy"].(float64))
	monthG := time.Month(m["gm"].(float64))
	dayG := int(m["gd"].(float64))

	t := time.Date(yearG, monthG, dayG, 0, 0, 0, 0, time.UTC)
	fmt.Println("the greg date: " + t.String())

	return t
}
