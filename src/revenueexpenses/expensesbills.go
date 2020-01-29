package revenueexpenses

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ExpensesTypes struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type ExpensesTypesConfig struct {
	AllExpensesTypes []ExpensesTypes `json:"data"`
}

func TotalExpensesData(userID, storeID int, from, to time.Time) []ExpensesTypes {
	db := mysqldb.Connect()
	defer db.Close()

	var totalexpenses []ExpensesTypes

	results, err := db.Query(`SELECT expenses_list.expenses_name,
		SUM(expensestrans.expenses_price) AS total_expences
		FROM  expensestrans 
		LEFT JOIN expenses_list ON expensestrans.expenses_list_id = expenses_list.id
		WHERE expensestrans.user_id = ? AND  expensestrans.store_id = ? AND DATE(expensestrans.created_at) BETWEEN DATE(?) AND (?)
		GROUP BY expensestrans.expenses_list_id `, userID, storeID, from, to)

	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {

		var totals ExpensesTypes
		err = results.Scan(&totals.Name, &totals.Total)

		if err != nil {
			println(err.Error())
		}

		totalexpenses = append(totalexpenses, totals)
	}

	return totalexpenses
}

func TotalExpensesDataJSON(w http.ResponseWriter, r *http.Request) {
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

	totalexpenses := TotalExpensesData(userID, storeID, fromt, tot)

	json.NewEncoder(w).Encode(ExpensesTypesConfig{AllExpensesTypes: totalexpenses})
}
