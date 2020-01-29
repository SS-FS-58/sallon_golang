package main

import (
	"fmt"
	"mysqldb"
	"time"
	"unicode"

	ics "github.com/PuloV/ics-golang"
)

type Giortes struct {
	Status           int       `json:"status"`
	StatusMessage    string    `json:"status_message"`
	CelebreationDate time.Time `json:"data"`
}

func GetStatheresGiortesFromGoogleCalendar() {

	parser := ics.New()
	// ics.FilePath = "/"
	// ics.DeleteTempFiles = false
	// ics.RepeatRuleApply = true
	// inputChan := parser.GetInputChan()
	// inputChan <- "basic.ics"

	// // wait to kill the main goroute
	// parser.Wait()
	// cal, _ := parser.GetCalendars()
	// for _, e := range cal[0].GetEvents() {
	// 	fmt.Println("More info: ", e.GetStart().Format("2006-01-02"), e.GetSummary())
	// }
	parserChan := parser.GetInputChan()
	parserChan <- "https://calendar.google.com/calendar/ical/kkbrqm0u9ak137ch20ductt23k@group.calendar.google.com/public/basic.ics"
	// t := time.Now()
	outputChan := parser.GetOutputChan()
	//  print events
	var total []*ics.Event
	go func() {
		for event := range outputChan {
			total = append(total, event)

		}
	}()

	// wait to kill the main goroute
	parser.Wait()
	db := mysqldb.Connect()
	defer db.Close()
	for _, event := range total {
		_, err := db.Exec(`INSERT INTO greek_name_days (
			celebreation_date,
			name_day,
			created_at,
			updated_at) 
			VALUES (?,?,NOW(),NOW())  ON DUPLICATE KEY UPDATE celebreation_date = celebreation_date`, event.GetStart(), event.GetSummary())

		if err != nil {
			println(err.Error())

		}

		fmt.Println(event.GetSummary())

	}
}
func GetStatheresGiortesFromDB() Giortes {
	db := mysqldb.Connect()
	defer db.Close()

	var giorty Giortes
	t := time.Now()
	db.QueryRow(`SELECT id,celebreation_date,name_day
	FROM  greek_name_days where celebreation_date = ?`, t.Format("2006-01-02")).Scan(&giorty.Status, &giorty.CelebreationDate, &giorty.StatusMessage)

	// var simera
	// for _, g := range giortes {
	// 	ss := strings.Split(g.StatusMessage, ",")
	// 	simera = append(simera,)

	// }

	return giorty
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
