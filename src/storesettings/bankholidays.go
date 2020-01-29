package storesettings

import (
	"encoding/json"
	"fmt"
	tr "gettext"
	"log"
	"mysqldb"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MultipBankHolidays struct {
	TotalHolidays []GreekHolidays `json:"items"`
}

type GreekHolidays struct {
	King         string    `json:"kind"`
	Etag         string    `json:"etag"`
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	HTMLLink     string    `json:"htmlLink"`
	CreatedAt    string    `json:"created"`
	UpdatedAt    string    `json:"updated"`
	Summary      string    `json:"summary"`
	Creators     Creator   `json:"creator"`
	Organizers   Organizer `json:"organizer"`
	Starts       Start     `json:"start"`
	Ends         End       `json:"end"`
	Transaprency string    `json:"transparency"`
	Visibility   string    `json:"visibility"`
	ICalUID      string    `json:"iCalUID"`
	Sequence     int       `json:"sequence"`
}

type Creator struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Self        bool   `json:"self"`
}

type Organizer struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Self        bool   `json:"self"`
}

type Start struct {
	Date string `json:"date"`
}

type End struct {
	Date string `json:"date"`
}

type BankHolidaysPerStore struct {
	ID               int       `json:"id"`
	Summary          string    `json:"title"`
	BankHolidayStart time.Time `json:"start"`
	BankHolidayEnd   time.Time `json:"end"`
	Country          string    `json:"country"`
	StoreID          int       `json:"store_id"`
	UserID           int       `json:"user_id"`
	StoreName        string    `json:"store_name"`
}
type HolidaysConfig struct {
	BankHolidays []BankHolidaysPerStore `json:"data"`
}

func BankHolidays() {
	var primes = []string{"australian", "austrian", "brazilian", "canadian", "china", "christian", "danish", "dutch", "finnish", "french", "german", "greek", "hong_kong_c", "hong_kong", "indian", "indonesian", "iranian", "irish", "islamic", "italian", "japanese", "jewish", "malaysia", "mexican", "new_zealand", "norwegian", "philippines", "polish", "portuguese", "russian", "singapore", "sa", "south_korea", "spain", "swedish", "taiwan", "thai", "uk", "usa"}
	db := mysqldb.Connect()
	defer db.Close()
	start := time.Now()
	numComplete := 0

	for _, symbol := range primes {

		go func(symbol string) {
			var url = "https://www.googleapis.com/calendar/v3/calendars/en." + symbol + "%23holiday%40group.v.calendar.google.com/events?key=AIzaSyDDys-k9RKuTBWIKLORovbuv-2_8FbKOJg"
			bankHolidays := &MultipBankHolidays{}
			err := getJSON(url, &bankHolidays)

			if err != nil {
				log.Println(err.Error())
			}
			for _, c := range bankHolidays.TotalHolidays {
				_, err := db.Exec(`INSERT INTO bank_holidays (
					summary,
					date_start,
					date_end,
					country,
					created_at,
					updated_at) 
					VALUES (?,?,?,?,NOW(),NOW())`, c.Summary, c.Starts.Date, c.Ends.Date, symbol)

				if err != nil {
					println(err.Error())

				}
			}

			numComplete++
		}(symbol)
	}

	for numComplete < len(primes) {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)

	fmt.Printf("Execution Time: %s", elapsed)

}

var myClient = &http.Client{Timeout: 30 * time.Second}

//getJson function to return the json data from a given url
func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetAllBankHolidaysPerStore(storeID int) []BankHolidaysPerStore {
	db := mysqldb.Connect()
	defer db.Close()

	var holidays []BankHolidaysPerStore

	results, err := db.Query(`SELECT bank_holidays.id, bank_holidays.summary,
		bank_holidays.date_start,
		bank_holidays.date_end,
		bank_holidays.country,
		shops.company_name,
		bank_holidays_per_store.store_id,
		bank_holidays_per_store.user_id
	FROM  bank_holidays_per_store
	LEFT JOIN bank_holidays ON bank_holidays.country  = bank_holidays_per_store.bank_holidays_country
	LEFT JOIN shops ON shops.id = bank_holidays_per_store.store_id
	WHERE bank_holidays_per_store.store_id = ?`, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var holiday BankHolidaysPerStore
		err = results.Scan(&holiday.ID, &holiday.Summary, &holiday.BankHolidayStart, &holiday.BankHolidayEnd, &holiday.Country, &holiday.StoreName, &holiday.StoreID, &holiday.UserID)
		if err != nil {
			println(err.Error())
		}
		holidays = append(holidays, holiday)
	}
	return holidays
}

func AllHolidaysJSONPerStore(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	vars := mux.Vars(r)
	storeID, err := strconv.Atoi(vars["storeid"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}

	var bankhol []BankHolidaysPerStore
	holidays := GetAllBankHolidaysPerStore(storeID)
	for _, b := range holidays {
		b.Summary = tr.Translate(b.Summary)
		bankhol = append(bankhol, b)
	}
	err = json.NewEncoder(w).Encode(HolidaysConfig{BankHolidays: bankhol})
	if err != nil {
		println(err.Error())
	}
}
