package bookings

import (
	"encoding/json"
	tr "gettext"
	"mysqldb"
	"net/http"
	"os"
	"time"
)

type CreateFeatureRantevou struct {
	UserID                int    `json:"user_id"`
	RantevouStart         string `json:"rantevou_start"`
	RantevouEnd           string `json:"rantevou_end"`
	RantevouStoreID       int    `json:"rantevou_store_id"`
	RantevouHairdresserID int    `json:"rantevou_hairdresser_id"`
	RantevouCustomerID    int    `json:"select_customer_id"`
	AServiceID            int    `json:"service_id"`
	AllRantevouPrices     string `json:"rantevou_price"`
	Status                string `json:"status"`
	AllRantevouDurations  string `json:"rantevou_duration"`
	RantevouComments      string `json:"rantevou_comments"`
	HairdresserName       string `json:"hairdresser_name"`
	IsAllDay              bool   `json:"is_all_day"`
	PromotionID           int    `json:"promotion_id"`
}

func InsertFeatureRantevouJSON(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var rantevou CreateFeatureRantevou

	err := decoder.Decode(&rantevou)
	if err != nil {
		print(err.Error())
	}

	strt := rantevou.RantevouStart
	ettime := rantevou.RantevouEnd

	startTime, _ := time.Parse("02/01/2006 15:04", strt)
	endTIme, _ := time.Parse("02/01/2006 15:04", ettime)

	stmt, err := db.Prepare(`INSERT INTO appointments (user_id,
		hairdresser_id,
		customer_id,
		store_id,
		service_id,
		start_time,
		end_time,
		service_status,
		comments,
		is_all_day,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?,?,?,?,?, NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}
	res, err := stmt.Exec(rantevou.UserID, rantevou.RantevouHairdresserID, rantevou.RantevouCustomerID, rantevou.RantevouStoreID, rantevou.AServiceID, startTime, endTIme, rantevou.Status, rantevou.RantevouComments, rantevou.IsAllDay)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert feature appointments into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}
	err = insertAppointmentPromotions(rantevou.UserID, rantevou.RantevouStoreID, rantevou.PromotionID, int(id), rantevou.RantevouHairdresserID, rantevou.RantevouCustomerID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed insert appointment with promotion to database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted feature appointment Into the Database"})
}

func insertAppointmentPromotions(userID, storeID, promotionID, appointmentID, hairdresserID, customerID int) error {
	db := mysqldb.Connect()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO appointments_with_promotions (user_id,
		store_id,
		hairdresser_id,
		customer_id,
		promotion_id,
		appointment_id,
		created_at,
		updated_at) VALUES (?,?,?,?,?,?, NOW(), NOW())`)
	if err != nil {
		println(err.Error())
	}
	res, err := stmt.Exec(userID, storeID, hairdresserID, customerID, promotionID, appointmentID)

	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
