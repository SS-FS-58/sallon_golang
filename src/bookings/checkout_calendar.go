package bookings

import (
	"customers"
	"encoding/json"
	"math/rand"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CheckOutCalendar struct {
	Active              string
	Title               string
	Username            string
	VatNumber           string
	ImageName           int
	UserID              int
	UserProfile         string
	CompanyName         string
	Role                string
	MenuTitle           string
	MessageError        string
	CustomerName        string
	CustomerPhone       string
	CustomerEmail       string
	HairdresserName     string
	AppointmentCloseDay time.Time
	CustomerCheckout    []CustomerCheckout
	Total               float64 `json:"total"`
	TotalDiscount       float64 `json:"total_discount"`
	SingleCustomer      customers.Customer
	StoreID             int
	HairdresserID       int
	CustomerID          int
	EventID             string
}

type CustomerCheckout struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	CustomerID       int       `json:"customer_id"`
	CustomerName     string    `json:"customer_name"`
	CustomerEmail    string    `json:"customer_email"`
	CustomerSurname  string    `json:"customer_surname"`
	StoreID          int       `json:"store_id"`
	ResourceID       int       `json:"resourceId"`
	ServiceID        int       `json:"service_id"`
	ServiceStatus    string    `json:"service_status"`
	EventStart       time.Time `json:"event_start"`
	EventEnd         time.Time `json:"event_end"`
	StartEvent       string    `json:"start"`
	EndEvent         string    `json:"end"`
	EventTitle       string    `json:"title"`
	IsAllDay         bool      `json:"is_all_day"`
	RantevouComments string    `json:"rantevou_comments"`
	ServiceName      string    `json:"service_name"`
	ServicePrice     float64   `json:"service_price"`
	ServiceDiscount  float64   `json:"service_discount"`
	HairdresserName  string    `json:"hairdresser_name"`
	TotalPerRow      float64   `json:"total_per_row"`
	SwitchFormula    bool      `json:"switch_formula"`
}
type CheckOutData struct {
	SalesTransID       int    `json:"sales_trans_id"`
	UserID             int    `json:"user_id"`
	HairdresserID      int    `json:"hairdresser_id"`
	HairdresserName    string `json:"hairdresser_name"`
	StoreID            int    `json:"store_id"`
	CustomerID         int    `json:"customer_id"`
	PromotionProductID int    `json:"promotion_product_id"`

	ServiceName      string    `json:"service_name"`
	ServiceID        int       `json:"service_id"`
	ServicePrice     float64   `json:"service_price"`
	ServiceQty       int       `json:"service_qty"`
	ServiceDiscount  float64   `json:"service_discount"`
	PaymentType      string    `json:"payment_type"`
	ServiceLineTotal float64   `json:"service_line_total"`
	IsService        bool      `json:"is_service"`
	ServiceStarTime  time.Time `json:"service_start_time"`
	ServiceEndTime   time.Time `json:"service_end_time"`
	IsNew            bool      `json:"is_new"`
	ServiceComments  string    `json:"service_comments"`
	EventID          int       `json:"event_id"`
	InvoiceNumber    int       `json:"invoice_number"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Invoices struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	StoreID      int       `json:"store_id"`
	CustomerID   int       `json:"customer_id"`
	SubTotalCost float64   `json:"sub_total_cost"`
	Discount     float64   `json:"discount"`
	Total        float64   `json:"total"`
	StatusID     int       `json:"status_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type SpendProducts struct {
	ProductName string `json:"product_name"`
}

type CheckoutConfig struct {
	AllCheckout []CustomerCheckout `json:"data"`
}

var AllAppointmets []CustomerCheckout

