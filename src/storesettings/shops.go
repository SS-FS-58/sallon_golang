package storesettings

import (
	"encoding/json"
	"fmt"
	"log"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type HTTPResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
	VatNumber   bool   `json:"vat_len"`
}

//StoreUsers holds the data for all store users
type Shops struct {
	Active       string
	Title        string
	Username     string
	VatNumber    string
	ImageName    int
	UserID       int
	UserProfile  string
	CompanyName  string
	Role         string
	MenuTitle    string
	MessageError string
}

type Shop struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	CompanyName         string    `json:"company_name"`
	VatNumber           string    `json:"vat_number"`
	TaxOffice           string    `json:"tax_office"`
	CompanyAddress      string    `json:"company_address"`
	CompanyStreetNumber int       `json:"company_street_number"`
	CompanyCity         string    `json:"company_city"`
	CompanyState        string    `json:"company_state"`
	CompanyZipCode      string    `json:"company_zip_code"`
	CompanyCountry      string    `json:"company_country"`
	HomeTelephone       string    `json:"work_telephone"`
	MobileTelephone     string    `json:"mobile_telephone"`
	Password            string    `json:"password"`
	IsActive            bool      `json:"is_active"`
	IncludeBankHolidays string    `json:"include_bank_holidays"`
	BankHolidaysCountry string    `json:"bank_holidays_country"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

//ShopsConfig holds the data for json structure
type ShopsConfig struct {
	Shops []Shop `json:"data"`
}

//CreateShops create a single shop using REST API
func CreateShops(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var shop Shop

	err := decoder.Decode(&shop)

	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM shops WHERE vat_number= ?  and user_id = ?", shop.VatNumber, shop.UserID) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "The Vat Number already exists"})
		return
	}
	shop.IsActive = true
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shop.Password), bcrypt.DefaultCost)
	if err != nil {
		print(err.Error())
	}

	stmt, _ := db.Prepare(`INSERT INTO shops (user_id,
		vat_number,
		company_name,
		company_address,
		company_street_number,
		company_city,
		company_state,
		company_zip_code,
		tax_office,
		company_country,
		work_telephone,
		mobile_telephone,
		password, 
		is_active,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(shop.UserID, shop.VatNumber, shop.CompanyName, shop.CompanyAddress, shop.CompanyStreetNumber, shop.CompanyCity, shop.CompanyState, shop.CompanyZipCode, shop.TaxOffice, shop.CompanyCountry, shop.HomeTelephone, shop.MobileTelephone, hashedPassword, shop.IsActive)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert user into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	if shop.IncludeBankHolidays == "include" {
		var intID int = int(id)
		stmt1, _ := db.Prepare(`INSERT INTO bank_holidays_per_store (user_id,
			store_id,
			bank_holidays_country,
			created_at,
			updated_at) VALUES (?,?,?, NOW(), NOW())`)
		res1, err := stmt1.Exec(shop.UserID, intID, shop.BankHolidaysCountry)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert user into database"})
		}
		id1, err := res1.LastInsertId()
		if err != nil {
			println(err.Error(), id1)
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
		}
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted Post Into the Database", Body: strconv.Itoa(int(id))})

}

func UpdateShopPassword(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var shop Shop
	err := decoder.Decode(&shop)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	shopID := vars["id"]
	ID, err := strconv.Atoi(shopID)

	if err != nil {
		println(err.Error())
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shop.Password), bcrypt.DefaultCost)
	if err != nil {
		print(err.Error())
	}
	println(shop.Password)
	stmt, _ := db.Prepare("UPDATE shops SET password  = ? WHERE id = ?")
	_, err = stmt.Exec(hashedPassword, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update shop in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update shop in the Database"})

}

