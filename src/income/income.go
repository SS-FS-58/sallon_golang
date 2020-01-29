package income

import (
	"bookings"
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"
)

type Income struct {
	Title     string
	Active    string
	MenuTitle string
	Username  string
	VatNumber string
	ImageName int
	UserID    int
}

type IncomeCheckout struct {
	SalesTransID     int       `json:"sales_trans_id"`
	UserID           int       `json:"user_id"`
	HairdresserID    int       `json:"hairdresser_id"`
	HairdresserName  string    `json:"hairdresser_name"`
	StoreID          int       `json:"store_id"`
	CustomerID       int       `json:"customer_id"`
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

type IncomeInvoice struct {
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
type HTTPResp struct {
	Status        int    `json:"status"`
	Description   string `json:"description"`
	Body          string `json:"body"`
	InvoiceNumber int    `json:"invoice_number"`
}

func InsertIncomeCheckoutDataToDB(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	customerProductsJSON := r.PostFormValue("customerProductsJSON")
	var checkout []IncomeCheckout
	err := json.Unmarshal([]byte(customerProductsJSON), &checkout)

	if err != nil {
		println(err.Error())

	}
	var id int64
	for _, c := range checkout {

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
		err = bookings.UpdateProductQty(c.ServiceID, c.ServiceQty)
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed update product qty"})
		}

	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted ustomer services into the Database", Body: strconv.Itoa(int(id))})

}