func AllAppointmentPerCustomerPerDay(userID, storeID int, customerid int, eventtime string) []CustomerCheckout {

	appointmentime, _ := time.Parse("2006-01-02T15:04:05", eventtime)
	t := appointmentime.Format("2006-01-02")
	db := mysqldb.Connect()
	defer db.Close()

	var events []CustomerCheckout

	results, err := db.Query(`SELECT appointments.id,
		appointments.user_id,
		appointments.customer_id,
		customers.customer_name,
		customers.customer_surname,
		customers.customer_email,
		appointments.store_id,
		appointments.hairdresser_id,
		appointments.service_status,
		appointments.start_time,
		appointments.end_time,
		appointments.is_all_day,
		appointments.comments,
		services.service_name,
        services.service_price,
		services.service_discount,
		services.id,
		hairdressers.hairdresser_name
		FROM customers
		INNER JOIN appointments ON customers.id = appointments.customer_id
		JOIN services ON services.id = appointments.service_id
		JOIN hairdressers ON hairdressers.id = appointments.hairdresser_id
		WHERE appointments.store_id= ? AND appointments.user_id = ? AND appointments.service_status = 'pending' AND customers.id = ? AND DATE(appointments.end_time) = ? `, storeID, userID, customerid, t)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var event CustomerCheckout

		err = results.Scan(&event.ID, &event.UserID, &event.CustomerID, &event.CustomerName, &event.CustomerSurname, &event.CustomerEmail, &event.StoreID, &event.ResourceID, &event.ServiceStatus, &event.EventStart, &event.EventEnd, &event.IsAllDay, &event.RantevouComments, &event.ServiceName, &event.ServicePrice, &event.ServiceDiscount, &event.ServiceID, &event.HairdresserName)
		if err != nil {
			println(err.Error())
		}

		events = append(events, event)
		AllAppointmets = events
		// fmt.Printf("%+v", AllAppointmets)

	}

	return events
}

func AllAppointmentPerCustomerPerDayJSON(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(CheckoutConfig{AllCheckout: AllAppointmets})
	if err != nil {
		println(err.Error())
	}
}

