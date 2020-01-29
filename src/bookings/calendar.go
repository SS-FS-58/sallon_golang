package bookings

import (
	"customers"
	"encoding/json"
	"fmt"
	tr "gettext"
	"mysqldb"
	"net/http"
	"os"
	"storesettings"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Calendar struct {
	Active       string
	Title        string
	Username     string
	ImageName    int
	UserID       int
	VatNumber    string
	UserProfile  string
	CompanyName  string
	Role         string
	MenuTitle    string
	MessageError string
	Customers    []customers.Customer
	Stores       []storesettings.Shop
}

type CreateRantevou struct {
	UserID                int    `json:"user_id"`
	RantevouStart         string `json:"rantevou_start"`
	RantevouEnd           string `json:"rantevou_end"`
	RantevouStoreID       int    `json:"rantevou_store_id"`
	RantevouHairdresserID int    `json:"rantevou_hairdresser_id"`
	RantevouCustomerID    int    `json:"select_customer_id"`
	AllRantevouServices   string `json:"all_services"`
	AllRantevouPrices     string `json:"rantevou_price"`
	AllRantevouStatuses   string `json:"all_statuses"`
	AllRantevouDurations  string `json:"rantevou_duration"`
	RantevouComments      string `json:"rantevou_comments"`
	HairdresserName       string `json:"hairdresser_name"`
	IsAllDay              bool   `json:"is_all_day"`
}

type CalendarEvents struct {
	ID                   int       `json:"id"`
	UserID               int       `json:"user_id"`
	CustomerID           int       `json:"customer_id"`
	CustomerName         string    `json:"customer_name"`
	CustomerEmail        string    `json:"customer_email"`
	CustomerMobileNumber string    `json:"customer_mobile_number"`
	CustomerSurname      string    `json:"customer_surname"`
	StoreID              int       `json:"store_id"`
	ResourceID           int       `json:"resourceId"`
	ServiceID            int       `json:"service_id"`
	ServiceStatus        string    `json:"service_status"`
	EventStart           time.Time `json:"event_start"`
	EventEnd             time.Time `json:"event_end"`
	StartEvent           string    `json:"start"`
	EndEvent             string    `json:"end"`
	EventTitle           string    `json:"title"`
	IsAllDay             bool      `json:"is_all_day"`
	RantevouComments     string    `json:"rantevou_comments"`
	ServiceName          string    `json:"service_name"`
	HairdresserName      string    `json:"hairdresser_name"`
	SwitchFormula        string    `json:"switch_formula"`
	HairdresserID        int       `json:"hairdresser_id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type CalendarDataResource struct {
	RowID      int    `json:"id"`
	Title      string `json:"title"`
	EventColor string `json:"eventColor"`
}
type HTTPResp struct {
	Status        int    `json:"status"`
	Description   string `json:"description"`
	Body          string `json:"body"`
	InvoiceNumber int    `json:"invoice_number"`
}
type UpdateCalendar struct {
	ID         int    `json:"id"`
	EventStart string `json:"start"`
	EventEnd   string `json:"end"`
	ResourceID int    `json:"resource_id"`
}

type CancelAppointment struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func InsertRantevouJSON(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var rantevou CreateRantevou

	err := decoder.Decode(&rantevou)
	if err != nil {
		print(err.Error())
	}
	allServiceStatus := strings.Split(rantevou.AllRantevouStatuses, ":")
	allServiceIDs := strings.Split(rantevou.AllRantevouServices, ":")
	allservicePrices := strings.Split(rantevou.AllRantevouPrices, ":")
	allserviceDurations := strings.Split(rantevou.AllRantevouDurations, ":")
	allHairdressers := strings.Split(rantevou.HairdresserName, ":")
	allComments := strings.Split(rantevou.RantevouComments, "::")

	var hairdresserIDs = []int{}
	var serviceIDs = []int{}
	var servicePrices = []float64{}
	var serviceDurations = []int{}

	for _, i := range allHairdressers {
		j, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
		}
		hairdresserIDs = append(hairdresserIDs, j)
	}

	for _, i := range allServiceIDs {
		j, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
		}
		serviceIDs = append(serviceIDs, j)
	}

	for _, i := range allserviceDurations {
		j, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
		}
		serviceDurations = append(serviceDurations, j)
	}
	for _, i := range allservicePrices {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			println(err.Error())
		}
		servicePrices = append(servicePrices, j)
	}

	strt := rantevou.RantevouStart

	// twoDays := time.Hour * 24 * 2
	// start := eventStart.Add(twoDays).Format("2006-01-02T15:04:05")
	// 		end := eventEnd.Add(twoDays).Format("2006-01-02T15:04:05")
	startTime, _ := time.Parse("2006-01-02T15:04:05", strt)
	for i, status := range allServiceStatus {
		t := time.Duration(serviceDurations[i]) * time.Minute

		endTime := startTime.Add(t)
		start := startTime.Format("2006-01-02T15:04:05")
		end := endTime.Format("2006-01-02T15:04:05")
		stmt, err := db.Prepare(`INSERT INTO appointments (user_id,
		hairdresser_id,
		customer_id,
		store_id,
		service_id,
		start_time,
		end_time,
		service_status,
		comments,
		is_all_day,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
		if err != nil {
			println(err.Error())
		}
		res, err := stmt.Exec(rantevou.UserID, hairdresserIDs[i], rantevou.RantevouCustomerID, rantevou.RantevouStoreID, allServiceIDs[i], start, end, status, allComments[i], rantevou.IsAllDay)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert appointments into database"})
		}
		intID, err := res.LastInsertId()
		if err != nil {
			println(err.Error(), intID)
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
		}
		startTime = endTime
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted appointment Into the Database"})
}

func AllCalendarResources(userID, storeID int) []CalendarDataResource {

	db := mysqldb.Connect()
	defer db.Close()

	var resources []CalendarDataResource

	results, err := db.Query(`SELECT hairdresser_id,hairdresser_name,color FROM hairdresser_view WHERE user_id = ? AND company_id = ? AND is_active_hairdresser = 1`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var resource CalendarDataResource
		err = results.Scan(&resource.RowID, &resource.Title, &resource.EventColor)
		if err != nil {
			println(err.Error())
		}
		resources = append(resources, resource)
	}
	return resources
}
func AllCalendarEvents(userID, storeID int, eventtime string, endtime string) []CalendarEvents {

	db := mysqldb.Connect()
	defer db.Close()

	var events []CalendarEvents

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
		hairdressers.hairdresser_name
		FROM customers
		INNER JOIN appointments ON customers.id = appointments.customer_id
		JOIN services ON services.id = appointments.service_id
		JOIN hairdressers ON hairdressers.id = appointments.hairdresser_id
		WHERE  appointments.store_id = ? AND appointments.user_id = ? AND DATE(appointments.start_time) >= ? AND DATE(appointments.end_time) <= ? `, storeID, userID, eventtime, endtime)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var event CalendarEvents

		err = results.Scan(&event.ID, &event.UserID, &event.CustomerID, &event.CustomerName, &event.CustomerSurname, &event.CustomerEmail, &event.CustomerMobileNumber, &event.StoreID, &event.ResourceID, &event.ServiceStatus, &event.EventStart, &event.EventEnd, &event.IsAllDay, &event.RantevouComments, &event.ServiceName, &event.HairdresserName)
		if err != nil {
			println(err.Error())
		}

		event.EventTitle = event.CustomerName + " " + event.CustomerSurname

		events = append(events, event)

	}

	return events
}

func AllCommentsPerCustomer(customerID int) []CalendarEvents {
	db := mysqldb.Connect()
	defer db.Close()
	var allComments []CalendarEvents

	results, err := db.Query(`SELECT appointments.id,
		appointments.user_id,
		appointments.hairdresser_id,
		appointments.customer_id,
		appointments.store_id,
		appointments.service_id,
		services.service_name,
		appointments.start_time, 
		appointments.end_time,
		appointments.service_status,
		appointments.comments,
		appointments.created_at,
		appointments.updated_at FROM appointments 
		LEFT JOIN services ON appointments.service_id = services.id
		where appointments.customer_id = ? 
		ORDER BY appointments.created_at DESC`, customerID)
	if err != nil {
		println(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var singleComment CalendarEvents

		err = results.Scan(&singleComment.ID, &singleComment.UserID, &singleComment.HairdresserID, &singleComment.CustomerID, &singleComment.StoreID, &singleComment.ServiceID, &singleComment.ServiceName, &singleComment.EventStart, &singleComment.EventEnd, &singleComment.ServiceStatus, &singleComment.RantevouComments, &singleComment.CreatedAt, &singleComment.UpdatedAt)
		if err != nil {
			println(err.Error())
		}

		allComments = append(allComments, singleComment)

	}

	return allComments

}
func AllAppointmentPerCustomerJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, _ := strconv.Atoi(vars["customerid"])

	allAppointmentCustomers := AllCommentsPerCustomer(customerID)
	err := json.NewEncoder(w).Encode(allAppointmentCustomers)

	if err != nil {
		println(err.Error())
	}

}

func AllCalendarEventsJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userid"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid userid")
	}
	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid storeid")
	}
	eventtime := vars["eventtime"]
	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid eventtime")
	}
	endtime := vars["endtime"]
	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid endtime")
	}
	events := AllCalendarEvents(userID, storeID, eventtime, endtime)

	var eventss []CalendarEvents
	for _, e := range events {
		var start string
		var end string
		if e.IsAllDay == true {
			start = e.EventStart.Format("2006-01-02")
			end = e.EventEnd.Format("2006-01-02")
		} else {
			start = e.EventStart.Format("2006-01-02T15:04:05")
			end = e.EventEnd.Format("2006-01-02T15:04:05")
		}
		e.StartEvent = start
		e.EndEvent = end
		eventss = append(eventss, e)

	}

	err = json.NewEncoder(w).Encode(eventss)
	if err != nil {
		println(err.Error())
	}

}

func AllCalendarResourcesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userid"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid userid")
	}
	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid userid")
	}
	resources := AllCalendarResources(userID, storeID)
	err = json.NewEncoder(w).Encode(resources)
	if err != nil {
		println(err.Error())
	}
}
func UpdateCalendarDataJSONToDB(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var calendar UpdateCalendar
	err := decoder.Decode(&calendar)
	if err != nil {
		println(err.Error())
	}

	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE appointments SET start_time = ?,end_time = ?,updated_at = NOW() where id = ? ")
	if err != nil {
		println(err.Error())
	}
	_, err = stmt.Exec(calendar.EventStart, calendar.EventEnd, calendar.ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update table in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Failed to update table in the database"})
}

func UpdateCalendarDataJSONToDBWithHairdresserID(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var calendar UpdateCalendar
	err := decoder.Decode(&calendar)
	if err != nil {
		println(err.Error())
	}

	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE appointments SET hairdresser_id = ?, start_time = ?,end_time = ?,updated_at = NOW() where id = ? ")
	if err != nil {
		println(err.Error())
	}
	_, err = stmt.Exec(calendar.ResourceID, calendar.EventStart, calendar.EventEnd, calendar.ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update table in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully to update table in the database"})
}
func CancelAppointmentJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var calendar CancelAppointment
	err := decoder.Decode(&calendar)
	if err != nil {
		println(err.Error())
	}

	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE appointments SET service_status = ?,updated_at = NOW() where id = ? ")
	if err != nil {
		println(err.Error())
	}
	_, err = stmt.Exec(calendar.Status, calendar.ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update table in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully to update table in the database"})
}
