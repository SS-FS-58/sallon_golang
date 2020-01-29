package resportssales

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ServiceSales struct {
	Name                 string  `json:"name"`
	Qty                  int     `json:"qty"`
	TotalServiceDuration int     `json:"total_service_duration"`
	Total                float64 `json:"total"`
}

type ServiceSalesConfig struct {
	AllServicesales []ServiceSales `json:"data"`
}

func GetAlServicesSalesJSON(w http.ResponseWriter, r *http.Request) {
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

	totalSales := GetAllServicesSalesPerService(userID, storeID, fromt, tot)
	err = json.NewEncoder(w).Encode(ServiceSalesConfig{AllServicesales: totalSales})
	if err != nil {
		println(err.Error())
	}
}

func GetAllServicesSalesPerService(userID, storeID int, from, to time.Time) []ServiceSales {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []ServiceSales

	results, err := db.Query(`SELECT 
		service_name,
		SUM(service_line_total) AS total_price,
		SUM(service_qty) AS total_qty,
		SUM(service_duration) AS total_service_duration
	
	FROM sales_report_for_services
	
	WHERE user_id = ? and store_id = ? and DATE(created_at)  
	   between date(?) and date(?)
	GROUP BY service_name`, userID, storeID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total ServiceSales
		err = results.Scan(&total.Name, &total.Total, &total.Qty, &total.TotalServiceDuration)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}
