package storesettings

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

//Services holds the data for all store users
type Services struct {
	Active        string
	Title         string
	Username      string
	VatNumber     string
	ImageName     int
	UserID        int
	UserProfile   string
	CompanyName   string
	Role          string
	MenuTitle     string
	MessageError  string
	LenCategories int
	Categories    []Category
	SubCategories []SubCategory
	Services      []Service
}

type Service struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CategoryID      int       `json:"category_id"`
	SubCategoryID   int       `json:"sub_category_id"`
	Stores          []string  `json:"stores"`
	ServiceName     string    `json:"service_name"`
	ServiceDuration int       `json:"service_duration"`
	ServicePrice    float64   `json:"service_price"`
	ServiceDiscount float64   `json:"service_discount"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	StoreID         int       `json:"-"`
	StoreName       string    `json:"store_name"`
	SwitchFormula   bool      `json:"switch_formula"`
}

type ServiceView struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	ServiceName     string    `json:"service_name"`
	ServiceDuration int       `json:"service_duration"`
	ServicePrice    float64   `json:"service_price"`
	ServiceDiscount float64   `json:"service_discount"`
	StoreID         int       `json:"store_id"`
	StoreName       string    `json:"store_name"`
	CategoryName    string    `json:"category_name"`
	SubCategoryName string    `json:"sub_category_name"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	SwitchFormula   bool      `json:"switch_formula"`
}

//ServiceConfig holds the data for json structure
type ServiceConfig struct {
	AllServices []ServiceView `json:"data"`
}

