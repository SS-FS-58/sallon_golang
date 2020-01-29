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

//Categories holds the data for all store users
type Categories struct {
	Active       string
	Title        string
	Username     string
	ImageName    int
	UserID       int
	UserProfile  string
	CompanyName  string
	Role         string
	MenuTitle    string
	MessageError string
	Categories   []Category
}

type Category struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Stores       []string  `json:"stores"`
	CategoryName string    `json:"category_name"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type ServiceCategoriesStores struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	StoreID    int       `json:"store_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//CategoryConfig holds the data for json structure
type CategoryConfig struct {
	AllCategories []Category `json:"data"`
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var category Category

	err := decoder.Decode(&category)
	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM service_categories WHERE category_name= ? and user_id = ?", category.CategoryName, category.UserID) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Category exists"})
		return
	}
	category.IsActive = true

	stmt, _ := db.Prepare(`INSERT INTO service_categories (user_id,
		category_name,
		is_active,
		created_at,
		updated_at) VALUES (?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(category.UserID, category.CategoryName, category.IsActive)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert category into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted category Into the Database", Body: strconv.Itoa(int(id))})
}
func GetAllCategoriesPerUser(userID int) []Category {
	db := mysqldb.Connect()
	defer db.Close()

	var categories []Category

	results, err := db.Query(`SELECT id, user_id,
		category_name ,
		is_active,
		created_at,
		updated_at FROM service_categories where user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var category Category
		err = results.Scan(&category.ID, &category.UserID, &category.CategoryName, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		categories = append(categories, category)
	}
	return categories
}

func AllCategoriesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	categories := GetAllCategoriesPerUser(userID)
	err = json.NewEncoder(w).Encode(CategoryConfig{AllCategories: categories})
	if err != nil {
		println(err.Error())
	}
}

//UpdateCategoryJSON update the shop to the database using json
func UpdateCategoryJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var category Category
	err := decoder.Decode(&category)

	vars := mux.Vars(r)

	categoryID := vars["id"]
	ID, err := strconv.Atoi(categoryID)

	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM service_categories WHERE user_id = ? and category_name= ?", category.UserID, category.CategoryName) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Category exists"})
		return
	}

	stmt, err := db.Prepare(`UPDATE service_categories SET
			user_id = ?,
			category_name = ?,
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(category.UserID, category.CategoryName, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update category in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update category in the Database"})
}

//DisableCategory update the shop if its active or not
func DisableCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var category Category
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

	stmt, _ := db.Prepare("UPDATE service_categories SET is_active = ? WHERE id = ?")
	_, err = stmt.Exec(category.IsActive, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update category in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update category in the Database"})
}

//DeleteCategory delete a service
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	categoryID := vars["id"]
	categoryname := vars["categoryname"]
	userid := vars["userid"]

	ID, _ := strconv.Atoi(categoryID)
	if rowExists("SELECT id FROM service_sub_categories WHERE category_name= ?  and user_id = ?", categoryname, userid) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Cannot delete the category because is not empty. Please check Sub Categories"})
		return
	}
	stmt, err := db.Prepare("DELETE FROM service_categories where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete category from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted category from database "})
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
func (s Services) LengthCategories() int {
	return s.LenCategories
}
