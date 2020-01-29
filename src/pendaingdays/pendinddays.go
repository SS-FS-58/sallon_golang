package pendaingdays

import (
	"database/sql"
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type PendingDays struct {
	Active       string
	Title        string
	Username     string
	VatNumber    string
	ImageName    int
	UserID       int
	Email        string
	UserProfile  string
	CompanyName  string
	Role         string
	MenuTitle    string
	MessageError string
	UserEmail    string
}

type PendingData struct {
	StoreID              int       `json:"store_id"`
	CustomerName         string    `json:"customer_name"`
	CustomerEmail        string    `json:"customer_email"`
	CustomerMobileNumber string    `json:"customer_mobile_number"`
	HomeTelephone        string    `json:"home_phone_number"`
	CustomerSurname      string    `json:"customer_surname"`
	HairdresserName      string    `json:"hairdresser_name"`
	ServiceName          string    `json:"service_name"`
	ServicePrice         float64   `json:"service_price"`
	StoreName            string    `json:"store_name"`
	Email                string    `json:"email"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type PendingDataConfig struct {
	AllPendingData []PendingData `json:"data"`
}

func AllPendingDataForUser(user_id int) []PendingData {

	db := mysqldb.Connect()
	defer db.Close()

	var events []PendingData

	results, err := db.Query(`SELECT customers.customer_name,
		customers.customer_surname,
		customers.mobile_phone_number,
		customers.home_phone_number,
		customers.customer_email,
		services.service_name,
		services.service_price,
		hairdressers.hairdresser_name,
		shops.company_name,
		users.email,
		shops.id,
		appointments.updated_at

		FROM appointments
		LEFT JOIN hairdressers ON appointments.hairdresser_id = hairdressers.id
		LEFT JOIN customers ON appointments.customer_id = customers.id
		LEFT JOIN services ON appointments.service_id = services.id
		LEFT JOIN shops ON appointments.store_id = shops.id
		LEFT JOIN users ON appointments.user_id = users.id
		WHERE appointments.user_id =? AND service_status = "pending" `, user_id)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var c PendingData

		err = results.Scan(&c.CustomerName, &c.CustomerSurname, &c.CustomerMobileNumber, &c.HomeTelephone, &c.CustomerEmail, &c.ServiceName, &c.ServicePrice, &c.HairdresserName, &c.StoreName, &c.Email, &c.StoreID, &c.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		events = append(events, c)
		// fmt.Printf("%+v", AllAppointmets)

	}

	return events
}

func AllSalesTransCustomerJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])

	allPendingData := AllPendingDataForUser(userID)
	err := json.NewEncoder(w).Encode(PendingDataConfig{AllPendingData: allPendingData})
	if err != nil {
		println(err.Error())
	}
}
