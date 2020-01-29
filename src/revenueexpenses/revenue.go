package revenueexpenses

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"storesettings"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type RevenueReports struct {
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
	Stores       []storesettings.Shop
}

type RevenueTypes struct {
	TotalRevenue float64 `json:"total_revenue"`
	IsService    bool    `json:"is_service"`
}
type TotalIncomePerService struct {
	IsService bool      `json:"is_service"`
	SingleDay time.Time `json:"single_day"`
	Total     float64   `json:"total"`
}

type Charts struct {
	Type  string `json:"type"`
	Datas `json:"data"`
}
type Datas struct {
	Labels   []time.Time `json:"labels"`
	Datasets []Datasets  `json:"datasets"`
}

type Datasets struct {
	Label           bool   `json:"label"`
	BackgroundColor string `json:"backgroundColor"`

	Data []float64 `json:"data"`
}

type ResponseTime struct {
	UserID       int       `json:"user_id"`
	ContactID    string    `json:"nodeID"`
	ResponseTime float64   `json:"responseTime"`
	NodeName     string    `json:"node_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetAllSalesPerServiceProductPerDateJSONGraph(w http.ResponseWriter, r *http.Request) {
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

	var primes = []string{"true", "false"}

	// start := time.Now(

	var totalIncomePerService [][]TotalIncomePerService

	for _, symbol := range primes {

		totalIncomePerService = append(totalIncomePerService, GetAllSalesPerServiceProductPerDayGraph(userID, storeID, symbol, fromt, tot))

	}

	tableCharts := Datas{}.Datasets
	labels := Datas{}.Labels
	var employeeNames []bool
	// for _, t := range totalSales {
	// 	if t.IsService == true {
	// 		service := "Services"
	// 		employeeNames = append(employeeNames,service)
	// 	} else {
	// 		product := "Products"
	// 		employeeNames = append(employeeNames,service)
	// 	}
	// }

	for _, s := range totalIncomePerService {
		for _, si := range s {
			labels = append(labels, si.SingleDay)

			employeeNames = append(employeeNames, si.IsService)
		}

	}

	employeeName := removeDuplicatesUnordered(employeeNames)
	colors := []string{"#26B99A", "#03586A"}

	for i, n := range employeeName {
		totals := []float64{}

		for _, t := range totalIncomePerService {
			for _, ti := range t {
				if n == ti.IsService {

					totals = append(totals, ti.Total)

				}
			}

		}

		t := Datasets{
			Label:           n,
			BackgroundColor: colors[i],
			Data:            totals,
		}

		tableCharts = append(tableCharts, t)

	}

	c := Charts{
		Type: "line",
		Datas: Datas{
			Labels: labels,

			Datasets: tableCharts,
		},
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		println(err.Error())
	}
}

func CallRevenueData(userID, storeID int, from, to time.Time) []RevenueTypes {
	db := mysqldb.Connect()
	defer db.Close()

	var totalrevenues []RevenueTypes

	results, err := db.Query(`select sum(service_line_total) as total_sum,is_service
	from salestrans
	where user_id = ? and store_id = ? and date(created_at) between date(?) and date(?)
	group by is_service `, userID, storeID, from, to)

	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {

		var totals RevenueTypes
		err = results.Scan(&totals.TotalRevenue, &totals.IsService)

		if err != nil {
			println(err.Error())
		}

		totalrevenues = append(totalrevenues, totals)
	}

	return totalrevenues
}

func CallRevenueDataJSON(w http.ResponseWriter, r *http.Request) {
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

	totalrevenues := CallRevenueData(userID, storeID, fromt, tot)

	json.NewEncoder(w).Encode(totalrevenues)
}

func GetAllSalesPerServiceProductPerDayGraph(userID, storeID int, isservice string, from, to time.Time) []TotalIncomePerService {
	services := 0
	if isservice == "false" {
		services = 0
	} else if isservice == "true" {
		services = 1
	}
	db := mysqldb.Connect()
	defer db.Close()

	var totals []TotalIncomePerService

	results, err := db.Query(`select  DATE(d.wdate) AS created_at, d.wval + ifnull(x.sum_total,0) AS sum_total
	from daily_values d
	left outer join
		(
		SELECT  DATE(created_at) AS created_as,
			SUM(service_line_total) AS sum_total
			
		FROM salestrans
		WHERE is_service =? AND user_id =? and store_id = ? and DATE(created_at)  
		  between date(?) and date(?)
		   GROUP BY DATE(created_at)
		) x on d.wdate = x.created_as
	where d.wdate between date(?) and date(?)
	ORDER BY DATE(d.wdate)`, services, userID, storeID, from, to, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total TotalIncomePerService

		err = results.Scan(&total.SingleDay, &total.Total)
		if err != nil {
			println(err.Error())
		}
		if isservice == "false" {
			total.IsService = false
		}
		if isservice == "true" {
			total.IsService = true
		}
		totals = append(totals, total)
	}
	return totals
}
func removeDuplicatesUnordered(elements []bool) []bool {
	encountered := map[bool]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []bool{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}
