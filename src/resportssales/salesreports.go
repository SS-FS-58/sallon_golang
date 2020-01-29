package resportssales

import (
	"encoding/json"
	"mysqldb"
	"net/http"
	"storesettings"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type SalesReports struct {
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

type TotalSalesPerEmployee struct {
	HairdresserID         int     `json:"hairdresser_id'`
	HairdresserName       string  `json:"hairdresser_name"`
	TotalQtyServices      int     `json:"total_qty_services"`
	TotalpPriceServices   float64 `json:"total_price_services"`
	TotalServiceDiscounts float64 `json:"total_service_discount"`
	TotalServiceDuration  int     `json:"total_service_duration"`
}
type SalesServicesPerEmployeeConfig struct {
	AllSalesServicesPerEmployee []TotalSalesPerEmployee `json:"data"`
}

type TotalIncomePerEmployee struct {
	HairdresserName string    `json:"hairdresser_name"`
	HairdresserID   int       `json:"hairdresser_int"`
	SingleDay       time.Time `json:"single_day"`
	Total           float64   `json:"total"`
}

type SalesGraphForEmployess struct {
	AllNames             []string    `json:"all_names"`
	AllPricesPerEmployee [][]float64 `json:"all_prices_per_employee"`
	AllDates             []string    `json:"all_dates"`
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
	Label                     string    `json:"label"`
	BackgroundColor           string    `json:"backgroundColor"`
	BorderColor               string    `json:"borderColor"`
	PointBorderColor          string    `json:"pointBorderColor"`
	PointBackgroundColor      string    `json:"pointBackgroundColor"`
	PointHoverBackgroundColor string    `json:"pointHoverBackgroundColor"`
	PointHoverBorderColor     string    `json:"pointHoverBorderColor"`
	PointBorderWidth          int       `json:"pointBorderWidth"`
	FillColor                 string    `json:"fillColor"`
	StrokeColor               string    `json:"strokeColor"`
	PointColor                string    `json:"pointColor"`
	PointStrokeColor          string    `json:"pointStrokeColor"`
	PointHighlightFill        string    `json:"pointHighlightFill"`
	PointHighlightStroke      string    `json:"pointHighlightStroke"`
	Data                      []float64 `json:"data"`
}

type ResponseTime struct {
	UserID       int       `json:"user_id"`
	ContactID    string    `json:"nodeID"`
	ResponseTime float64   `json:"responseTime"`
	NodeName     string    `json:"node_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetAlServiceslSalesPerEmployeeJSON(w http.ResponseWriter, r *http.Request) {
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

	totalSales := GetAlServiceslSalesPerEmployee(userID, storeID, fromt, tot)
	err = json.NewEncoder(w).Encode(SalesServicesPerEmployeeConfig{AllSalesServicesPerEmployee: totalSales})
	if err != nil {
		println(err.Error())
	}
}

func GetAlServiceslSalesPerEmployeeJSONGraph(w http.ResponseWriter, r *http.Request) {
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

	totalSales := GetAlServiceslSalesPerEmployeeGraph(userID, storeID, fromt, tot)
	tableCharts := Datas{}.Datasets
	labels := Datas{}.Labels

	var employeeNames []string
	for _, s := range totalSales {
		labels = append(labels, s.SingleDay)

		employeeNames = append(employeeNames, s.HairdresserName)

	}

	employeeName := removeDuplicatesUnordered(employeeNames)
	hairdressers := storesettings.GetAllHairdressersPerUser(userID)
	colors := []string{}
	for _, h := range hairdressers {
		for _, e := range employeeName {
			if h.HairdresserName == e {
				colors = append(colors, h.Color)
			}
		}
	}

	for i, n := range employeeName {
		totals := []float64{}

		for _, t := range totalSales {
			if n == t.HairdresserName {

				totals = append(totals, t.Total)

			}
		}

		t := Datasets{
			Label:                     n,
			BackgroundColor:           colors[i],
			BorderColor:               colors[i],
			PointBorderColor:          colors[i],
			PointBackgroundColor:      colors[i],
			PointHoverBackgroundColor: colors[i],
			PointHoverBorderColor:     colors[i],
			PointBorderWidth:          1,
			FillColor:                 colors[i],
			Data:                      totals,
		}

		tableCharts = append(tableCharts, t)

	}

	c := Charts{
		Type: "line",
		Datas: Datas{
			Labels:   labels,
			Datasets: tableCharts,
		},
	}

	// fmt.Println(statistics)
	// res2B, _ := json.Marshal(statistics)
	// fmt.Println("\n", string(res2B))
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		println(err.Error())
	}
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}
func GetAlServiceslSalesPerEmployee(userID, storeID int, from, to time.Time) []TotalSalesPerEmployee {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []TotalSalesPerEmployee

	results, err := db.Query(`SELECT hairdresser_id,
		hairdresser_name,
		SUM(service_qty),
		SUM(service_line_total),
		SUM(service_discount),
		SUM(service_duration)
 	FROM sales_report_for_services 
 	WHERE user_id = ? AND store_id = ? AND DATE(created_at)  
	between date(?) and date(?)
 	GROUP BY hairdresser_id`, userID, storeID, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total TotalSalesPerEmployee
		err = results.Scan(&total.HairdresserID, &total.HairdresserName, &total.TotalQtyServices, &total.TotalpPriceServices, &total.TotalServiceDiscounts, &total.TotalServiceDuration)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}

func GetAlServiceslSalesPerEmployeeGraph(userID, storeID int, from, to time.Time) []TotalIncomePerEmployee {
	db := mysqldb.Connect()
	defer db.Close()

	var totals []TotalIncomePerEmployee

	results, err := db.Query(`select 
		DISTINCT DATE(w.wdate) AS sorted_date,
		w.wval + ifnull(x.total_value,0) AS newval,
	   w.hairdresser_id,
	   w.hairdresser_name
	from
	   (
		   select 
				   d.wdate,
				   d.wval,
				   p.hairdresser_id,
				   p.hairdresser_name
		   from daily_values d
		   cross join sales_report_for_services p 
	   ) w
	   left join
	   (
		   SELECT  s.hairdresser_id,
				  s.hairdresser_name,
				  DATE(s.created_at) as doc_date,
				   SUM(s.service_line_total) total_value
		   FROM    sales_report_for_services s
				   LEFT JOIN  hairdressers h ON s.hairdresser_id  = h.id
				   WHERE s.user_id = ? AND s.store_id = ? AND DATE(s.created_at)  
			  between date(?) and date(?)
		   GROUP   BY s.hairdresser_id,
				  s.hairdresser_name,DATE(s.created_at)
	   ) x on w.wdate = x.doc_date and w.hairdresser_id = x.hairdresser_id
	where w.wdate between date(?) and date(?)

	ORDER BY sorted_date`, userID, storeID, from, to, from, to)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var total TotalIncomePerEmployee
		err = results.Scan(&total.SingleDay, &total.Total, &total.HairdresserID, &total.HairdresserName)
		if err != nil {
			println(err.Error())
		}
		totals = append(totals, total)
	}
	return totals
}
