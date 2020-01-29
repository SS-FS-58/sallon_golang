package cancellation

import (
	"bookings"
	"database/sql"
	"encoding/json"
	"mysqldb"
	"net/http"
	"storesettings"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CancellationReports struct {
	Active       string
	Title        string
	Username     string
	VatNumber    string
	ImageName    int
	UserID       int
	UserProfile  string
	UserEmail    string
	CompanyName  string
	Role         string
	MenuTitle    string
	MessageError string
	Stores       []storesettings.Shop
}
type AllCancellationConfig struct {
	AllCancellation []bookings.CalendarEvents `json:"data"`
}
type AllAppointmentPerDay struct {
	SingleDay             time.Time `json:"single_day"`
	TotalComplettedPerDay int       `json:"total_completed"`
	ServiceStatus         string    `json:"service_status"`
}

type CancelledAppointmentPerDay struct {
	SingleDay      time.Time `json:"single_day"`
	TotalCancelled int       `json:"total_cancelled"`
}

type Charts struct {
	Type  string `json:"type"`
	Datas `json:"data"`
}
type Datas struct {
	Labels   []time.Time `json:"labels"`
	Datasets []Datasets  `json:"datasets"`
}

type Datasets struct {
	Label           string `json:"label"`
	BackgroundColor string `json:"backgroundColor"`

	Data []int `json:"data"`
}

type ResponseTime struct {
	UserID       int       `json:"user_id"`
	ContactID    string    `json:"nodeID"`
	ResponseTime float64   `json:"responseTime"`
	NodeName     string    `json:"node_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func AllCancelledAppointment(userID, storeID int, from, to time.Time) []bookings.CalendarEvents {

	db := mysqldb.Connect()
	defer db.Close()

	var events []bookings.CalendarEvents

	results, err := db.Query(`SELECT appointments.id,
		appointments.user_id,
		appointments.customer_id,
		customers.customer_name,
		customers.customer_surname,
		customers.customer_email,
		customers.mobile_phone_number,
		appointments.store_id,
		appointments.hairdresser_id,
		appointments.service_status,
		appointments.start_time,
		appointments.end_time,
		appointments.is_all_day,
		appointments.comments,
		services.service_name,
		appointments.created_at,
		appointments.updated_at
	 FROM appointments
	 LEFT JOIN customers ON appointments.customer_id = customers.id
	 LEFT JOIN services ON appointments.service_id = services.id
	 WHERE appointments.user_id = ? AND appointments.store_id = ? 
	 AND DATE(appointments.created_at) between date(?)
	 AND DATE(?) AND  appointments.service_status="cancelled" `, userID, storeID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var event bookings.CalendarEvents

		err = results.Scan(&event.ID, &event.UserID, &event.CustomerID, &event.CustomerName, &event.CustomerSurname, &event.CustomerEmail, &event.CustomerMobileNumber, &event.StoreID, &event.ResourceID, &event.ServiceStatus, &event.EventStart, &event.EventEnd, &event.IsAllDay, &event.RantevouComments, &event.ServiceName, &event.CreatedAt, &event.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		events = append(events, event)

	}

	return events
}

func AllCancelledAppointmentJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	from, _ := strconv.ParseInt(vars["from"], 10, 64)
	fromt := time.Unix(from, 0)
	to, _ := strconv.ParseInt(vars["to"], 10, 64)
	tot := time.Unix(to, 0)

	totalCancellation := AllCancelledAppointment(userID, storeID, fromt, tot)

	err = json.NewEncoder(w).Encode(AllCancellationConfig{AllCancellation: totalCancellation})
	if err != nil {
		println(err.Error())
	}
}
func GetAllAppointmentsPesServiceStatusSONGraph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	from, _ := strconv.ParseInt(vars["from"], 10, 64)
	fromt := time.Unix(from, 0)
	to, _ := strconv.ParseInt(vars["to"], 10, 64)
	tot := time.Unix(to, 0)

	totalAppoitments := TotalAppointmentPerDay(userID, storeID, fromt, tot)
	tableCharts := Datas{}.Datasets
	labels := Datas{}.Labels

	var serviceStatus []string
	for _, s := range totalAppoitments {
		labels = append(labels, s.SingleDay)

		serviceStatus = append(serviceStatus, s.ServiceStatus)

	}

	serviceStatu := removeDuplicatesUnordered(serviceStatus)
	for i, v := range serviceStatu {
		if v == "pending" {
			serviceStatu = append(serviceStatu[:i], serviceStatu[i+1:]...)
			break
		}
	}

	colors := []string{"#27b99a", "#b92746"}

	for i, n := range serviceStatu {
		totals := []int{}

		for _, t := range totalAppoitments {
			if n == t.ServiceStatus {

				totals = append(totals, t.TotalComplettedPerDay)

			}
		}

		t := Datasets{
			Label:           n,
			BackgroundColor: colors[i],
			Data:            totals,
		}

		tableCharts = append(tableCharts, t)

	}

	c := Charts{
		Type: "line",
		Datas: Datas{
			Labels:   labels,
			Datasets: tableCharts,
		},
	}

	// fmt.Println(statistics)
	// res2B, _ := json.Marshal(statistics)
	// fmt.Println("\n", string(res2B))
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		println(err.Error())
	}
}

func GetAllAppointmentsPerStatusJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	from, _ := strconv.ParseInt(vars["from"], 10, 64)
	fromt := time.Unix(from, 0)
	to, _ := strconv.ParseInt(vars["to"], 10, 64)
	tot := time.Unix(to, 0)

	totalAppoitments := TotalAppointmentPerDay(userID, storeID, fromt, tot)

	err = json.NewEncoder(w).Encode(totalAppoitments)
	if err != nil {
		println(err.Error())
	}
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}
func TotalAppointmentPerDay(userID, storeID int, from, to time.Time) []AllAppointmentPerDay {

	db := mysqldb.Connect()
	defer db.Close()

	var totals []AllAppointmentPerDay

	results, err := db.Query(`select 
		DISTINCT DATE(w.wdate) AS sorted_date,
		w.wval + ifnull(x.count_status,0) AS newval,
	  
	   w.service_status
	from
	   (
		   select 
				   d.wdate,
				   d.wval,
				  
				   p.service_status
		   from daily_values d
		   cross join appointments p 
	   ) w
	   left join
	   (
		   SELECT  s.service_status,
				  
				  DATE(s.created_at) as doc_date,
				   COUNT(service_status) AS count_status
		   FROM    appointments s
				  
				   WHERE s.user_id = ? AND s.store_id = ? AND DATE(s.created_at)  
			  between date(?) and date(?)
		   GROUP   BY s.service_status,DATE(s.created_at)
	   ) x on w.wdate = x.doc_date and w.service_status = x.service_status
	where w.wdate between date(?) and date(?)
	ORDER BY sorted_date`, userID, storeID, from, to, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total AllAppointmentPerDay

		err = results.Scan(&total.SingleDay, &total.TotalComplettedPerDay, &total.ServiceStatus)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		totals = append(totals, total)

	}

	return totals
}

func AllCancelledAppointmentPerDay(userID, storeID int, from, to time.Time) []CancelledAppointmentPerDay {

	db := mysqldb.Connect()
	defer db.Close()

	var totalCancelled []CancelledAppointmentPerDay

	results, err := db.Query(`SELECT  DATE(created_at) AS created_as,
							COUNT(service_status) AS count_status
							FROM appointments
							WHERE service_status = "cancelled" AND user_id =? and store_id = ? and DATE(created_at)  
							between date(?) and date(?)

							GROUP BY DATE(created_at),service_status
							ORDER BY DATE(created_at)`, userID, storeID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var cancelled CancelledAppointmentPerDay

		err = results.Scan(&cancelled.SingleDay, &cancelled.TotalCancelled)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		totalCancelled = append(totalCancelled, cancelled)

	}

	return totalCancelled
}
