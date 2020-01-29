package storesettings

import (
	"encoding/json"
	"fmt"
	"log"
	"mysqldb"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//SubCategories holds the data for all store users
type SubCategories struct {
	Active        string
	Title         string
	Username      string
	ImageName     int
	UserID        int
	UserProfile   string
	CompanyName   string
	Role          string
	MenuTitle     string
	MessageError  string
	SubCategories []SubCategory
}

type SubCategory struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CategoryID      string    `json:"category_id"`
	CategoryName    string    `json:"category_name"`
	SubCategoryName string    `json:"sub_category_name"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

//SubCategoryConfig holds the data for json structure
type SubCategoryConfig struct {
	AllSubCategories []SubCategory `json:"data"`
}

func CreateSubCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var subcategory SubCategory

	err := decoder.Decode(&subcategory)

	if err != nil {
		print(err.Error())
	}

	subcategory.IsActive = true
	categoryName := strings.Split(subcategory.CategoryName, ":")
	subCategoryTitle := strings.Split(subcategory.SubCategoryName, ":")
	categoryID := strings.Split(subcategory.CategoryID, ":")
	var t2 = []int{}

	for _, i := range categoryID {
		j, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
		}
		t2 = append(t2, j)
	}

	var intID int64
	for i, s := range categoryName {
		if rowExists("SELECT id FROM service_sub_categories WHERE category_name= ? and sub_category_name= ? and user_id = ? ", s, subCategoryTitle[i], subcategory.UserID) {
			json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Sub Category exists: " + subCategoryTitle[i]})
			return
		}
		stmt, _ := db.Prepare(`INSERT INTO service_sub_categories (user_id,
		category_id,
		category_name,
		sub_category_name,
		is_active,
		created_at,
		updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
		res, err := stmt.Exec(subcategory.UserID, t2[i], s, subCategoryTitle[i], subcategory.IsActive)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert sub category into database"})
		}
		intID, err = res.LastInsertId()
		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
		}

	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted sub category Into the Database", Body: strconv.Itoa(int(intID))})
}
func GetAllSubCategoriesPerUser(userID int) []SubCategory {
	db := mysqldb.Connect()
	defer db.Close()

	var subcategories []SubCategory

	results, err := db.Query(`SELECT id, user_id,
		category_id,
		category_name,
		sub_category_name,
		is_active,
		created_at,
		updated_at FROM service_sub_categories where user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var subcategory SubCategory
		err = results.Scan(&subcategory.ID, &subcategory.UserID, &subcategory.CategoryID, &subcategory.CategoryName, &subcategory.SubCategoryName, &subcategory.IsActive, &subcategory.CreatedAt, &subcategory.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		subcategories = append(subcategories, subcategory)
	}
	return subcategories
}
func GetAllSubCategoriesPerUserPerGroup(userID int) []SubCategory {
	db := mysqldb.Connect()
	defer db.Close()

	var subcategories []SubCategory

	results, err := db.Query(`SELECT id, user_id,
		category_id,
		category_name,
		sub_category_name,
		is_active,
		created_at,
		updated_at FROM service_sub_categories where user_id = ? group by category_id`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var subcategory SubCategory
		err = results.Scan(&subcategory.ID, &subcategory.UserID, &subcategory.CategoryID, &subcategory.CategoryName, &subcategory.SubCategoryName, &subcategory.IsActive, &subcategory.CreatedAt, &subcategory.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		subcategories = append(subcategories, subcategory)
	}
	return subcategories
}

func AllSubCategoriesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	category, err := strconv.Atoi(vars["catid"])
	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	subcategories := GetAllSubCategoriesPerUser(userID)

	var subcategoriesPerCategory []SubCategory

	for _, s := range subcategories {
		cateID, _ := strconv.Atoi(s.CategoryID)
		if cateID == category {
			subcategoriesPerCategory = append(subcategoriesPerCategory, s)
		}
	}

	err = json.NewEncoder(w).Encode(SubCategoryConfig{AllSubCategories: subcategoriesPerCategory})
	if err != nil {
		println(err.Error())
	}
}
func AllSubCategoriesJSONAnaUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	subcategories := GetAllSubCategoriesPerUser(userID)

	err = json.NewEncoder(w).Encode(SubCategoryConfig{AllSubCategories: subcategories})
	if err != nil {
		println(err.Error())
	}
}

//UpdateSubCategoryJSON update the shop to the database using json
func UpdateSubCategoryJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var subcategory SubCategory

	err := decoder.Decode(&subcategory)

	if err != nil {
		print(err.Error())
	}

	vars := mux.Vars(r)

	subcategoryID := vars["id"]
	ID, err := strconv.Atoi(subcategoryID)

	if err != nil {
		print(err.Error())
	}

	if rowExists("SELECT id FROM service_sub_categories WHERE category_name= ? and sub_category_name= ? and user_id = ?", subcategory.CategoryName, subcategory.SubCategoryName, subcategory.UserID) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: "Sub Category exists"})
		return
	}

	stmt, err := db.Prepare(`UPDATE service_sub_categories SET
			user_id = ?,
			category_id = ?,
			category_name = ?,
			sub_category_name = ?,
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(subcategory.UserID, subcategory.CategoryID, subcategory.CategoryName, subcategory.SubCategoryName, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update sub category in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update sub category in the Database"})
}

//DisableSubCategory update the shop if its active or not
func DisableSubCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var subcategory SubCategory
	err := decoder.Decode(&subcategory)

	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	subcategoryID := vars["id"]
	ID, err := strconv.Atoi(subcategoryID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE service_sub_categories SET is_active = ? WHERE id = ?")
	_, err = stmt.Exec(subcategory.IsActive, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update sub category in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update sub category in the Database"})
}

//DeleteSubCategory delete a service
func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	subcategoryID := vars["id"]

	ID, _ := strconv.Atoi(subcategoryID)

	stmt, err := db.Prepare("DELETE FROM service_sub_categories where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete sub category from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted  sub category from database "})
}
