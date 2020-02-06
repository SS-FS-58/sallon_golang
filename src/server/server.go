package main

import (
	"bookings"
	"cancellation"
	"customers"
	"expenses"
	"fmt"
	tr "gettext"
	"html/template"
	"income"
	"log"
	"net/http"
	"os"
	"pendaingdays"
	"promotions"
	"resportssales"
	"revenueexpenses"
	"storesettings"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var templateCache = map[string]*template.Template{}

var templateFuncMap = template.FuncMap{
	"safehtml":    safeHTML,
	"touppercase": CompanyNameToUpperCase,
	"gettext":     tr.Translate,
	"add":         add,
}

func CacheMiddleware(h http.Handler, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int(duration.Seconds())))
		w.Header().Set("Expires", time.Now().Add(duration).Format(http.TimeFormat))
		h.ServeHTTP(w, r)
	})
}

func renderView(viewName string, w http.ResponseWriter, data interface{}) {
	if tmpl, ok := templateCache[viewName]; ok {
		tmpl.Execute(w, data)
		return
	}
	tmpl, err := template.New("base.html").Funcs(templateFuncMap).ParseFiles("views/base.html", "views/"+viewName)
	if err != nil {
		log.Fatalf("Error parsing templates %s: %v", viewName, err.Error())
	}
	templateCache[viewName] = tmpl
	err = tmpl.Execute(w, data)
	if err != nil {
		println(err.Error())
	}
}

func safeHTML(html string) template.HTML {
	return template.HTML(html)
}

func renderSubscriptionPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetSubscription()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("subscription.html", w, context)
}
func renderBillingPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetBilling()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("billing-history.html", w, context)
}
func renderCustomersPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCustomers()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	context.UserEmail = user.Email
	renderView("customers.html", w, context)
}
func renderViewCustomerPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetViewCustomer()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	singleCustomer := customers.SingleCustomerPerUser(ID, user.ID)
	context.MenuTitle = singleCustomer.CustomerName + " " + singleCustomer.CustomerSurname
	context.SingleCustomer = singleCustomer

	allAppointmentsPerCustomer := bookings.AllCommentsPerCustomer(ID)
	var lenCompleted []string
	var lenCanceled []string
	var lenPending []string

	for _, a := range allAppointmentsPerCustomer {
		if a.ServiceStatus == "complete" {

			lenCompleted = append(lenCompleted, a.ServiceStatus)
		}
		if a.ServiceStatus == "cancelled" {
			lenCanceled = append(lenCanceled, a.ServiceStatus)
		}
		if a.ServiceStatus == "pending" {
			lenPending = append(lenPending, a.ServiceStatus)
		}
	}
	context.LenCompletedServices = len(lenCompleted)
	context.LenCancelled = len(lenCanceled)
	context.LenPending = len(lenPending)
	var totalSalesServices float64
	var totalSalesProducts float64
	var totalDiscount float64
	allSalesPerCustomer := bookings.AllSalesTransPerCustomer(ID)

	for _, s := range allSalesPerCustomer {
		totalDiscount += s.ServiceDiscount
		if s.IsService == true {
			totalSalesServices += s.ServiceLineTotal
		}
		if s.IsService == false {
			totalSalesProducts += s.ServiceLineTotal
		}
	}
	context.TotalServices = totalSalesServices
	context.TotalProducts = totalSalesProducts
	context.TotalDiscount = totalDiscount
	context.Total = (totalSalesServices + totalSalesProducts)

	context.SingleCustomer = singleCustomer
	renderView("view-customer.html", w, context)
}
func renderCashRegisterPage(w http.ResponseWriter, r *http.Request) {

	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCashRegister()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("cashregister.html", w, context)

}
func renderAccountingPage(w http.ResponseWriter, r *http.Request) {

	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetAccounting()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("accounting.html", w, context)

}
func renderCampaignPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCampaign()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("campaign.html", w, context)
}
func renderReportsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetReports()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("reports.html", w, context)
}
func renderReportSalesPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetReportSales()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	allShops := storesettings.GetAllShopsPerUser(user.ID)
	context.Stores = allShops
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("report-sales.html", w, context)
}
func renderUnfinishedCanceledPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetUnfinishedCanceled()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	context.UserEmail = user.Email
	allShops := storesettings.GetAllShopsPerUser(user.ID)
	context.Stores = allShops
	renderView("unfinished-canceled.html", w, context)
}
func renderReportRevenueExpensesPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetReportRevenueExpenses()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	allShops := storesettings.GetAllShopsPerUser(user.ID)
	context.Stores = allShops
	renderView("report-revenue-expenses.html", w, context)
}
func renderReportPendingDaysPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetReportPendingDays()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	context.UserEmail = user.Email
	renderView("report-pending-days.html", w, context)
}
func renderReportCustomersWithBalancePage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetReportCustomersWithBalance()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("report-customers-with-balance.html", w, context)
}
func renderStoreSettingsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	context := GetStoreSettings()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("store-settings.html", w, context)
}
func renderStoreSettingsSmsEmailPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	context := GetStoreSettingsSmsEmail()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("sms-email-settings.html", w, context)
}
func renderShopsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetShops()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}
	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("shops.html", w, context)
}
func renderHairdressersPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetHairdressers()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}
	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("hairdressers.html", w, context)
}
func renderStoreProductsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetStoreProducts()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	var allActiveSuppliers []storesettings.Supplier
	suppliers := storesettings.GetAllSuppliersPerUser(user.ID)
	for _, s := range suppliers {
		if s.IsActive == true {
			allActiveSuppliers = append(allActiveSuppliers, s)
		}
	}
	context.Suppliers = allActiveSuppliers
	renderView("products.html", w, context)
}
func renderStoreSuppliersPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetStoreSuppliers()
	username := getUserName(r)
	println(username)
	user := QueryUser(username)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("suppliers.html", w, context)
}
func renderCampaignGooglePage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCampaignGoogle()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("campaigngoogle.html", w, context)
}
func renderCampaignFacebookPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCampaignFacebook()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("campaignfacebook.html", w, context)
}
func renderCampaignSmsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCampaignSms()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("campaignsms.html", w, context)
}
func renderStoreSettingsServicesPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetStoreSettingsServices()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID

	categories := storesettings.GetAllCategoriesPerUser(user.ID)
	context.LenCategories = len(categories)
	context.Categories = categories

	subCategories := storesettings.GetAllSubCategoriesPerUser(user.ID)
	context.SubCategories = subCategories
	renderView("services.html", w, context)
}
func renderCampaignEmailPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCampaignEmail()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("campaignemail.html", w, context)
}
func renderCalendarPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetCalendar()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	// allcustomers := customers.GetAllCustomersPerUser(user.ID)
	// context.Customers = allcustomers

	allShops := storesettings.GetAllShopsPerUser(user.ID)
	context.Stores = allShops

	// // get client ip address
	// ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	// // print out the ip address
	// fmt.Println(w, ip+"\n\n")

	// // sometimes, the user acccess the web server via a proxy or load balancer.
	// // The above IP address will be the IP address of the proxy or load balancer and not the user's machine.

	// // let's get the request HTTP header "X-Forwarded-For (XFF)"
	// // if the value returned is not null, then this is the real IP address of the user.
	// fmt.Println(w, "X-Forwarded-For :"+r.Header.Get("X-FORWARDED-FOR"))
	renderView("calendar.html", w, context)
}
func renderCheckoutCalendarPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		langParam := r.URL.Query().Get("lang")

		pathtolocale := os.Getenv("LOCALE_LANG")

		tr.Setup(langParam, langParam, pathtolocale)

		tr.ChangeLocale(langParam)
		context := GetCheckoutCalendar()
		username := getUserName(r)
		vatNumber := getVatNumberSession(r)
		if vatNumber != "" {
			context.VatNumber = vatNumber
		}

		user := QueryUser(username)
		context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
		context.ImageName = user.ID
		context.UserID = user.ID

		hairdresserID := r.FormValue("hairdresserID")
		hairdresserIDInt, _ := strconv.Atoi(hairdresserID)
		customerID := r.FormValue("customerID")
		eventDate := r.FormValue("eventDate")
		customerName := r.FormValue("customerName")
		customerEmail := r.FormValue("customerEmail")
		customerMobileNumber := r.FormValue("customerPhoneNumber")

		hairdreserName := r.FormValue("hairdreserName")
		storeID := r.FormValue("storeID")
		storeIDInt, _ := strconv.Atoi(storeID)
		customerIDInt, _ := strconv.Atoi(customerID)
		// serviceStatus := r.FormValue("serviceStatus")

		context.HairdresserName = hairdreserName
		context.CustomerName = customerName
		context.CustomerEmail = customerEmail
		context.CustomerPhone = customerMobileNumber
		eventDateTime, _ := time.Parse("2006-01-02T15:04:05", eventDate)
		context.AppointmentCloseDay = eventDateTime
		context.MenuTitle = hairdreserName

		allAppointmentsPerCustomerPerDay := bookings.AllAppointmentPerCustomerPerDay(user.ID, storeIDInt, customerIDInt, eventDate)
		services := storesettings.GetAllServicesPerUserPerStore(user.ID, storeIDInt)
		var allAppointmentsPercusPerDay []bookings.CustomerCheckout
		var total float64
		var totalDiscount float64
		for _, c := range allAppointmentsPerCustomerPerDay {
			for _, s := range services {
				if c.ServiceName == s.ServiceName {
					c.SwitchFormula = s.SwitchFormula
					totaPerRow := c.ServicePrice - c.ServiceDiscount
					c.TotalPerRow = totaPerRow
					totalDiscount += c.ServiceDiscount
					allAppointmentsPercusPerDay = append(allAppointmentsPercusPerDay, c)
					total += c.ServicePrice

				}
			}

		}
		context.CustomerCheckout = allAppointmentsPercusPerDay
		context.TotalDiscount = totalDiscount
		context.Total = total - totalDiscount

		customerPoints := customers.SingleCustomerPerUser(customerIDInt, user.ID)
		context.SingleCustomer = customerPoints
		context.StoreID = storeIDInt
		context.HairdresserID = hairdresserIDInt
		context.CustomerID = customerIDInt
		// context.EventID = eventID

		renderView("checkout-calendar.html", w, context)

	}

}

func renderPromotionsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetPromotions()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID
	renderView("promotions.html", w, context)
}
func renderRewardsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetRewards()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("rewards.html", w, context)
}
func renderRecommendationsPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetRecommendations()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	renderView("recommendations.html", w, context)
}
func notFound(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := Get404()

	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}
	context.Username = username
	if len(username) != 0 {
		renderView("404.html", w, context)
		return
	} else if len(username) < 1 {
		var loginTemplate = template.Must(template.New("4041.html").Funcs(templateFuncMap).ParseFiles("views/4041.html"))
		err := loginTemplate.Execute(w, context)

		if err != nil {
			println(err.Error())
		}
	}

}
func renderMyProfilePage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	context := GetMyProfile()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)

	context.ImageName = user.ID
	context.UserProfile = user.Username
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.User = user
	renderView("my-profile.html", w, context)
}
func renderExpensesPage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetExpenses()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID

	renderView("expenses.html", w, context)

}
func renderIncomePage(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	context := GetIncome()
	username := getUserName(r)
	vatNumber := getVatNumberSession(r)
	if vatNumber != "" {
		context.VatNumber = vatNumber
	}

	user := QueryUser(username)
	context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
	context.ImageName = user.ID
	context.UserID = user.ID

	renderView("income.html", w, context)

}

