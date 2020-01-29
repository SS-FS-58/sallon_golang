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
)

//Suppliers holds the data for all store users
type Suppliers struct {
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

type Supplier struct {
	ID                   int       `json:"id"`
	UserID               int       `json:"user_id"`
	SupplierName         string    `json:"supplier_name"`
	VatNumber            string    `json:"vat_number"`
	FieldOfBusiness      string    `json:"field_of_business"`
	SupplierAddress      string    `json:"supplier_address"`
	SupplierStreetNumber int       `json:"supplier_street_number"`
	SupplierCity         string    `json:"supplier_city"`
	SupplierState        string    `json:"supplier_state"`
	SupplierZipCode      string    `json:"supplier_zip_code"`
	SupplierCountry      string    `json:"supplier_country"`
	WorkTelephone        string    `json:"work_telephone"`
	Website              string    `json:"website"`
	Email                string    `json:"email"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

//SupplierConfig holds the data for json structure
type SupplierConfig struct {
	Suppliers []Supplier `json:"data"`
}

func QuerySupplier(vatNmuber string) Supplier {
	db := mysqldb.Connect()
	defer db.Close()

	var supplier Supplier
	err := db.QueryRow(`SELECT id, user_id,
		vat_number,
		supplier_name,
		supplier_address,
		supplier_street_number,
		supplier_city,
		supplier_state,
		supplier_zip_code,
		field_of_business,
		supplier_country,
		work_telephone,
		website,
		email, 
		is_active,
		created_at,
		updated_at FROM suppliers where vat_number = ?`, vatNmuber).Scan(&supplier.ID, &supplier.UserID, &supplier.VatNumber, &supplier.SupplierName, &supplier.SupplierAddress, &supplier.SupplierStreetNumber, &supplier.SupplierCity, &supplier.SupplierState, &supplier.SupplierZipCode, &supplier.FieldOfBusiness, &supplier.SupplierCountry, &supplier.WorkTelephone, &supplier.Website, &supplier.Email, &supplier.IsActive, &supplier.CreatedAt, &supplier.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
	}
	return supplier
}

//CreateSupplier create a single shop using REST API
func CreateSupplier(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var supplier Supplier

	err := decoder.Decode(&supplier)

	if err != nil {
		print(err.Error())
	}
	fmt.Printf("%+v", supplier)

	suppliers := QuerySupplier(supplier.VatNumber)
	if len(suppliers.VatNumber) != 0 {

		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "VAT Number already exists"})
		return
	}
	supplier.IsActive = true
	stmt, _ := db.Prepare(`INSERT INTO suppliers (user_id,
		vat_number,
		supplier_name,
		supplier_address,
		supplier_street_number,
		supplier_city,
		supplier_state,
		supplier_zip_code,
		field_of_business,
		supplier_country,
		work_telephone,
		website,
		email, 
		is_active,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(supplier.UserID, supplier.VatNumber, supplier.SupplierName, supplier.SupplierAddress, supplier.SupplierStreetNumber, supplier.SupplierCity, supplier.SupplierState, supplier.SupplierZipCode, supplier.FieldOfBusiness, supplier.SupplierCountry, supplier.WorkTelephone, supplier.Website, supplier.Email, supplier.IsActive)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert supplier into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted supplier Into the Database", Body: strconv.Itoa(int(id))})

}

func CheckIfVatExistsForSuppliers(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var supplier Supplier

	err := decoder.Decode(&supplier)
	if err != nil {
		print(err.Error())
	}

	err = db.QueryRow("SELECT vat_number FROM suppliers where vat_number = ?", supplier.VatNumber).Scan(&supplier.VatNumber)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200})

}

func GetAllSuppliersPerUser(userID int) []Supplier {
	db := mysqldb.Connect()
	defer db.Close()

	var suppliers []Supplier

	results, err := db.Query(`SELECT id, user_id,
		vat_number,
		supplier_name,
		supplier_address,
		supplier_street_number,
		supplier_city,
		supplier_state,
		supplier_zip_code,
		field_of_business,
		supplier_country,
		work_telephone,
		website,
		email, 
		is_active,
		created_at,
		updated_at FROM suppliers where user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var supplier Supplier
		err = results.Scan(&supplier.ID, &supplier.UserID, &supplier.VatNumber, &supplier.SupplierName, &supplier.SupplierAddress, &supplier.SupplierStreetNumber, &supplier.SupplierCity, &supplier.SupplierState, &supplier.SupplierZipCode, &supplier.FieldOfBusiness, &supplier.SupplierCountry, &supplier.WorkTelephone, &supplier.Website, &supplier.Email, &supplier.IsActive, &supplier.CreatedAt, &supplier.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers
}

func AllSuppliersJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	suppliers := GetAllSuppliersPerUser(userID)
	err = json.NewEncoder(w).Encode(SupplierConfig{Suppliers: suppliers})
	if err != nil {
		println(err.Error())
	}
}

//UpdateSupplierJSON update the supplier to the database using json
func UpdateSupplierJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var supplier Supplier
	err := decoder.Decode(&supplier)

	vars := mux.Vars(r)

	supplierID := vars["id"]
	ID, err := strconv.Atoi(supplierID)

	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`UPDATE suppliers SET
		user_id = ?,
		vat_number = ?,
		supplier_name = ?,
		supplier_address = ?,
		supplier_street_number = ?,
		supplier_city = ?,
		supplier_state = ?,
		supplier_zip_code = ?,
		field_of_business = ?,
		supplier_country = ?,
		work_telephone = ?,
		website = ?,
		email = ?, 
		updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(supplier.UserID, supplier.VatNumber, supplier.SupplierName, supplier.SupplierAddress, supplier.SupplierStreetNumber, supplier.SupplierCity, supplier.SupplierState, supplier.SupplierZipCode, supplier.FieldOfBusiness, supplier.SupplierCountry, supplier.WorkTelephone, supplier.Website, supplier.Email, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update supplier in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update supplier in the Database"})
}

//DisableSupplier update the supplier if its active or not
func DisableSupplier(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var supplier Supplier
	err := decoder.Decode(&supplier)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	supplierID := vars["id"]
	ID, err := strconv.Atoi(supplierID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE suppliers SET is_active = ? WHERE id = ?")
	_, err = stmt.Exec(supplier.IsActive, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update supplier in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update supplier in the Database"})
}

//DeleteSupplier delete a service
func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	supplierID := vars["id"]

	decoder := json.NewDecoder(r.Body)
	var supplier Supplier

	err := decoder.Decode(&supplier)
	if err != nil {
		print(err.Error())
	}

	ID, _ := strconv.Atoi(supplierID)

	if rowExists("SELECT id FROM products WHERE supplier_id= ?  and user_id = ?", ID, supplier.UserID) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Cannot delete the category because is not empty. Please check Sub Categories"})
		return
	}

	stmt, err := db.Prepare("DELETE FROM suppliers where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete supplier from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted supplier from database "})
}
