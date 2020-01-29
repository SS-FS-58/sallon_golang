package promotions

import (
	"database/sql"
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Promotions struct {
	Title     string
	Active    string
	MenuTitle string
	Username  string
	VatNumber string
	ImageName int
	UserID    int
}
type Promotion struct {
	ID                     int       `json:"id"`
	PromotionCommonID      int       `json:"promotion_common_id"`
	StoreID                int       `json:"store_id"`
	StoreName              string    `json:"store_name"`
	UserID                 int       `json:"user_id"`
	PromotionTitle         string    `json:"promotion_title"`
	IsService              bool      `json:"is_service"`
	PromotionServiceName   string    `json:"promotion_service"`
	ServicesIDs            string    `json:"services_ids"`
	ServiceID              int       `json:"service_id"`
	ServiceName            string    `json:"services_name"`
	ServiceDuration        int       `json:"service_duration"`
	PromotionServicesNames string    `json:"promotion_services_names"`
	DaysDuration           int       `json:"days_duration"`
	PromotionSale          float64   `json:"promotion_sale"`
	PromotionDescription   string    `json:"promotion_description"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type PromotionProducts struct {
	ID                     int       `json:"promotion_id"`
	PromotionCommonID      int       `json:"promotion_common_id"`
	StoreID                int       `json:"store_id"`
	StoreName              string    `json:"store_name"`
	UserID                 int       `json:"user_id"`
	PromotionTitle         string    `json:"promotion_title"`
	IsService              bool      `json:"is_service"`
	PromotionServiceName   string    `json:"promotion_service"`
	ServicesIDs            string    `json:"services_ids"`
	ServiceID              int       `json:"service_id"`
	ServiceName            string    `json:"services_name"`
	ServiceDuration        int       `json:"service_duration"`
	PromotionServicesNames string    `json:"promotion_services_names"`
	DaysDuration           int       `json:"days_duration"`
	PromotionSale          float64   `json:"promotion_sale"`
	PromotionDescription   string    `json:"promotion_description"`
	ProductPrice           float64   `json:"product_price"`
	ProductID              int       `json:"product_id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type InsertPromotion struct {
	ID                   int     `json:"id"`
	StoreID              int     `json:"store_id"`
	UserID               int     `json:"user_id"`
	PromotionTitle       string  `json:"promotion_title"`
	IsService            bool    `json:"is_service"`
	PromotionServices    string  `json:"promotion_service"`
	DaysDuration         int     `json:"days_duration"`
	PromotionSale        float64 `json:"promotion_sale"`
	PromotionDescription string  `json:"promotion_description"`
}
type HTTPResp struct {
	Status        int    `json:"status"`
	Description   string `json:"description"`
	Body          string `json:"body"`
	InvoiceNumber int    `json:"invoice_number"`
}
type AllPromotionsServicesConfig struct {
	AllServicesPromotions []Promotion `json:"data"`
}

type AllPromotionsProductsConfig struct {
	AllProductsPromotions []Promotion `json:"data"`
}

func InsertPromotionDataToDB(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	promotionJSON := r.PostFormValue("promotionJSON")
	var promotion InsertPromotion
	err := json.Unmarshal([]byte(promotionJSON), &promotion)

	if err != nil {
		println(err.Error())

	}

	allServices := strings.Split(promotion.PromotionServices, ",")
	var id int64

	promotionID, err := insertSinglePromotion(promotion)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert promotion in the database: " + err.Error()})

	}
	for _, s := range allServices {
		sID, _ := strconv.Atoi(s)
		stmt, _ := db.Prepare(`INSERT INTO promotions_commons (promotion_id,
			promotion_service_id,
			created_at,
			updated_at) VALUES (?,?, NOW(), NOW())`)
		res, err := stmt.Exec(int(promotionID), sID)
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert promotion ids in the database: " + err.Error()})

		}
		id, err = res.LastInsertId()
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed get last insert id from database: " + err.Error()})

		}

	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted promotion into the Database", Body: strconv.Itoa(int(id))})

}

