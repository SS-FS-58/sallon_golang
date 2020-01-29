package storesettings

import (
	"database/sql"
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type WeeklyHours struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	StoreID      int       `json:"store_id"`
	DayOfTheWeek int       `json:"day_of_the_week"`
	NameWeekDay  string    `json:"name_week_day"`
	FromTime     string    `json:"start"`
	ToTime       string    `json:"end"`
	StoreName    string    `json:"store_name"`
	Color        string    `json:"color"`
	Dow          []int     `json:"dow"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `josn:"updated_at"`
	Rendering    string    `json:"rendering"`
}

type AllWeeklyHoursConfig struct {
	AllWeekly []WeeklyHours `json:"data"`
}

func InsertWeeklyHoursDataToDB(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	weeklyHoursJSON := r.PostFormValue("weeklyHoursJSON")
	var weeklyHours WeeklyHours

	err := json.Unmarshal([]byte(weeklyHoursJSON), &weeklyHours)

	if err != nil {
		println(err.Error())

	}
	if rowExists("SELECT id FROM weekly_hours WHERE  user_id = ? and store_id = ? and day_of_the_week = ? and from_time = ? ", weeklyHours.UserID, weeklyHours.StoreID, weeklyHours.DayOfTheWeek, weeklyHours.FromTime) {
		witchDay := weeklyHours.DayOfTheWeek
		p := map[int]string{
			1: "Monday",
			2: "Tuesday",
			3: "Wednesday",
			4: "Thursday",
			5: "Friday",
			6: "Saturday",
			0: "Sunday",
		}

		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "The: " + p[witchDay] + " with start time: " + weeklyHours.FromTime + " already exists"})
		return
	}

	stmt, _ := db.Prepare(`INSERT INTO weekly_hours (user_id,
		store_id,
		day_of_the_week,
		from_time,
		to_time,
		created_at,
		updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(weeklyHours.UserID, weeklyHours.StoreID, weeklyHours.DayOfTheWeek, weeklyHours.FromTime, weeklyHours.ToTime)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert weekly hours to the database: " + err.Error()})

	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed get last insert id from database: " + err.Error()})

	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted weekly hours into the Database", Body: strconv.Itoa(int(id))})

}
func AllWeeklyHoursPerUserPerStore(userID int) []WeeklyHours {

	db := mysqldb.Connect()
	defer db.Close()

	var weeklyHours []WeeklyHours

	results, err := db.Query(`SELECT weekly_hours.id,
		weekly_hours.user_id,
		weekly_hours.store_id,
		weekly_hours.day_of_the_week,
		weekly_hours.from_time,
		weekly_hours.to_time,
		shops.company_name,
		weekly_hours.created_at,
		weekly_hours.updated_at
	   FROM weekly_hours 
		LEFT JOIN shops ON weekly_hours.store_id = shops.id
	   WHERE weekly_hours.user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var weeklyHour WeeklyHours

		err = results.Scan(&weeklyHour.ID, &weeklyHour.UserID, &weeklyHour.StoreID, &weeklyHour.DayOfTheWeek, &weeklyHour.FromTime, &weeklyHour.ToTime, &weeklyHour.StoreName, &weeklyHour.CreatedAt, &weeklyHour.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}
		witchDay := weeklyHour.DayOfTheWeek
		p := map[int]string{
			1: "Monday",
			2: "Tuesday",
			3: "Wednesday",
			4: "Thursday",
			5: "Friday",
			6: "Saturday",
			0: "Sunday",
		}
		weeklyHour.NameWeekDay = p[witchDay]
		// fmt.Println("Week day : ", time.Weekday(weeklyHour.DayOfTheWeek).String())
		weeklyHour.Dow = append(weeklyHour.Dow, weeklyHour.DayOfTheWeek)
		// if val, ok := p[witchDay]; ok {
		// 	println(val)
		// } else {
		// 	println(p[witchDay])
		// }
		weeklyHour.Rendering = "inverse-background"
		weeklyHours = append(weeklyHours, weeklyHour)
		// fmt.Printf("%+v", AllAppointmets)

	}

	return weeklyHours
}
func AllWeeklyHoursPerUserJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])

	allWeeklyHoursPerUserPerStore := AllWeeklyHoursPerUserPerStore(userID)

	err := json.NewEncoder(w).Encode(AllWeeklyHoursConfig{AllWeekly: allWeeklyHoursPerUserPerStore})
	if err != nil {
		println(err.Error())
	}
}

func AllWeeklyHoursPerUserPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])
	storeID, _ := strconv.Atoi(vars["storeid"])
	var weekdays []WeeklyHours

	allWeeklyHoursPerUserPerStore := AllWeeklyHoursPerUserPerStore(userID)
	for _, w := range allWeeklyHoursPerUserPerStore {
		if w.StoreID == storeID {
			weekdays = append(weekdays, w)
		}
	}
	err := json.NewEncoder(w).Encode(weekdays)
	if err != nil {
		println(err.Error())
	}
}
func DeleteWeekDay(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	dayID := vars["id"]

	ID, _ := strconv.Atoi(dayID)

	stmt, _ := db.Prepare("DELETE FROM weekly_hours where id = ?")

	_, err := stmt.Exec(ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete day from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted day from database "})
}
