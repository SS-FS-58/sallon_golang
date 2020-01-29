package bookings

import (
	"database/sql"
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AllEventConfig struct {
	AllEvents []CalendarEvents `json:"data"`
}
type SalesTransConfig struct {
	AllSales []CheckOutData `json:"data"`
}

func AllSalesTransPerCustomer(customerid int) []CheckOutData {

	db := mysqldb.Connect()
	defer db.Close()

	var events []CheckOutData

	results, err := db.Query(`SELECT 
		salestrans.id AS  'salestrans_id',
		salestrans.user_id,
		salestrans_common.hairdresser_id,
		salestrans.store_id,
		salestrans.customer_id,
		salestrans.service_id,
		salestrans.service_price,
		salestrans.service_qty,
		salestrans.service_discount,
		salestrans.payment_type,
		salestrans.service_line_total,
		salestrans.is_service,
		salestrans.invoice_number,
		salestrans.created_at,
		salestrans.updated_at,
		hairdressers.hairdresser_name,
		CONCAT_WS('', services.service_name, products.product_name ) AS  service_name,
		appointments.comments,
		appointments.id 
		FROM salestrans_common
        LEFT JOIN salestrans ON salestrans_common.salestrans_id = salestrans.id
		LEFT JOIN services ON salestrans.service_id= services.id AND salestrans.is_service = true
		LEFT JOIN products ON salestrans.service_id = products.id AND salestrans.is_service = false
		LEFT JOIN hairdressers ON salestrans_common.hairdresser_id = hairdressers.id
		LEFT JOIN appointments ON  salestrans_common.appointment_id = appointments.id
	 WHERE appointments.customer_id = ? `, customerid)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var c CheckOutData

		err = results.Scan(&c.SalesTransID, &c.UserID, &c.HairdresserID, &c.StoreID, &c.CustomerID, &c.ServiceID, &c.ServicePrice, &c.ServiceQty, &c.ServiceDiscount, &c.PaymentType, &c.ServiceLineTotal, &c.IsService, &c.InvoiceNumber, &c.CreatedAt, &c.UpdatedAt, &c.HairdresserName, &c.ServiceName, &c.ServiceComments, &c.ServiceID)
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
	customerID, _ := strconv.Atoi(vars["customerid"])

	allsalesPerCustomer := AllSalesTransPerCustomer(customerID)
	err := json.NewEncoder(w).Encode(SalesTransConfig{AllSales: allsalesPerCustomer})
	if err != nil {
		println(err.Error())
	}
}

func AllCalendarEventsPerCustomer(customerID int) []CalendarEvents {

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
		hairdressers.hairdresser_name,
		appointments.created_at,
		appointments.updated_at
		FROM appointments
		LEFT JOIN customers ON appointments.customer_id = customers.id
		LEFT JOIN services ON appointments.service_id = services.id
		LEFT JOIN hairdressers ON appointments.hairdresser_id = hairdressers.id
		WHERE  appointments.customer_id= ? AND NOT appointments.service_status="complete"  `, customerID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var event CalendarEvents

		err = results.Scan(&event.ID, &event.UserID, &event.CustomerID, &event.CustomerName, &event.CustomerSurname, &event.CustomerEmail, &event.CustomerMobileNumber, &event.StoreID, &event.ResourceID, &event.ServiceStatus, &event.EventStart, &event.EventEnd, &event.IsAllDay, &event.RantevouComments, &event.ServiceName, &event.HairdresserName, &event.CreatedAt, &event.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		events = append(events, event)

	}

	return events
}

func AllCalendarEventsPerCustomerJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, _ := strconv.Atoi(vars["customerid"])

	alleventsPerCustomer := AllCalendarEventsPerCustomer(customerID)

	err := json.NewEncoder(w).Encode(AllEventConfig{AllEvents: alleventsPerCustomer})
	if err != nil {
		println(err.Error())
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var event CalendarEvents
	err := decoder.Decode(&event)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	eventID := vars["id"]
	ID, err := strconv.Atoi(eventID)

	if err != nil {
		println(err.Error())
	}
	stmt, _ := db.Prepare("UPDATE appointments SET comments = ? WHERE id = ?")
	_, err = stmt.Exec(event.RantevouComments, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update appointmnet in the database: " + err.Error()})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update appointmnet in the Database"})
}