func insertSinglePromotion(promotion InsertPromotion) (int64, error) {
	db := mysqldb.Connect()
	defer db.Close()
	stmt, _ := db.Prepare(`INSERT INTO promotions (user_id,
		store_id,
		promotion_title,
		days_duration,
		promotion_sale,
		promotion_description,
		is_service,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(promotion.UserID, promotion.StoreID, promotion.PromotionTitle, promotion.DaysDuration, promotion.PromotionSale, promotion.PromotionDescription, promotion.IsService)

	if err != nil {
		return 0, err

	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err

	}
	return id, nil
}

func UpdateSinglePromotionDataToDB(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	promID, _ := strconv.Atoi(vars["promid"])
	db := mysqldb.Connect()
	defer db.Close()

	promotionJSON := r.PostFormValue("epromotionJSON")
	var promotion InsertPromotion
	err := json.Unmarshal([]byte(promotionJSON), &promotion)

	if err != nil {
		println(err.Error())

	}
	err = deleteSinglePromotion(promID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete promotion from  database: " + err.Error()})
		return
	}
	allServices := strings.Split(promotion.PromotionServices, ",")
	var id int64

	promotionID, err := insertSinglePromotion(promotion)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert promotion in the database: " + err.Error()})

	}
	for _, s := range allServices {
		sID, _ := strconv.Atoi(s)
		stmt, _ := db.Prepare(`INSERT INTO promotions_commons (promotion_id,
			promotion_service_id,
			created_at,
			updated_at) VALUES (?,?, NOW(), NOW())`)
		res, err := stmt.Exec(int(promotionID), sID)
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert promotion ids in the database: " + err.Error()})

		}
		id, err = res.LastInsertId()
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed get last insert id from database: " + err.Error()})

		}

	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted promotion into the Database", Body: strconv.Itoa(int(id))})

}
func deleteSinglePromotion(promID int) error {
	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM promotions where id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(promID)

	if err != nil {
		return err
	}
	return nil
}

func AllPromotionPerUser(userID int) []Promotion {

	db := mysqldb.Connect()
	defer db.Close()

	var promotions []Promotion

	results, err := db.Query(`SELECT
		promotions_id,
		promotion_title,
		GROUP_CONCAT(service_name) AS services,
		GROUP_CONCAT(s_id) as service_ids,
		days_duration,
		promotion_sale,
		store_id,
		user_id,
		company_name,
		promotion_description,
		is_service,
		created_at,
		updated_at
	   FROM
		 promotion_view
	 WHERE user_id = ?
	   GROUP BY
		 promotions_id`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var p Promotion

		err = results.Scan(&p.ID, &p.PromotionTitle, &p.PromotionServiceName, &p.ServicesIDs, &p.DaysDuration, &p.PromotionSale, &p.StoreID, &p.UserID, &p.StoreName, &p.PromotionDescription, &p.IsService, &p.CreatedAt, &p.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		promotions = append(promotions, p)
		// fmt.Printf("%+v", AllAppointmets)

	}

	return promotions
}
func AllPromotionServicesJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])

	allServicesPromotion := AllPromotionPerUser(userID)
	var promotionServices []Promotion
	for _, s := range allServicesPromotion {

		if s.IsService == true {

			promotionServices = append(promotionServices, s)
		}

	}

	err := json.NewEncoder(w).Encode(AllPromotionsServicesConfig{AllServicesPromotions: promotionServices})
	if err != nil {
		println(err.Error())
	}
}
func AllPromotionServicesPerUserPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])
	storeID, _ := strconv.Atoi(vars["storeid"])
	serviceID, _ := strconv.Atoi(vars["serviceid"])

	db := mysqldb.Connect()
	defer db.Close()

	var promotions []Promotion

	results, err := db.Query(`SELECT promotions.id AS promotion_id,
		promotions.user_id,
		promotions.store_id,
		promotions.promotion_title,
		promotions.days_duration,
		promotions.promotion_sale,
		promotions.promotion_description,
		promotions.is_service,
		promotions_commons.promotion_service_id,
		services.service_name,
		services.service_duration,
		promotions.created_at,
		promotions.updated_at
		FROM promotions_commons
		LEFT JOIN promotions ON promotions_commons.promotion_id = promotions.id
		LEFT JOIN services On promotions_commons.promotion_service_id =services.id
		WHERE promotions.user_id = ? AND promotions.store_id = ? AND promotions_commons.promotion_service_id = ? AND promotions.is_service = true
	  	ORDER BY promotions.created_at DESC`, userID, storeID, serviceID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var p Promotion

		err = results.Scan(&p.ID, &p.UserID, &p.StoreID, &p.PromotionTitle, &p.DaysDuration, &p.PromotionSale, &p.PromotionDescription, &p.IsService, &p.ServiceID, &p.ServiceName, &p.ServiceDuration, &p.CreatedAt, &p.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		promotions = append(promotions, p)
		// fmt.Printf("%+v", AllAppointmets)

	}

	err = json.NewEncoder(w).Encode(promotions)
	if err != nil {
		println(err.Error())
	}
}

func AllPromotionProductsPerUserPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])
	storeID, _ := strconv.Atoi(vars["storeid"])
	productID, _ := strconv.Atoi(vars["productid"])

	db := mysqldb.Connect()
	defer db.Close()

	var promotions []PromotionProducts

	results, err := db.Query(`SELECT promotions.id AS promotion_id,
		promotions.user_id,
		promotions.store_id,
		promotions.promotion_title,
		promotions.promotion_sale,
		promotions.promotion_description,
		promotions.is_service,
		promotions_commons.promotion_service_id,
		products.product_name,
		products.product_price,
		products.id,
		promotions.created_at,
		promotions.updated_at
		FROM promotions_commons
		LEFT JOIN promotions ON promotions_commons.promotion_id = promotions.id
		LEFT JOIN products ON promotions_commons.promotion_service_id =products.id
		WHERE promotions.user_id = ? AND promotions.store_id = ? AND promotions_commons.promotion_service_id = ? AND promotions.is_service = false
	  	ORDER BY promotions.created_at DESC`, userID, storeID, productID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var p PromotionProducts

		err = results.Scan(&p.ID, &p.UserID, &p.StoreID, &p.PromotionTitle, &p.PromotionSale, &p.PromotionDescription, &p.IsService, &p.ServiceID, &p.ServiceName, &p.ProductPrice, &p.ProductID, &p.CreatedAt, &p.UpdatedAt)
		if err != sql.ErrNoRows && err != nil {
			println(err.Error())
		}

		promotions = append(promotions, p)
		// fmt.Printf("%+v", AllAppointmets)

	}

	err = json.NewEncoder(w).Encode(promotions)
	if err != nil {
		println(err.Error())
	}
}
func AllPromotionProductsJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])

	allProductsPromotion := AllPromotionPerUser(userID)
	var promotionProducts []Promotion
	for _, s := range allProductsPromotion {

		if s.IsService == false {

			promotionProducts = append(promotionProducts, s)
		}

	}

	err := json.NewEncoder(w).Encode(AllPromotionsProductsConfig{AllProductsPromotions: promotionProducts})
	if err != nil {
		println(err.Error())
	}
}
func removeDuplicates(elements []Promotion) []Promotion {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []Promotion{}

	for v := range elements {
		if encountered[elements[v].ID] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v].ID] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func joinDuplicates(elements []Promotion) []string {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v].ID] == true {
			// Do not add duplicate.
			result = append(result, elements[v].PromotionServiceName)

		} else {
			// Record this element as an encountered element.
			encountered[elements[v].ID] = true
			// Append to result slice.
		}
	}
	// Return the new slice.
	return result
}
func Deletepromotion(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	promID := vars["promid"]

	ID, _ := strconv.Atoi(promID)

	err := deleteSinglePromotion(ID)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete promotion from database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted promotion from database "})
}