func CreateServices(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var service Service

	err := decoder.Decode(&service)
	if err != nil {
		print(err.Error())
	}

	stores := service.Stores
	var allServices []string
	var allStoreNames []string
	var allstoreIDs []int
	var allstoreIDss []int
	var serviceID int
	for _, s := range stores {
		var ser Service
		sint, _ := strconv.Atoi(s)
		err = db.QueryRow(`SELECT first.id,
			first.service_name,
			 third.id,
			 third.company_name
			FROM services first
			 LEFT JOIN services_stores second on first.id = second.service_id 
			 INNER JOIN shops third on second.store_id = third.id 
			 WHERE first.user_id = ? AND first.service_name = ? `, service.UserID, service.ServiceName).Scan(&ser.ID, &ser.ServiceName, &ser.StoreID, &ser.StoreName)

		if err != nil && err != sql.ErrNoRows {
			log.Println(err.Error())

		}
		if err != sql.ErrNoRows {
			allServices = append(allServices, ser.ServiceName)
			allstoreIDs = append(allstoreIDs, ser.StoreID)
			serviceID = ser.ID
			allStoreNames = append(allStoreNames, ser.StoreName)

		}
		if !rowExists("SELECT id FROM services_stores WHERE service_id = ? and store_id= ?", ser.ID, sint) {
			allstoreIDss = append(allstoreIDss, sint)
		}

	}

	if allstoreIDs == nil {

		allcats := QuerySubCategory(service.SubCategoryID)
		stmt, _ := db.Prepare(`INSERT INTO services (user_id,category_id,sub_category_id,
		service_name ,
		service_duration,
		service_price,
		service_discount,
		
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?, NOW(), NOW())`)
		res, err := stmt.Exec(service.UserID, allcats.CategoryID, service.SubCategoryID, service.ServiceName, service.ServiceDuration, service.ServicePrice, service.ServiceDiscount)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert service into database"})
		}
		id, err := res.LastInsertId()
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
		}

		intID := strconv.Itoa(int(id))
		stores := service.Stores

		for _, store := range stores {
			wstmt, err := db.Prepare("INSERT INTO services_stores (service_id, store_id,is_active,switch_formula,created_at, updated_at) VALUES (?,?,1,?,NOW(),NOW())")
			if err != nil {
				println(err.Error())
			}
			j, err := strconv.Atoi(store)
			if err != nil {
				println(err.Error())
			}
			wres, err := wstmt.Exec(intID, j, service.SwitchFormula)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert into database"})
			}
			wid, err := wres.LastInsertId()
			if err != nil {
				println(err.Error(), wid)
			}

		}
		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted service Into the Database", Body: strconv.Itoa(int(id))})

	} else {
		if allstoreIDss == nil {

			json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Service exists"})
			return
		}
		for _, store := range allstoreIDss {

			wstmt, err := db.Prepare("INSERT INTO services_stores (service_id, store_id,is_active,switch_formula,created_at, updated_at) VALUES (?,?,1,?,NOW(),NOW())")
			if err != nil {
				println(err.Error())
			}

			wres, err := wstmt.Exec(serviceID, store, service.SwitchFormula)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert into database"})
			}
			wid, err := wres.LastInsertId()
			if err != nil {
				println(err.Error(), wid)
			}

		}
		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted hairdresser into the Database"})

	}
}
func GetAllServicesPerUser(userID int) []ServiceView {
	db := mysqldb.Connect()
	defer db.Close()

	var services []ServiceView

	results, err := db.Query(`SELECT service_id,
		user_id,
		service_name,
		service_duration,
		service_price,
		service_discount,
		store_id,
		company_name,
		category_name,
		sub_category_name,
		is_active,
		switch_formula,
		created_at,
		updated_at
	FROM services_view where user_id = ? ORDER BY service_name ASC`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var service ServiceView
		err = results.Scan(&service.ID, &service.UserID, &service.ServiceName, &service.ServiceDuration, &service.ServicePrice, &service.ServiceDiscount, &service.StoreID, &service.StoreName, &service.CategoryName, &service.SubCategoryName, &service.IsActive, &service.SwitchFormula, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		services = append(services, service)
	}
	return services
}
func GetAllServicesPerUserPerStore(userID, storeID int) []ServiceView {
	db := mysqldb.Connect()
	defer db.Close()

	var services []ServiceView

	results, err := db.Query(`SELECT service_id,
		user_id,
		service_name,
		service_duration,
		service_price,
		service_discount,
		store_id,
		company_name,
		category_name,
		sub_category_name,
		is_active,
		switch_formula,
		created_at,
		updated_at
	FROM services_view where user_id = ? AND store_id = ? AND is_active = 1 ORDER BY service_name ASC`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var service ServiceView
		err = results.Scan(&service.ID, &service.UserID, &service.ServiceName, &service.ServiceDuration, &service.ServicePrice, &service.ServiceDiscount, &service.StoreID, &service.StoreName, &service.CategoryName, &service.SubCategoryName, &service.IsActive, &service.SwitchFormula, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		services = append(services, service)
	}
	return services
}

func AllServicesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	services := GetAllServicesPerUser(userID)
	err = json.NewEncoder(w).Encode(ServiceConfig{AllServices: services})
	if err != nil {
		println(err.Error())
	}
}

func GetAllServicesPerUserPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}

	storeID, err := strconv.Atoi(vars["storeid"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid store id")
	}

	var services []ServiceView

	results, err := db.Query(`SELECT service_id,
		user_id,
		service_name,
		service_duration,
		service_price,
		service_discount,
		store_id,
		company_name,
		category_name,
		sub_category_name,
		is_active,
		switch_formula,
		created_at,
		updated_at
	FROM services_view where user_id = ? AND store_id = ? AND is_active = 1 ORDER BY service_name ASC`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var service ServiceView
		err = results.Scan(&service.ID, &service.UserID, &service.ServiceName, &service.ServiceDuration, &service.ServicePrice, &service.ServiceDiscount, &service.StoreID, &service.StoreName, &service.CategoryName, &service.SubCategoryName, &service.IsActive, &service.SwitchFormula, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		services = append(services, service)
	}
	err = json.NewEncoder(w).Encode(services)
	if err != nil {
		println(err.Error())
	}
}

//UpdateServiceJSON update the shop to the database using json
func UpdateServiceJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var service Service
	err := decoder.Decode(&service)

	vars := mux.Vars(r)

	serviceID := vars["id"]
	ID, err := strconv.Atoi(serviceID)

	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`UPDATE services SET
			user_id = ?,
			service_name = ? ,
			service_duration = ?,
			service_price = ?,
			service_discount = ?,
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(service.UserID, service.ServiceName, service.ServiceDuration, service.ServicePrice, service.ServiceDiscount, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update service in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update service in the Database"})
}

//DisableService update the shop if its active or not
func DisableService(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var service Service
	err := decoder.Decode(&service)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	serviceID := vars["id"]
	ID, err := strconv.Atoi(serviceID)

	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE services_stores SET is_active = ? WHERE store_id = ? AND service_id = ?")
	_, err = stmt.Exec(service.IsActive, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update service in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update service in the Database"})
}

//DeleteService delete a service
func DeleteService(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	serviceID := vars["id"]

	ID, _ := strconv.Atoi(serviceID)

	stmt, err := db.Prepare("DELETE FROM services where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete service from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted service from database "})
}
func HasFormulaService(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var service Service
	err := decoder.Decode(&service)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	serviceID := vars["id"]
	ID, err := strconv.Atoi(serviceID)

	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE services_stores SET switch_formula = ? WHERE store_id = ? AND service_id = ?")
	_, err = stmt.Exec(service.IsActive, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update service in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update service in the Database"})
}
func QuerySubCategory(subcategoryID int) SubCategory {
	db := mysqldb.Connect()
	defer db.Close()

	var subcategory = SubCategory{}
	err := db.QueryRow("SELECT id, user_id,category_id,category_name,sub_category_name,is_active,created_at,updated_at FROM service_sub_categories where id = ? AND is_active = 1", subcategoryID).Scan(&subcategory.ID, &subcategory.UserID, &subcategory.CategoryID, &subcategory.CategoryName, &subcategory.SubCategoryName, &subcategory.IsActive, &subcategory.CreatedAt, &subcategory.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
	}
	return subcategory
}