func GetAllShopsPerUser(userID int) []Shop {
	db := mysqldb.Connect()
	defer db.Close()

	var shops []Shop

	results, err := db.Query(`SELECT shops.id, shops.user_id,
		shops.vat_number,
		shops.company_name,
		shops.company_address,
		shops.company_street_number,
		shops.company_city,
		shops.company_state,
		shops.company_zip_code,
		shops.tax_office,
		shops.company_country,
		shops.work_telephone,
		shops.mobile_telephone,
		bank_holidays_per_store.bank_holidays_country, 
		shops.is_active,
		shops.created_at,
		shops.updated_at FROM shops
		LEFT JOIN bank_holidays_per_store ON shops.id = bank_holidays_per_store.store_id
		where shops.is_active = 1 and shops.user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var shop Shop
		err = results.Scan(&shop.ID, &shop.UserID, &shop.VatNumber, &shop.CompanyName, &shop.CompanyAddress, &shop.CompanyStreetNumber, &shop.CompanyCity, &shop.CompanyState, &shop.CompanyZipCode, &shop.TaxOffice, &shop.CompanyCountry, &shop.HomeTelephone, &shop.MobileTelephone, &shop.BankHolidaysCountry, &shop.IsActive, &shop.CreatedAt, &shop.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		shops = append(shops, shop)
	}
	return shops
}

func AllShopsJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	shops := GetAllShopsPerUser(userID)
	err = json.NewEncoder(w).Encode(ShopsConfig{Shops: shops})
	if err != nil {
		println(err.Error())
	}
}

//UpdateShopJSON update the shop to the database using json
func UpdateShopJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var shop Shop
	err := decoder.Decode(&shop)

	vars := mux.Vars(r)

	shopID := vars["id"]
	ID, err := strconv.Atoi(shopID)

	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`UPDATE shops SET
				 user_id = ?,
				vat_number = ?,
				company_name = ?,
				company_address = ?,
				company_street_number = ?,
				company_city = ?,
				company_state = ?,
				company_zip_code = ?,
				tax_office = ?,
				company_country = ?,
				work_telephone = ?,
				mobile_telephone = ?, 
				updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(shop.UserID, shop.VatNumber, shop.CompanyName, shop.CompanyAddress, shop.CompanyStreetNumber, shop.CompanyCity, shop.CompanyState, shop.CompanyZipCode, shop.TaxOffice, shop.CompanyCountry, shop.HomeTelephone, shop.MobileTelephone, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update shop in the Database"})
	}

	if shop.BankHolidaysCountry != "" {
		if rowExists("SELECT id FROM bank_holidays_per_store WHERE store_id= ?  and user_id = ?", ID, shop.UserID) {
			stmt1, err := db.Prepare(`UPDATE bank_holidays_per_store SET
				bank_holidays_country = ?,
			   updated_at = NOW() WHERE store_id = ?`)
			_, err = stmt1.Exec(shop.BankHolidaysCountry, ID)
			if err != nil {
				log.Print(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update shop in the Database"})
			}
		} else {
			stmt2, _ := db.Prepare(`INSERT INTO bank_holidays_per_store (user_id,
				store_id,
				bank_holidays_country,
				created_at,
				updated_at) VALUES (?,?,?, NOW(), NOW())`)
			res2, err := stmt2.Exec(shop.UserID, ID, shop.BankHolidaysCountry)

			if err != nil {
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert user into database"})
			}
			id1, err := res2.LastInsertId()
			if err != nil {
				println(err.Error(), id1)
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
			}
		}

	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update shop in the Database"})
}

//DisableShop update the shop if its active or not
func DisableShop(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var shop Shop
	err := decoder.Decode(&shop)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	shopID := vars["id"]
	ID, err := strconv.Atoi(shopID)

	if err != nil {
		println(err.Error())
	}
	shop.IsActive = false
	stmt, _ := db.Prepare("UPDATE shops SET is_active = ? WHERE id = ?")
	_, err = stmt.Exec(shop.IsActive, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update shop in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update shop in the Database"})
}
func QueryShop(vatNumber string) Shop {
	db := mysqldb.Connect()
	defer db.Close()

	var shop = Shop{}
	err := db.QueryRow(`SELECT id, user_id,
		vat_number,
		company_name,
		company_address,
		company_street_number,
		company_city,
		company_state,
		company_zip_code,
		tax_office,
		company_country,
		work_telephone,
		mobile_telephone,
		password, 
		is_active,
		created_at,
		updated_at FROM shops where vat_number = ?`, vatNumber).Scan(&shop.ID, &shop.UserID, &shop.VatNumber, &shop.CompanyName, &shop.CompanyAddress, &shop.CompanyStreetNumber, &shop.CompanyCity, &shop.CompanyState, &shop.CompanyZipCode, &shop.TaxOffice, &shop.CompanyCountry, &shop.HomeTelephone, &shop.MobileTelephone, &shop.Password, &shop.IsActive, &shop.CreatedAt, &shop.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
	}
	return shop
}