func InsertCheckoutDataToDB(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	customerServicesJSON := r.PostFormValue("customerServicesJSON")
	var checkout []CheckOutData
	err := json.Unmarshal([]byte(customerServicesJSON), &checkout)

	if err != nil {
		println(err.Error())

	}
	customerServicesJSONPoints := r.PostFormValue("customerPoints")
	var points int
	err = json.Unmarshal([]byte(customerServicesJSONPoints), &points)
	if err != nil {
		println(err.Error())
	}
	go func() {
		err = UpdateCustomersPoints(points, checkout[0].CustomerID)
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed update customer points"})

			return
		}
	}()

	var id int64
	for _, c := range checkout {
		go func() {
			err := UpdateStatusCompleted(c.EventID)
			if err != nil {
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed update appointment status for: " + c.ServiceName})

				return
			}
		}()

		if c.IsNew == true {
			idstring := make(chan string, 1)
			go func() {
				id, err := InsertappointmentFromCheckout(c.UserID, c.HairdresserID, c.CustomerID, c.StoreID, c.ServiceID, c.ServiceStarTime, c.ServiceEndTime, c.ServiceComments)
				if err != nil {
					println(err.Error())
					json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed create appointmentfor the service: " + c.ServiceName})
					return

				}
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				idstring <- id
			}()
			appID := <-idstring
			appointmentID, _ := strconv.Atoi(appID)
			c.EventID = appointmentID
		}
		stmt, _ := db.Prepare(`INSERT INTO salestrans (user_id,
			customer_id,
			store_id,
			service_id,
			service_price,
			service_qty,
			service_discount,
			payment_type,
			service_line_total,
			is_service,
			invoice_number,
			created_at,
			updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
		res, err := stmt.Exec(c.UserID, c.CustomerID, c.StoreID, c.ServiceID, c.ServicePrice, c.ServiceQty, c.ServiceDiscount, c.PaymentType, c.ServiceLineTotal, c.IsService, c.InvoiceNumber)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert customer services into database"})
		}
		id, err = res.LastInsertId()
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
		}
		idInt := int(id)
		if c.IsService == true {
			id, err = insertSalesTransCommonFromCheckout(idInt, c.HairdresserID, c.EventID)
			if err != nil {
				print(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert service into database"})

			}
		}

		if c.IsService == false {
			err = UpdateProductQty(c.ServiceID, c.ServiceQty)
			if err != nil {
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed update product qty"})
			}
			err = insertSalesProductsPromotions(c.UserID, c.StoreID, c.PromotionProductID, idInt, c.CustomerID)
			if err != nil {
				println(err.Error())
			}
		}

	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted ustomer services into the Database", Body: strconv.Itoa(int(id))})

}
func UpdateCustomersPoints(points, customerID int) error {
	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE customers SET customer_points = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(points, customerID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatusCompleted(id int) error {
	db := mysqldb.Connect()
	defer db.Close()
	status := "complete"
	stmt, err := db.Prepare("UPDATE appointments SET service_status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductQty(productID, qty int) error {
	db := mysqldb.Connect()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE products SET qty = qty - ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(qty, productID)
	if err != nil {
		return err
	}
	return nil
}

func insertSalesTransCommonFromCheckout(salesID, hairdresserID, appointmentID int) (int64, error) {
	db := mysqldb.Connect()
	defer db.Close()
	stmt, err := db.Prepare(`INSERT INTO salestrans_common (salestrans_id,
		hairdresser_id,
		appointment_id,
		created_at,
		updated_at) VALUES (?,?,?, NOW(), NOW())`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(salesID, hairdresserID, appointmentID)

	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

//InsertappointmentFromCheckout insert to db service that comes from checkout
func InsertappointmentFromCheckout(userID, hairdresserID, customerID, storeID, servicID int, startTime, endTime time.Time, comments string) (string, error) {
	db := mysqldb.Connect()
	defer db.Close()
	serviceStatus := "complete"
	isAllDay := false
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
		return "", err
	}
	res, err := stmt.Exec(userID, hairdresserID, customerID, storeID, servicID, startTime, endTime, serviceStatus, comments, isAllDay)

	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func InsertInvoiceJSON(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var invoice Invoices

	err := decoder.Decode(&invoice)
	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`INSERT INTO invoices (user_id,
		store_id,
		customer_id,
		sub_total,
		discount,
		total_cost,
		status_id,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}
	res, err := stmt.Exec(invoice.UserID, invoice.StoreID, invoice.CustomerID, invoice.SubTotalCost, invoice.Discount, invoice.Total, invoice.StatusID)

	if err != nil {
		println(err.Error())
	}
	_, err = res.LastInsertId()
	if err != nil {
		println(err.Error())
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, InvoiceNumber: invoice.StatusID, Description: "Failed update product qty"})

}

func AllInvoicesPerCustomer(storeID int) []Invoices {

	db := mysqldb.Connect()
	defer db.Close()

	var invoices []Invoices

	results, err := db.Query(`SELECT user_id,
		store_id,
		customer_id,
		sub_total,
		discount,
		total_cost,
		status_id,
		created_at,
		updated_at
		FROM invoices
		WHERE store_id = ? `, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var invoice Invoices

		err = results.Scan(&invoice.UserID, &invoice.StoreID, &invoice.CustomerID, &invoice.SubTotalCost, &invoice.Discount, &invoice.Total, &invoice.StatusID, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {
			println(err.Error())
		}

		invoices = append(invoices, invoice)

	}

	return invoices
}
func AllInvoicesPerCustomerJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID, _ := strconv.Atoi(vars["storeid"])

	allInvoicesPerCustomer := AllInvoicesPerCustomer(storeID)
	err := json.NewEncoder(w).Encode(allInvoicesPerCustomer)
	if err != nil {
		println(err.Error())
	}
}
func insertSalesProductsPromotions(userID, storeID, promotionID, salestransID, customerID int) error {
	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO sales_products_with_promotions (user_id,
		store_id,
		customer_id,
		promotion_id,
		salestrans_id,
		created_at,
		updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}
	res, err := stmt.Exec(userID, storeID, customerID, promotionID, salestransID)

	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
