package customers

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type InactiveCustomerConfig struct {
	AllInactiveCustomers []Customer `json:"data"`
}

func AllInactiveCustomersJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}
	println(userID)

	from, _ := strconv.ParseInt(vars["from"], 10, 64)
	fromt := time.Unix(from, 0)
	to, _ := strconv.ParseInt(vars["to"], 10, 64)
	tot := time.Unix(to, 0)

	inactiveCustomers := GetInactiveAllCustomersPerUser(userID, fromt, tot)

	err = json.NewEncoder(w).Encode(InactiveCustomerConfig{AllInactiveCustomers: inactiveCustomers})
	if err != nil {
		println(err.Error())
	}
}

func GetInactiveAllCustomersPerUser(userID int, from, to time.Time) []Customer {
	db := mysqldb.Connect()
	defer db.Close()
	var customers []Customer

	results, err := db.Query(`SELECT customers.id, 
		customers.user_id,
		customers.customer_name,
		customers.customer_surname,
		customers.gender_of_customer,
		customers.date_of_birth,
		customers.customer_address,
		customers.customer_street_number,
		customers.customer_city,
		customers.customer_state,
		customers.customer_zip_code,
		customers.customer_country,
		customers.home_phone_number,
		customers.mobile_phone_number,
		customers.customer_email,
		customers.customer_points,
		customers.is_active,
		customers.pelatis_lianikis,
		customers.created_at,
		customers.updated_at
		FROM customers
		WHERE NOT EXISTS
			(SELECT appointments.created_at
			FROM   appointments
			WHERE customers.user_id = ? AND appointments.customer_id = customers.id AND DATE(appointments.created_at)  
			between date(?) and date(?)) `, userID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var customer Customer
		err = results.Scan(&customer.ID, &customer.UserID, &customer.CustomerName, &customer.CustomerSurname, &customer.CustomerGender, &customer.DateOfBirth, &customer.CustomerAddress, &customer.CustomerStreetNumber, &customer.CustomerCity, &customer.CustomerState, &customer.CustomerZipCode, &customer.CustomerCountry, &customer.HomeTelephone, &customer.MobileTelephone, &customer.CustomerEmail, &customer.CustomerPoints, &customer.IsActive, &customer.PelatisLianikis, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		customers = append(customers, customer)
	}
	return customers
}
