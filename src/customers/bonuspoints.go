package customers

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateCustomersPoints(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	customerID, _ := strconv.Atoi(vars["id"])

	decoder := json.NewDecoder(r.Body)
	var singleCustomer Customer

	err := decoder.Decode(&singleCustomer)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Cannot decode message body "})

	}

	stmt, err := db.Prepare("UPDATE customers SET customer_points = customer_points + ? WHERE id = ?")
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Cannot update customer points " + err.Error()})
	}
	_, err = stmt.Exec(singleCustomer.CustomerPoints, customerID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Cannot update customer points " + err.Error()})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "You successfully updated customer points: "})
}
