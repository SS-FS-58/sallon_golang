package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CountPendingApp struct {
	Total int `json:"total"`
}
type TotalPerMonth struct {
	Month int     `json:"month"`
	Total float64 `json:"total"`
}

func CountAllPendingAppointment(userID, storeID int, today time.Time) int {

	db := mysqldb.Connect()
	defer db.Close()
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM appointments where user_id = ? AND store_id = ? AND service_status = "pending" AND DATE(start_time)  = ?`, userID, storeID, today.Format("2006-01-02")).Scan(&count)

	if err != nil {
		log.Println(err.Error())
	}

	return count
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return count
}
func CountAllPendingAppointmentJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	today, _ := strconv.ParseInt(vars["today"], 10, 64)
	todayt := time.Unix(today, 0)

	countPendings := CountAllPendingAppointment(userID, storeID, todayt)

	err = json.NewEncoder(w).Encode(countPendings)
	if err != nil {
		println(err.Error())
	}
}

func CountAllCancelledAppointment(userID, storeID int, today time.Time) int {

	db := mysqldb.Connect()
	defer db.Close()

	var totalpendings int
	err := db.QueryRow(`SELECT COUNT(*) FROM appointments 
						WHERE user_id = ? AND store_id = ? AND service_status = "cancelled" AND DATE(start_time) = ?`, userID, storeID, today.Format("2006-01-02")).Scan(&totalpendings)
	if err != nil {
		log.Println(err.Error())
	}
	return totalpendings
}
func CountAllCancelledAppointmentJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	today, _ := strconv.ParseInt(vars["today"], 10, 64)
	todayt := time.Unix(today, 0)

	countPendings := CountAllCancelledAppointment(userID, storeID, todayt)

	err = json.NewEncoder(w).Encode(countPendings)
	if err != nil {
		println(err.Error())
	}
}

func GetAllSalesPerMonthPerStore(userID, storeID int) []TotalPerMonth {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []TotalPerMonth

	results, err := db.Query(`select 
		DISTINCT month(w.wdate) AS sorted_date,
		w.wval + ifnull(x.total_value,0) AS newval
	   
	from
	   (
		   select 
				   date(d.wdate) as wdate,
				   d.wval
				  
		   from daily_values d
		   cross join salestrans p 
	   ) w
	   left join
	   (
		   SELECT  
				  
				  month(s.created_at) as doc_date,
				   SUM(s.service_line_total) total_value
		   FROM    salestrans s
				   
				   WHERE s.user_id = ? AND s.store_id = ? 
		   GROUP BY month(s.created_at)
	   ) x on month(w.wdate) = x.doc_date
       order by sorted_date`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total TotalPerMonth
		err = results.Scan(&total.Month, &total.Total)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}
func GetAllSalesPerMonthPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	totals := GetAllSalesPerMonthPerStore(userID, storeID)

	err = json.NewEncoder(w).Encode(totals)
	if err != nil {
		println(err.Error())
	}
}

func GetAllExpensesPerMonthPerStore(userID, storeID int) []TotalPerMonth {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []TotalPerMonth

	results, err := db.Query(`select 
		DISTINCT month(w.wdate) AS sorted_date,
		w.wval + ifnull(x.total_value,0) AS newval
	   
	from
	   (
		   select 
				   date(d.wdate) as wdate,
				   d.wval
				  
		   from daily_values d
		   cross join expensestrans p 
	   ) w
	   left join
	   (
		   SELECT  
				  
				  month(s.created_at) as doc_date,
				   SUM(s.expenses_price) total_value
		   FROM    expensestrans s
				   
				   WHERE s.user_id = ? AND s.store_id = ? 
		   GROUP BY month(s.created_at)
	   ) x on month(w.wdate) = x.doc_date
       order by sorted_date`, userID, storeID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total TotalPerMonth
		err = results.Scan(&total.Month, &total.Total)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}
func GetAllExpensesPerMonthPerStoreJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}

	storeID, err := strconv.Atoi(vars["storeid"])
	if err != nil {
		print(err.Error())

	}

	totals := GetAllExpensesPerMonthPerStore(userID, storeID)

	err = json.NewEncoder(w).Encode(totals)
	if err != nil {
		println(err.Error())
	}

}