func main() {

	http.Handle("/static/", CacheMiddleware(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))), 7*24*time.Hour))

	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	router.Handle("/admin", AdminCheck(http.HandlerFunc(Admin), true))
	router.HandleFunc("/admin/login", login)
	router.HandleFunc("/register", renderRegister)
	router.HandleFunc("/admin/register", Register).Methods("POST")
	router.HandleFunc("/activate/{email}/{token}", packageRegister).Methods("GET")
	router.Handle("/admin/logout", AdminCheck(http.HandlerFunc(logout), true))
	router.HandleFunc("/admin/api/insert-user-package", CreateProducts).Methods("POST")

	router.Handle("/admin/subscription", AdminCheck(http.HandlerFunc(renderSubscriptionPage), true))
	router.Handle("/admin/customers", AdminCheck(http.HandlerFunc(renderCustomersPage), true))
	router.Handle("/admin/billing", AdminCheck(http.HandlerFunc(renderBillingPage), true))
	router.HandleFunc("/cashregister", renderCashRegisterPage)
	router.HandleFunc("/accounting", renderAccountingPage)
	router.HandleFunc("/campaign", renderCampaignPage)
	router.Handle("/admin/reports", AdminCheck(http.HandlerFunc(renderReportsPage), true))
	router.Handle("/admin/store-settings", AdminCheck(http.HandlerFunc(renderStoreSettingsPage), true))
	router.Handle("/admin/sms-email-settings", AdminCheck(http.HandlerFunc(renderStoreSettingsSmsEmailPage), true))
	router.Handle("/admin/stores", AdminCheck(http.HandlerFunc(renderShopsPage), true))
	router.Handle("/admin/hairdressers", AdminCheck(http.HandlerFunc(renderHairdressersPage), true))
	router.Handle("/admin/suppliers", AdminCheck(http.HandlerFunc(renderStoreSuppliersPage), true))
	router.Handle("/admin/products", AdminCheck(http.HandlerFunc(renderStoreProductsPage), true))
	router.Handle("/admin/services", AdminCheck(http.HandlerFunc(renderStoreSettingsServicesPage), true))

	router.HandleFunc("/admin/rewards", renderRewardsPage)
	router.Handle("/admin/promotions", AdminCheck(http.HandlerFunc(renderPromotionsPage), true))
	router.HandleFunc("/admin/recommendations", renderRecommendationsPage)

	router.Handle("/admin/api/all-stores-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllShopsJSONPerUser), true))

	//My Campaigns
	router.Handle("/admin/campaigngoogle", AdminCheck(http.HandlerFunc(renderCampaignGooglePage), true))
	router.Handle("/admin/campaignfacebook", AdminCheck(http.HandlerFunc(renderCampaignFacebookPage), true))
	router.Handle("/admin/campaignsms", AdminCheck(http.HandlerFunc(renderCampaignSmsPage), true))
	router.Handle("/admin/campaignemail", AdminCheck(http.HandlerFunc(renderCampaignEmailPage), true))

	//My Profile
	router.Handle("/admin/my-profile", AdminCheck(http.HandlerFunc(renderMyProfilePage), true))
	router.Handle("/admin/api/profile/{id}", AdminCheck(http.HandlerFunc(GetSingleUser), true))
	router.Handle("/admin/api/update-user/{id}", AdminCheck(http.HandlerFunc(UpdateSingleUser), true))

	//Store settings

	router.Handle("/admin/api/create-shop", AdminCheck(http.HandlerFunc(storesettings.CreateShops), true))
	router.Handle("/admin/api/change-password-for-shop/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateShopPassword), true))
	router.Handle("/admin/api/all-shops-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllShopsJSONPerUser), true))
	router.Handle("/admin/api/update-single-shop/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateShopJSON), true))
	router.Handle("/admin/api/disable-single-shop/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableShop), true))

	router.Handle("/admin/api/insert-wekkly-hours", AdminCheck(http.HandlerFunc(storesettings.InsertWeeklyHoursDataToDB), true))
	router.Handle("/admin/api/all-weekly-hours-per-user/{userid}", AdminCheck(http.HandlerFunc(storesettings.AllWeeklyHoursPerUserJSON), true))
	router.Handle("/admin/api/all-weekly-hours-per-user-per-store/{userid}/{storeid}", AdminCheck(http.HandlerFunc(storesettings.AllWeeklyHoursPerUserPerStoreJSON), true))

	router.Handle("/admin/api/delete-day-of-the-week/{id}", AdminCheck(http.HandlerFunc(storesettings.DeleteWeekDay), true))

	//Services settings
	router.Handle("/admin/api/create-service", AdminCheck(http.HandlerFunc(storesettings.CreateServices), true))
	router.Handle("/admin/api/all-services-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllServicesJSONPerUser), true))
	router.Handle("/admin/api/update-service/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateServiceJSON), true))
	router.Handle("/admin/api/enable-disable-service/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableService), true))
	router.Handle("/admin/api/enable-disable-has-formula/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.HasFormulaService), true))
	router.Handle("/admin/api/delete-single-service/{id}", AdminCheck(http.HandlerFunc(storesettings.DeleteService), true))

	//Service Categories settings
	router.Handle("/admin/api/create-category", AdminCheck(http.HandlerFunc(storesettings.CreateCategory), true))
	router.Handle("/admin/api/all-categories-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllCategoriesJSONPerUser), true))
	router.Handle("/admin/api/update-category/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateCategoryJSON), true))
	router.Handle("/admin/api/enable-disable-category/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableCategory), true))
	router.Handle("/admin/api/delete-single-category/{id}/{categoryname}/{userid}", AdminCheck(http.HandlerFunc(storesettings.DeleteCategory), true))

	//Service Sub Categories settings
	router.Handle("/admin/api/create-sub-category", AdminCheck(http.HandlerFunc(storesettings.CreateSubCategory), true))
	router.Handle("/admin/api/all-sub-categories-per-user/{id}/{catid}", AdminCheck(http.HandlerFunc(storesettings.AllSubCategoriesJSONPerUser), true))
	router.Handle("/admin/api/all-sub-categories/{id}", AdminCheck(http.HandlerFunc(storesettings.AllSubCategoriesJSONAnaUser), true))
	router.Handle("/admin/api/update-sub-category/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateSubCategoryJSON), true))
	router.Handle("/admin/api/enable-disable-sub-category/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableSubCategory), true))
	router.Handle("/admin/api/delete-single-sub-category/{id}", AdminCheck(http.HandlerFunc(storesettings.DeleteSubCategory), true))

	//Suppliers settings
	router.Handle("/admin/api/create-supplier", AdminCheck(http.HandlerFunc(storesettings.CreateSupplier), true))
	router.Handle("/admin/api/all-suppliers-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllSuppliersJSONPerUser), true))
	router.Handle("/admin/api/update-supplier/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateSupplierJSON), true))
	router.Handle("/admin/api/enable-disable-supplier/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableSupplier), true))
	router.Handle("/admin/api/check-vat-exists-for-supplier", AdminCheck(http.HandlerFunc(storesettings.CheckIfVatExistsForSuppliers), true))
	router.Handle("/admin/api/delete-single-supplier/{id}", AdminCheck(http.HandlerFunc(storesettings.DeleteSupplier), true))

	//Products settings
	router.Handle("/admin/api/create-product", AdminCheck(http.HandlerFunc(storesettings.CreateProducts), true))
	router.Handle("/admin/api/all-products-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllProductsJSONPerUser), true))
	router.Handle("/admin/api/update-product/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateProductJSON), true))
	router.Handle("/admin/api/enable-disable-product/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableProduct), true))
	router.Handle("/admin/api/enable-disable-can-split-product/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.CanSPlitProduct), true))
	router.Handle("/admin/api/all-products-per-user-per-store/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.GetAllProductsPerUserPerStoreJSON), true))
	router.Handle("/admin/api/product-qty/{id}", AdminCheck(http.HandlerFunc(storesettings.TotalProductQTY), true))

	//Hairdresser settings
	router.Handle("/admin/api/create-hairdresser", AdminCheck(http.HandlerFunc(storesettings.CreateHairdresser), true))
	router.Handle("/admin/api/all-hairdressers-per-user/{id}", AdminCheck(http.HandlerFunc(storesettings.AllHairdressersJSONPerUser), true))
	router.Handle("/admin/api/update-hairdresser/{id}", AdminCheck(http.HandlerFunc(storesettings.UpdateHairdresserJSON), true))
	router.Handle("/admin/api/enable-disable-hairdresser/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.DisableHairdresser), true))
	router.Handle("/admin/api/change-hairdresser-color/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.ChangeHairdresserColor), true))

	//Customer settings
	router.Handle("/admin/api/create-customer", AdminCheck(http.HandlerFunc(customers.CreateCustomer), true))
	router.Handle("/admin/api/all-customers-per-user/{id}", AdminCheck(http.HandlerFunc(customers.AllCustomersJSONPerUser), true))
	router.Handle("/admin/api/update-customer/{id}", AdminCheck(http.HandlerFunc(customers.UpdateCustomerJSON), true))
	router.Handle("/admin/view-customer/{id}/{name}", AdminCheck(http.HandlerFunc(renderViewCustomerPage), true))
	router.Handle("/admin/api/all-sales-per-customer/{customerid}", AdminCheck(http.HandlerFunc(bookings.AllSalesTransCustomerJSON), true))
	router.Handle("/admin/api/all-events-per-customer/{customerid}", AdminCheck(http.HandlerFunc(bookings.AllCalendarEventsPerCustomerJSON), true))
	router.Handle("/admin/api/update-comment/{id}", AdminCheck(http.HandlerFunc(bookings.UpdateComment), true))

	router.Handle("/admin/api/all-inactive-customers-per-user/{id}/{from}/{to}", AdminCheck(http.HandlerFunc(customers.AllInactiveCustomersJSONPerUser), true))
	router.Handle("/admin/api/all-bthday-customers-per-user/{id}/{bthday}", AdminCheck(http.HandlerFunc(customers.AllBthadyCustomersJSONPerUser), true))

	router.Handle("/admin/api/insert-bonus-points/{id}", AdminCheck(http.HandlerFunc(customers.UpdateCustomersPoints), true))

	// Expenses
	router.Handle("/admin/expenses", AdminCheck(http.HandlerFunc(renderExpensesPage), true))

	router.Handle("/admin/api/create-cost-category", AdminCheck(http.HandlerFunc(expenses.CreateCostCategory), true))
	router.Handle("/admin/api/all-cost-categories-per-user/{id}", AdminCheck(http.HandlerFunc(expenses.AllCostCategoriesJSONPerUser), true))
	router.Handle("/admin/api/update-cost-category/{id}", AdminCheck(http.HandlerFunc(expenses.UpdateCostCategoryJSON), true))
	router.Handle("/admin/api/enable-disable-cost-category/{id}", AdminCheck(http.HandlerFunc(expenses.DisableCostCategory), true))
	router.Handle("/admin/api/delete-single-cost-category/{id}", AdminCheck(http.HandlerFunc(expenses.DeleteCostCategory), true))

	router.Handle("/admin/api/create-expense", AdminCheck(http.HandlerFunc(expenses.InsertExpenseToDB), true))
	router.Handle("/admin/api/all-expenses-per-user/{id}", AdminCheck(http.HandlerFunc(expenses.AllExpensesJSONPerUser), true))
	router.Handle("/admin/api/paid-unpaid/{id}", AdminCheck(http.HandlerFunc(expenses.PaidUnpaid), true))
	router.Handle("/admin/api/delete-single-expense/{id}", AdminCheck(http.HandlerFunc(expenses.DeleteExpense), true))
	router.Handle("/admin/api/edit-expense/{id}", AdminCheck(http.HandlerFunc(expenses.UpdateExpenseJSON), true))

	// Income
	router.Handle("/admin/income", AdminCheck(http.HandlerFunc(renderIncomePage), true))
	router.Handle("/admin/api/insert-income-checkout-data", AdminCheck(http.HandlerFunc(income.InsertIncomeCheckoutDataToDB), true))

	// Calendar
	router.Handle("/admin/calendar", AdminCheck(http.HandlerFunc(renderCalendarPage), true))
	router.Handle("/admin/api/all-calendar-resources/{userid}/{storeid}", AdminCheck(http.HandlerFunc(bookings.AllCalendarResourcesJSONPerUser), true))
	router.Handle("/admin/api/all-services-per-user-per-store/{storeid}/{id}", AdminCheck(http.HandlerFunc(storesettings.GetAllServicesPerUserPerStoreJSON), true))
	router.Handle("/admin/api/create-rantevou", AdminCheck(http.HandlerFunc(bookings.InsertRantevouJSON), true))
	router.Handle("/admin/api/all-calendar-events/{userid}/{storeid}/{eventtime}/{endtime}", AdminCheck(http.HandlerFunc(bookings.AllCalendarEventsJSONPerUser), true))
	router.Handle("/admin/api/update-calendar", AdminCheck(http.HandlerFunc(bookings.UpdateCalendarDataJSONToDB), true))
	router.Handle("/admin/api/update-calendar-with-hairdresser-id", AdminCheck(http.HandlerFunc(bookings.UpdateCalendarDataJSONToDBWithHairdresserID), true))
	router.Handle("/admin/api/all-appointments-per-customer/{customerid}", AdminCheck(http.HandlerFunc(bookings.AllAppointmentPerCustomerJSON), true))
	router.Handle("/admin/api/create-feature-rantevou", AdminCheck(http.HandlerFunc(bookings.InsertFeatureRantevouJSON), true))

	router.Handle("/admin/api/all-bank-holidays-per-store/{storeid}", AdminCheck(http.HandlerFunc(storesettings.AllHolidaysJSONPerStore), true))

	router.Handle("/admin/api/cancelled-single-appointment", AdminCheck(http.HandlerFunc(bookings.CancelAppointmentJSON), true))

	//Checkout

	router.Handle("/admin/checkout", AdminCheck(http.HandlerFunc(renderCheckoutCalendarPage), true))
	router.Handle("/admin/api/insert-checkout-data", AdminCheck(http.HandlerFunc(bookings.InsertCheckoutDataToDB), true))
	router.Handle("/admin/api/insert-invoice-data", AdminCheck(http.HandlerFunc(bookings.InsertInvoiceJSON), true))
	router.Handle("/admin/api/all-invoices-per-store/{storeid}", AdminCheck(http.HandlerFunc(bookings.AllInvoicesPerCustomerJSON), true))

	//Promotions

	router.Handle("/admin/api/insert-promotion-data", AdminCheck(http.HandlerFunc(promotions.InsertPromotionDataToDB), true))
	router.Handle("/admin/api/all-promotion-services-per-user/{userid}", AdminCheck(http.HandlerFunc(promotions.AllPromotionServicesJSON), true))
	router.Handle("/admin/api/all-promotion-products-per-user/{userid}", AdminCheck(http.HandlerFunc(promotions.AllPromotionProductsJSON), true))
	router.Handle("/admin/api/update-promotion-data/{promid}", AdminCheck(http.HandlerFunc(promotions.UpdateSinglePromotionDataToDB), true))
	router.Handle("/admin/api/delete-promotion-data/{promid}", AdminCheck(http.HandlerFunc(promotions.Deletepromotion), true))
	router.Handle("/admin/api/all-promotion-services-per-service-per-user-per-store/{userid}/{storeid}/{serviceid}", AdminCheck(http.HandlerFunc(promotions.AllPromotionServicesPerUserPerStoreJSON), true))
	router.Handle("/admin/api/all-promotion-services-per-products-per-user-per-store/{userid}/{storeid}/{productid}", AdminCheck(http.HandlerFunc(promotions.AllPromotionProductsPerUserPerStoreJSON), true))

	// router.Handle("/admin/api/all-appointment-per-customer", AdminCheck(http.HandlerFunc(bookings.AllAppointmentPerCustomerPerDayJSON), true))

	// Report
	router.Handle("/admin/reportsales", AdminCheck(http.HandlerFunc(renderReportSalesPage), true))
	router.Handle("/admin/api/all-services-sales-per-employee/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(resportssales.GetAlServiceslSalesPerEmployeeJSON), true))
	router.Handle("/admin/api/all-services-sales-per-employee-graph/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(resportssales.GetAlServiceslSalesPerEmployeeJSONGraph), true))
	router.Handle("/admin/api/all-products-sales-per-employee/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(resportssales.GetAlProductsSalesPerEmployeeJSON), true))
	router.Handle("/admin/api/all-services-sales-per-service/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(resportssales.GetAlServicesSalesJSON), true))

	router.Handle("/admin/unfinishedcanceled", AdminCheck(http.HandlerFunc(renderUnfinishedCanceledPage), true))
	router.Handle("/admin/api/all-cancellations/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(cancellation.AllCancelledAppointmentJSON), true))
	router.Handle("/admin/api/all-cancellations-per-day-graph/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(cancellation.GetAllAppointmentsPesServiceStatusSONGraph), true))
	router.Handle("/admin/api/total-appointments-per-status/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(cancellation.GetAllAppointmentsPerStatusJSON), true))

	//Revenue Reports
	router.Handle("/admin/reportrevenueexpenses", AdminCheck(http.HandlerFunc(renderReportRevenueExpensesPage), true))
	router.Handle("/admin/api/total-revenue-sales-and-services/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(revenueexpenses.CallRevenueDataJSON), true))
	router.Handle("/admin/api/total-revenue-sales-and-services-graph/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(revenueexpenses.GetAllSalesPerServiceProductPerDateJSONGraph), true))

	router.Handle("/admin/api/total-expenses-sales/{id}/{storeid}/{from}/{to}", AdminCheck(http.HandlerFunc(revenueexpenses.TotalExpensesDataJSON), true))

	router.Handle("/admin/reportpendingdays", AdminCheck(http.HandlerFunc(renderReportPendingDaysPage), true))
	router.Handle("/admin/api/all-pending-services-per-user/{userid}", AdminCheck(http.HandlerFunc(pendaingdays.AllSalesTransCustomerJSON), true))
	router.Handle("/admin/reportcustomerswithbalance", AdminCheck(http.HandlerFunc(renderReportCustomersWithBalancePage), true))

	router.Handle("/admin/api/count-pending-appointment/{id}/{storeid}/{today}", AdminCheck(http.HandlerFunc(CountAllPendingAppointmentJSON), true))
	router.Handle("/admin/api/count-cancelled-appointment/{id}/{storeid}/{today}", AdminCheck(http.HandlerFunc(CountAllCancelledAppointmentJSON), true))
	router.Handle("/admin/api/total-sales-per-month/{id}/{storeid}", AdminCheck(http.HandlerFunc(GetAllSalesPerMonthPerStoreJSON), true))
	router.Handle("/admin/api/total-expenses-per-month/{id}/{storeid}", AdminCheck(http.HandlerFunc(GetAllExpensesPerMonthPerStoreJSON), true))

	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.Handle("/", router)

	// GetStatheresGiortes()

	// storesettings.BankHolidays()

	//	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	http.ListenAndServe(":8001", nil)

}
