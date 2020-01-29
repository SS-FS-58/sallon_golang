package customers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Customers struct {
	Active               string
	Title                string
	Username             string
	VatNumber            string
	ImageName            int
	UserID               int
	UserEmail            string
	UserProfile          string
	CompanyName          string
	Role                 string
	MenuTitle            string
	MessageError         string
	Customers            []Customer
	SingleCustomer       Customer
	LenCompletedServices int
	LenPending           int
	LenCancelled         int
	TotalServices        float64
	TotalProducts        float64
	TotalDiscount        float64
	Total                float64
}

type Customer struct {
	ID                   int       `json:"id"`
	UserID               int       `json:"user_id"`
	CustomerName         string    `json:"customer_name"`
	CustomerSurname      string    `json:"customer_surname"`
	CustomerGender       string    `json:"gender_of_customer"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	CustomerAddress      string    `json:"customer_address"`
	CustomerStreetNumber int       `json:"customer_street_number"`
	CustomerCity         string    `json:"customer_city"`
	CustomerState        string    `json:"customer_state"`
	CustomerZipCode      string    `json:"customer_zip_code"`
	CustomerCountry      string    `json:"customer_country"`
	HomeTelephone        string    `json:"home_phone_number"`
	MobileTelephone      string    `json:"mobile_phone_number"`
	CustomerEmail        string    `json:"customer_email"`
	CustomerPoints       int       `json:"customer_points"`
	IsActive             bool      `json:"is_active"`
	PelatisLianikis      bool      `json:"pelatis_lianikis"`
	TitleCelebration     string    `json:"-"`
	CelebreationType     string    `json:"-"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
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

//CustomerConfig holds the data for json structure
type CustomerConfig struct {
	AllCustomers []Customer `json:"data"`
}

type HTTPResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
	VatNumber   bool   `json:"vat_len"`
}

func rowExists(query string, args ...interface{}) bool {
	db := mysqldb.Connect()
	defer db.Close()
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

//CreateCustomer create a single shop using REST API
func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var customer Customer

	err := decoder.Decode(&customer)

	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM customers where customer_email = ?", &customer.CustomerEmail) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Category exists"})
		return
	}

	stmt, err := db.Prepare(`INSERT INTO customers (user_id,
		customer_name,
		customer_surname,
		gender_of_customer,
		date_of_birth,
		customer_address,
		customer_street_number,
		customer_city,
		customer_state,
		customer_zip_code,
		customer_country,
		home_phone_number,
		mobile_phone_number,
		customer_email,
		customer_points,
		is_active,
		pelatis_lianikis,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}
	customer.IsActive = true
	customer.DateOfBirth.Format("01/02/2006")
	res, _ := stmt.Exec(customer.UserID, customer.CustomerName, customer.CustomerSurname, customer.CustomerGender, customer.DateOfBirth, customer.CustomerAddress, customer.CustomerStreetNumber, customer.CustomerCity, customer.CustomerState, customer.CustomerZipCode, customer.CustomerCountry, customer.HomeTelephone, customer.MobileTelephone, customer.CustomerEmail, customer.CustomerPoints, customer.IsActive, customer.PelatisLianikis)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert customer into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted customer Into the Database", Body: strconv.Itoa(int(id))})

}

func UpdateCustomerJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var customer Customer
	err := decoder.Decode(&customer)

	vars := mux.Vars(r)

	customerID := vars["id"]
	ID, err := strconv.Atoi(customerID)

	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`UPDATE customers SET
			user_id = ?,
			customer_name = ?,
			customer_surname = ?,
			gender_of_customer = ?,
			date_of_birth = ?,
			customer_address = ?,
			customer_street_number = ?,
			customer_city = ?,
			customer_state = ?,
			customer_zip_code = ?,
			customer_country = ?,
			home_phone_number =?,
			mobile_phone_number = ?,
			customer_email = ?, 
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(customer.UserID, customer.CustomerName, customer.CustomerSurname, customer.CustomerGender, customer.DateOfBirth, customer.CustomerAddress, customer.CustomerStreetNumber, customer.CustomerCity, customer.CustomerState, customer.CustomerZipCode, customer.CustomerCountry, customer.HomeTelephone, customer.MobileTelephone, customer.CustomerEmail, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update customer  in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update customer in the Database"})
}

func GetAllCustomersPerUser(userID int) []Customer {
	db := mysqldb.Connect()
	defer db.Close()

	var customers []Customer

	results, err := db.Query(`SELECT id, user_id,
		customer_name,
		customer_surname,
		gender_of_customer,
		date_of_birth,
		customer_address,
		customer_street_number,
		customer_city,
		customer_state,
		customer_zip_code,
		customer_country,
		home_phone_number,
		mobile_phone_number,
		customer_email,
		customer_points,
		is_active,
		pelatis_lianikis,
		created_at,
		updated_at FROM customers where user_id = ?`, userID)
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
func AllCustomersJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	customers := GetAllCustomersPerUser(userID)
	err = json.NewEncoder(w).Encode(CustomerConfig{AllCustomers: customers})
	if err != nil {
		println(err.Error())
	}
}

func SingleCustomerPerUser(id, user_id int) Customer {
	db := mysqldb.Connect()
	defer db.Close()

	var customer = Customer{}
	db.QueryRow(`SELECT id, user_id,
		customer_name,
		customer_surname,
		gender_of_customer,
		date_of_birth,
		customer_address,
		customer_street_number,
		customer_city,
		customer_state,
		customer_zip_code,
		customer_country,
		home_phone_number,
		mobile_phone_number,
		customer_email,
		customer_points,
		is_active,
		pelatis_lianikis,
		created_at,
		updated_at FROM customers WHERE id = ? AND user_id = ?`, id, user_id).Scan(&customer.ID, &customer.UserID, &customer.CustomerName, &customer.CustomerSurname, &customer.CustomerGender, &customer.DateOfBirth, &customer.CustomerAddress, &customer.CustomerStreetNumber, &customer.CustomerCity, &customer.CustomerState, &customer.CustomerZipCode, &customer.CustomerCountry, &customer.HomeTelephone, &customer.MobileTelephone, &customer.CustomerEmail, &customer.CustomerPoints, &customer.IsActive, &customer.PelatisLianikis, &customer.CreatedAt, &customer.UpdatedAt)

	return customer
}
