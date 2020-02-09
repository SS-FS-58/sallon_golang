package storesettings

import (
	"database/sql"
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

type Hairdressers struct {
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
	Hairdressers []Hairdresser
}

type Hairdresser struct {
	ID                      int       `json:"id"`
	UserID                  int       `json:"user_id"`
	Stores                  []string  `json:"stores"`
	HairdresserName         string    `json:"hairdresser_name"`
	HairdressersMobilePhone string    `json:"hairdresser_mobile_phone"`
	HairdressersPhone       string    `json:"hairdresser_phone"`
	DisplayOrder            int       `json:"display_order"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	StoreID                 int       `json:"-"`
	StoreName               string    `json:"store_name"`
}

type HairdresserView struct {
	UserID                  int       `json:"user_id"`
	HairdresserID           int       `json:"hairdresser_id"`
	HairdresserName         string    `json:"hairdresser_name"`
	HairdressersMobilePhone string    `json:"hairdresser_mobile_phone"`
	HairdressersPhone       string    `json:"hairdresser_phone"`
	DisplayOrder            int       `json:"display_order"`
	IsActive                bool      `json:"is_active"`
	StoreName               string    `json:"store_name"`
	StoreID                 int       `json:"store_id"`
	Color                   string    `json:"color"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
type HairdresserStores struct {
	ID            int       `json:"id"`
	HairdresserID int       `json:"hairdresser_id"`
	StoreID       int       `json:"store_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type HairdresserConfig struct {
	AllHairdresser []HairdresserView `json:"data"`
}

func CreateHairdresser(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)


	var hairdresser Hairdresser

	err := decoder.Decode(&hairdresser)
	if err != nil {
		print(err.Error())
	}
	stores := hairdresser.Stores
	var allHairdressers []string
	var allStoreNames []string
	var allstoreIDs []int
	var allstoreIDss []int
	var hairdresserID int
	for _, s := range stores {
		var hair Hairdresser
		sint, _ := strconv.Atoi(s)
		err = db.QueryRow(`SELECT first.id,
			first.hairdresser_name,
			 third.id,
			 third.company_name
			FROM hairdressers first
			 LEFT JOIN hairdresser_stores second on first.id = second.hairdresser_id 
			 INNER JOIN shops third on second.store_id = third.id 
			 WHERE first.user_id = ? AND first.hairdresser_name = ?`, hairdresser.UserID, hairdresser.HairdresserName).Scan(&hair.ID, &hair.HairdresserName, &hair.StoreID, &hair.StoreName)

		if err != nil && err != sql.ErrNoRows {
			log.Println(err.Error())
			println("-------------------Error Hairdresser ok error no!-------------------------------" + err.Error())
	

		}
		if err != sql.ErrNoRows {
			allHairdressers = append(allHairdressers, hair.HairdresserName)
			allstoreIDs = append(allstoreIDs, hair.StoreID)
			hairdresserID = hair.ID
			allStoreNames = append(allStoreNames, hair.StoreName)

		}
		if !rowExists("SELECT id FROM hairdresser_stores WHERE hairdresser_id = ? and store_id= ?", hair.ID, sint) {
			allstoreIDss = append(allstoreIDss, sint)
		}

	}

	if allHairdressers == nil {

		stmt, err := db.Prepare(`INSERT INTO hairdressers (user_id,
		hairdresser_name,
		hairdresser_mobile_phone,
		hairdresser_phone,
		display_order,
		created_at,
		updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
		if err != nil {
			println(err.Error())
			

		}
		res, err := stmt.Exec(hairdresser.UserID, hairdresser.HairdresserName, hairdresser.HairdressersMobilePhone, hairdresser.HairdressersPhone, hairdresser.DisplayOrder)

		if err != nil {
			println(err.Error())
			
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to insert hairdresser into database")})
		}
		id, err := res.LastInsertId()
		if err != nil {
			println(err.Error())
			
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to get last insert id")})
		}

		intID := strconv.Itoa(int(id))
		color := "#66615B"
		for _, store := range stores {
			wstmt, err := db.Prepare("INSERT INTO hairdresser_stores (hairdresser_id, store_id,is_active_hairdresser,color,created_at, updated_at) VALUES (?,?,1,?,NOW(),NOW())")
			if err != nil {
				println(err.Error())
			}
			j, err := strconv.Atoi(store)
			if err != nil {
				println(err.Error())
				
			}
			wres, err := wstmt.Exec(intID, j, color)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to insert into database")})
			}
			wid, err := wres.LastInsertId()
			if err != nil {
				println(err.Error(), wid)
				
			}

		}

		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully insert hairdresser to the database"), Body: strconv.Itoa(int(id))})
	} else {
		if allstoreIDss == nil {
			json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: tr.Translate("Hairdresser exists")})
			return
		}
		color := "#66615B"
		for _, store := range allstoreIDss {

			wstmt, err := db.Prepare("INSERT INTO hairdresser_stores (hairdresser_id, store_id,is_active_hairdresser,color,created_at, updated_at) VALUES (?,?,1,?,NOW(),NOW())")
			if err != nil {
				println(err.Error())
			}

			wres, err := wstmt.Exec(hairdresserID, store, color)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				println(err.Error())
				json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to insert into database")})
			}
			wid, err := wres.LastInsertId()
			if err != nil {
				println(err.Error(), wid)
			}

		}
		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully insert hairdresser to the database")})

	}

}
func GetAllHairdressersPerUser(userID int) []HairdresserView {
	db := mysqldb.Connect()
	defer db.Close()

	var hairdressers []HairdresserView

	results, err := db.Query(`SELECT user_id,hairdresser_id,hairdresser_name,
		hairdresser_mobile_phone,
		hairdresser_phone,
		display_order,
		is_active_hairdresser,
		company_name,
		company_id,
		color,
		created_at,
		updated_at 
		FROM hairdresser_view
		WHERE user_id = ? AND is_active_hairdresser = ?`, userID, 1)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var hairdresser HairdresserView
		err = results.Scan(&hairdresser.UserID, &hairdresser.HairdresserID, &hairdresser.HairdresserName, &hairdresser.HairdressersMobilePhone, &hairdresser.HairdressersPhone, &hairdresser.DisplayOrder, &hairdresser.IsActive, &hairdresser.StoreName, &hairdresser.StoreID, &hairdresser.Color, &hairdresser.CreatedAt, &hairdresser.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		hairdressers = append(hairdressers, hairdresser)
	}
	return hairdressers
}
func AllHairdressersJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	hairdressers := GetAllHairdressersPerUser(userID)
	err = json.NewEncoder(w).Encode(HairdresserConfig{AllHairdresser: hairdressers})
	if err != nil {
		println(err.Error())
	}
}
func UpdateHairdresserJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var hairdresser Hairdresser
	err := decoder.Decode(&hairdresser)

	vars := mux.Vars(r)

	hairdressertID := vars["id"]
	ID, err := strconv.Atoi(hairdressertID)

	if err != nil {
		print(err.Error())
	}
	if rowExists("SELECT id FROM hairdressers WHERE hairdresser_name = ? and user_id= ?", hairdresser.HairdresserName, hairdresser.UserID) {
		json.NewEncoder(w).Encode(HTTPResp{Status: 501, Description: tr.Translate("Hairdresser exists")})
		return
	}
	stmt, err := db.Prepare(`UPDATE hairdressers SET
			user_id = ?,
			hairdresser_name = ?,
			hairdresser_mobile_phone = ?,
			hairdresser_phone = ?,
			display_order = ?, 
			updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(hairdresser.UserID, hairdresser.HairdresserName, hairdresser.HairdressersMobilePhone, hairdresser.HairdressersPhone, hairdresser.DisplayOrder, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update hairdressers in the Database")})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Update hairdressers in the Database")})
}
func DisableHairdresser(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)




	var hairdresser HairdresserView
	err := decoder.Decode(&hairdresser)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	hairdresserID := vars["id"]
	ID, _ := strconv.Atoi(hairdresserID)
	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	
	
	
	if err != nil {
		println(err.Error())
	}
	// hairdresser.IsActive = tinyint(0)
	stmt, _ := db.Prepare("UPDATE hairdresser_stores SET is_active_hairdresser = ? WHERE store_id = ? AND hairdresser_id = ?")
	_, err = stmt.Exec(0, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update hairdresser in the database")})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Updated hairdresser in the Database")})
}

func ChangeHairdresserColor(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var hairdresser HairdresserView
	err := decoder.Decode(&hairdresser)
	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	hairdresserID := vars["id"]
	ID, _ := strconv.Atoi(hairdresserID)
	storeID := vars["storeid"]
	StoreID, _ := strconv.Atoi(storeID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE hairdresser_stores SET color = ? WHERE store_id = ? AND hairdresser_id = ?")
	_, err = stmt.Exec(hairdresser.Color, StoreID, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: tr.Translate("Failed to update hairdresser in the database")})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: tr.Translate("Successfully Updated hairdresser in the Database")})
}
