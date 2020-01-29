package storesettings

import (
	"encoding/json"
	"fmt"
	tr "gettext"
	"log"
	"mysqldb"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Products struct {
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
	Product      []Product
	Suppliers    []Supplier
}

type Product struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	Stores            []string  `json:"stores"`
	ProductCode       string    `json:"product_code"`
	SupplierName      string    `json:"supplier_name"`
	SupplierID        int       `json:"supplier_id"`
	ProductName       string    `json:"product_name"`
	ProductML         int       `json:"product_ml"`
	ProductPrice      float64   `json:"product_price"`
	ProductMlPerPRice float32   `json:"product_ml_per_price"`
	ProductStock      float32   `json:"product_stock"`
	IsActive          bool      `json:"is_active"`
	CanSplit          bool      `json:"can_split"`
	StoreName         string    `json:"store_name"`
	StoreID           int       `json:"store_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ProductStores struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	StoreID   int       `json:"store_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//ProductConfig holds the data for json structure
type ProductConfig struct {
	AllProducts []Product `json:"data"`
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {

	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var product Product

	err := decoder.Decode(&product)
	if err != nil {
		print(err.Error())
	}

	stores := product.Stores

	for _, s := range stores {
		storeIDInt, _ := strconv.Atoi(s)
		if rowExists("SELECT id FROM products WHERE product_code = ? and store_id= ? AND user_id = ? AND product_ml= ?", product.ProductCode, storeIDInt, product.UserID, product.ProductML) {
			json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: tr.Translate("Product exists")})

			return
		}
		stmt, err := db.Prepare(`INSERT INTO products (product_code,
			user_id,
			supplier_id,
			store_id,
			supplier_name,
			product_name,
			qty,
			product_ml,
			product_price,
			product_ml_per_price,
			is_active_product,
			can_split_product,
			created_at,
			updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
		if err != nil {
			println(err.Error())
		}
		res, err := stmt.Exec(product.ProductCode, product.UserID, product.SupplierID, storeIDInt, product.SupplierName, product.ProductName, product.ProductStock, product.ProductML, product.ProductPrice, product.ProductMlPerPRice, product.IsActive, product.CanSplit)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to insert product into database")})
		}
		id, err := res.LastInsertId()
		if err != nil {
			println(err.Error(), id)
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to get last insert id")})
		}

	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Update product in the Database")})

}
func GetAllProductsPerUser(userID int) []Product {
	db := mysqldb.Connect()
	defer db.Close()

	var products []Product

	results, err := db.Query(`SELECT user_id,
		product_id,
		product_code,
		product_name,
		product_ml,
		product_price,
		product_ml_per_price,
		qty,
		supplier_id,
		supplier_name,
		is_active_product,
		can_split_product,
		company_id,
		company_name,
		created_at,
		updated_at FROM products_view where user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var product Product
		err = results.Scan(&product.UserID, &product.ID, &product.ProductCode, &product.ProductName, &product.ProductML, &product.ProductPrice, &product.ProductMlPerPRice, &product.ProductStock, &product.SupplierID, &product.SupplierName, &product.IsActive, &product.CanSplit, &product.StoreID, &product.StoreName, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		products = append(products, product)
	}
	return products
}
func GetAllProductsPerUserPerStoreJSON(w http.ResponseWriter, r *http.Request) {
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

	var products []Product

	results, err := db.Query(`SELECT user_id,
		product_id,
		product_code,
		product_name,
		product_ml,
		product_price,
		product_ml_per_price,
		qty,
		supplier_id,
		supplier_name,
		is_active_product,
		can_split_product,
		company_id,
		company_name,
		created_at,
		updated_at FROM products_view where user_id = ? AND company_id = ? AND is_active_product = 1 ORDER BY product_name ASC`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var product Product
		err = results.Scan(&product.UserID, &product.ID, &product.ProductCode, &product.ProductName, &product.ProductML, &product.ProductPrice, &product.ProductMlPerPRice, &product.ProductStock, &product.SupplierID, &product.SupplierName, &product.IsActive, &product.CanSplit, &product.StoreID, &product.StoreName, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		products = append(products, product)
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		println(err.Error())
	}

}
func AllProductsJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	products := GetAllProductsPerUser(userID)
	err = json.NewEncoder(w).Encode(ProductConfig{AllProducts: products})
	if err != nil {
		println(err.Error())
	}
}

//UpdateProductJSON update the supplier to the database using json
func UpdateProductJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var product Product
	err := decoder.Decode(&product)

	vars := mux.Vars(r)

	productID := vars["id"]
	ID, err := strconv.Atoi(productID)

	if err != nil {
		print(err.Error())
	}
	stmt, err := db.Prepare(`UPDATE products SET
			product_code = ?,
			user_id = ?,
			supplier_id = ?,
			supplier_name = ?,
			product_name = ?,
			qty = qty + ?,
			product_ml = ?,
			product_price = ?,
			product_ml_per_price = ?,
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(product.ProductCode, product.UserID, product.SupplierID, product.SupplierName, product.ProductName, product.ProductStock, product.ProductML, product.ProductPrice, product.ProductMlPerPRice, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update product in the Database")})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Update propduct in the Database")})
}

//DisableProduct update the supplier if its active or not
func DisableProduct(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var product Product
	err := decoder.Decode(&product)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	productID := vars["id"]
	ID, err := strconv.Atoi(productID)
	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE products SET is_active_product = ? WHERE store_id = ? AND id = ?")
	_, err = stmt.Exec(product.IsActive, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update product in the database")})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Updated product in the Database")})
}

//CanSPlitProduct update the supplier if its active or not
func CanSPlitProduct(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var product Product
	err := decoder.Decode(&product)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	productID := vars["id"]
	ID, err := strconv.Atoi(productID)
	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE products SET can_split_product = ? WHERE store_id = ? AND id = ?")
	_, err = stmt.Exec(product.CanSplit, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update product in the database")})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Updated product in the Database")})
}
func TotalProductQTY(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	db := mysqldb.Connect()
	defer db.Close()

	var product = Product{}
	db.QueryRow(`SELECT qty FROM products WHERE id  = ?`, productID).Scan(&product.ProductStock)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		println(err.Error())
	}
}
