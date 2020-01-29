package resportssales

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ProductSales struct {
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Total float64 `json:"total"`
}

type ProductSalesPerEmployeeConfig struct {
	AllSProductsalesPerEmployee []ProductSales `json:"data"`
}

func GetAlProductsSalesPerEmployeeJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	from, _ := strconv.ParseInt(vars["from"], 10, 64)
	fromt := time.Unix(from, 0)
	to, _ := strconv.ParseInt(vars["to"], 10, 64)
	tot := time.Unix(to, 0)

	totalSales := GetAlProductslSalesPerEmployee(userID, storeID, fromt, tot)
	err = json.NewEncoder(w).Encode(ProductSalesPerEmployeeConfig{AllSProductsalesPerEmployee: totalSales})
	if err != nil {
		println(err.Error())
	}
}

func GetAlProductslSalesPerEmployee(userID, storeID int, from, to time.Time) []ProductSales {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []ProductSales

	results, err := db.Query(`SELECT 
		products.product_name,
		SUM(salestrans.service_line_total) AS total_price,
		SUM(salestrans.service_qty) AS total_qty
	
	FROM salestrans
	LEFT JOIN products ON salestrans.service_id = products.id AND salestrans.is_service = false
	WHERE salestrans.is_service = false and salestrans.user_id = ? and salestrans.store_id = ? and DATE(salestrans.created_at)  
	   between date(?) and date(?)
	GROUP BY products.product_name, DATE(salestrans.created_at)`, userID, storeID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total ProductSales
		err = results.Scan(&total.Name, &total.Total, &total.Qty)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}
