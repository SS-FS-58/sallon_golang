package expenses

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

type Expenses struct {
	Title     string
	Active    string
	MenuTitle string
	Username  string
	VatNumber string
	ImageName int
	UserID    int
}

type CostCategoryName struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	StoreID          int       `json:"store_id"`
	StoreName        string    `json:"store_name"`
	CostCategoryName string    `json:"cost_category"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type HTTPResp struct {
	Status        int    `json:"status"`
	Description   string `json:"description"`
	Body          string `json:"body"`
	InvoiceNumber int    `json:"invoice_number"`
}

//CategoryConfig holds the data for json structure
type ExpensesConfig struct {
	AllCostCategories []CostCategoryName `json:"data"`
}

func CreateCostCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var category CostCategoryName

	err := decoder.Decode(&category)
	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM expenses_list WHERE user_id = ? and store_id = ? and expenses_name= ?", category.UserID, category.StoreID, category.CostCategoryName) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Category exists"})
		return
	}
	category.IsActive = true

	stmt, _ := db.Prepare(`INSERT INTO expenses_list (user_id,
		store_id,
		expenses_name,
		is_active,
		created_at,
		updated_at) VALUES (?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(category.UserID, category.StoreID, category.CostCategoryName, category.IsActive)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert cost category into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted cost category Into the Database", Body: strconv.Itoa(int(id))})
}
func GetAllCostCategoriesPerUser(userID int) []CostCategoryName {
	db := mysqldb.Connect()
	defer db.Close()

	var categories []CostCategoryName

	results, err := db.Query(`SELECT expenses_list.id, expenses_list.user_id,
		expenses_list.store_id,
		shops.company_name,
		expenses_list.expenses_name ,
		expenses_list.is_active,
		expenses_list.created_at,
		expenses_list.updated_at 
		FROM expenses_list 
		LEFT JOIN shops ON expenses_list.store_id = shops.id
		WHERE expenses_list.user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var category CostCategoryName
		err = results.Scan(&category.ID, &category.UserID, &category.StoreID, &category.StoreName, &category.CostCategoryName, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		categories = append(categories, category)
	}
	return categories
}

func AllCostCategoriesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	categories := GetAllCostCategoriesPerUser(userID)
	err = json.NewEncoder(w).Encode(ExpensesConfig{AllCostCategories: categories})
	if err != nil {
		println(err.Error())
	}
}

//UpdateCostCategoryJSON update the shop to the database using json
func UpdateCostCategoryJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var category CostCategoryName
	err := decoder.Decode(&category)

	vars := mux.Vars(r)

	categoryID := vars["id"]
	ID, err := strconv.Atoi(categoryID)

	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM expenses_list WHERE user_id = ? and store_id = ? and expenses_name= ?", category.UserID, category.StoreID, category.CostCategoryName) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Cost Category exists"})
		return
	}

	stmt, err := db.Prepare(`UPDATE expenses_list SET
			expenses_name = ?,
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(category.CostCategoryName, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update cost category in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update cost category in the Database"})
}

//DisableCostCategory update the shop if its active or not
func DisableCostCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var category CostCategoryName
	err := decoder.Decode(&category)

	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	categoryID := vars["id"]
	ID, err := strconv.Atoi(categoryID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE expenses_list SET is_active = ? WHERE id = ?")
	_, err = stmt.Exec(category.IsActive, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update cost category in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update cost category in the Database"})
}

//DeleteCostCategory delete a service
func DeleteCostCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	categoryID := vars["id"]

	ID, _ := strconv.Atoi(categoryID)

	stmt, err := db.Prepare("DELETE FROM expenses_list where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete cost category from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted cost category from database "})
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
