package main

import (
	"encoding/json"
	tr "gettext"
	"mysqldb"
	"net/http"
	"os"
)

type UserPackage struct {
	UserID  int     `json:"user_id"`
	Months  int     `json:"month_periods"`
	Stores  int     `json:"number_of_stores"`
	Amount  float64 `json:"amount"`
	Credits int     `json:"credits"`
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {

	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var userp UserPackage

	err := decoder.Decode(&userp)
	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`INSERT INTO users_packages (user_id,
			month_periods,
			number_of_stores,
			amount,
			credits,
			created_at,
			updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}

	res, err := stmt.Exec(userp.UserID, userp.Months, userp.Stores, userp.Amount, userp.Credits)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to insert package into database")})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error(), id)
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to get last insert id")})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Update package in the Database")})

}
