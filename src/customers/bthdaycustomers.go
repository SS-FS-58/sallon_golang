package customers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type BthdayCustomerConfig struct {
	AllBthdayCustomers []Customer `json:"data"`
}

func AllBthadyCustomersJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		print(err.Error())

	}
	println(userID)

	from, _ := strconv.ParseInt(vars["bthday"], 10, 64)
	fromt := time.Unix(from, 0)

	customers := GetAllCustomersPerUser(userID)
	var bthdayCustomers []Customer
	for _, g := range customers {
		if g.DateOfBirth.Format("02-01") == fromt.Format("02-01") {
			bthdayCustomers = append(bthdayCustomers, g)
		}
	}
	err = json.NewEncoder(w).Encode(BthdayCustomerConfig{AllBthdayCustomers: bthdayCustomers})
	if err != nil {
		println(err.Error())
	}
}
