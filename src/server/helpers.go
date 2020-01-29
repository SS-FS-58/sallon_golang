package main

import (
	"bookings"
	"cancellation"
	"customers"
	"database/sql"
	"expenses"
	"fmt"
	"income"
	"log"
	"mysqldb"
	"net/http"
	"pendaingdays"
	"promotions"
	"resportssales"
	"revenueexpenses"
	"storesettings"
)

type Salon struct {
	Title     string
	Active    string
	MenuTitle string
	Username  string
	VatNumber string
	ImageName int
	UserID    int
}
type HTTPResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func GetIndex() Users {
	var result Users

	result.Active = "index"
	result.Title = "Home"
	result.MenuTitle = "Home"
	return result
}

func GetSubscription() Salon {
	var result Salon

	result.Active = "subscription"
	result.Title = "subscription"
	result.MenuTitle = "subscription"

	return result

}
func GetBilling() Salon {
	var result Salon

	result.Active = "Billing"
	result.Title = "Billing"
	result.MenuTitle = "Billing"

	return result

}

// Get404 return error if page doesnt exists
func Get404() Salon {
	var result Salon
	result.Active = "404"
	result.Title = "Salon - Page not found"
	return result
}

func GetLogin() Users {
	var result Users

	result.Active = "login"
	result.Title = "Login"

	return result
}

func GetRegister() Users {
	var result Users

	result.Active = "register"
	result.Title = "register"

	return result
}

func GetMyProfile() Users {
	var result Users

	result.Active = "profile"
	result.Title = "My Profile"

	return result
}

//-------------Main menu items---------------//

func GetCustomers() customers.Customers {
	var result customers.Customers

	result.Active = "customers"
	result.Title = "Customers"
	result.MenuTitle = "Customers"

	return result
}

func GetViewCustomer() customers.Customers {
	var result customers.Customers

	result.Active = "viewcustomers"
	result.Title = "viewCustomers"
	result.MenuTitle = "ΝΙΚΟΛΑΟΣ ΜΠΑΤΖΙΟΣ"

	return result
}

func GetCashRegister() Salon {
	var result Salon

	result.Active = "cashregister"
	result.Title = "Cashier"
	result.MenuTitle = "Cashier"

	return result
}

func GetAccounting() Salon {
	var result Salon

	result.Active = "accounting"
	result.Title = "Λογιστήριο"
	result.MenuTitle = "Λογιστήριο"

	return result
}

func GetCampaign() Salon {
	var result Salon

	result.Active = "campaign"
	result.Title = "Campaigns"
	result.MenuTitle = "Campaigns"

	return result
}

func GetReports() Salon {
	var result Salon

	result.Active = "reports"
	result.Title = "Reports"
	result.MenuTitle = "Reports"

	return result
}
func GetReportSales() resportssales.SalesReports {
	var result resportssales.SalesReports

	result.Active = "reportsales"
	result.Title = "Sales Report"
	result.MenuTitle = "Sales Report"

	return result
}
func GetUnfinishedCanceled() cancellation.CancellationReports {
	var result cancellation.CancellationReports

	result.Active = "unfinishedcanceled"
	result.Title = "Cancelled Appointments"
	result.MenuTitle = "Cancelled Appointments"

	return result
}
func GetReportRevenueExpenses() revenueexpenses.RevenueReports {
	var result revenueexpenses.RevenueReports

	result.Active = "reportrevenueexpenses"
	result.Title = "Report Revenue Expenses"
	result.MenuTitle = "Report Revenue Expenses"

	return result
}
func GetReportPendingDays() pendaingdays.PendingDays {
	var result pendaingdays.PendingDays

	result.Active = "reportpendingdays"
	result.Title = "Report Pending Appointments"
	result.MenuTitle = "Report Pending Appointments"

	return result
}
func GetReportCustomersWithBalance() Salon {
	var result Salon

	result.Active = "reportcustomerswithbalance"
	result.Title = "Report Customers With Balance"
	result.MenuTitle = "Report Customers With Balance"

	return result
}
func GetStoreSettings() storesettings.Shops {
	var result storesettings.Shops

	result.Active = "storesettings"
	result.Title = "Παράμετροι καταστήματος"
	result.MenuTitle = "Παράμετροι καταστήματος"

	return result
}
func GetStoreSettingsSmsEmail() storesettings.Shops {
	var result storesettings.Shops

	result.Active = "storesettingssmsemail"
	result.Title = "Παράμετροι sms email"
	result.MenuTitle = "Παράμετροι sms email"

	return result
}

func GetShops() storesettings.Shops {
	var result storesettings.Shops

	result.Active = "shops"
	result.Title = "Stores"
	result.MenuTitle = "Stores"

	return result
}

//-------------sub menu items---------------//

func GetStoreSettingsServices() storesettings.Services {
	var result storesettings.Services

	result.Active = "storesettingsservices"
	result.Title = "Καταχώρηση Υπηρεσιών"
	result.MenuTitle = "Καταχώρηση Υπηρεσιών"

	return result
}

func GetHairdressers() storesettings.Hairdressers {
	var result storesettings.Hairdressers

	result.Active = "hairdressers"
	result.Title = "Hairdressers"
	result.MenuTitle = "Hairdressers"

	return result
}

func GetStoreSuppliers() storesettings.Suppliers {
	var result storesettings.Suppliers

	result.Active = "suppliers"
	result.Title = "Suppliers"
	result.MenuTitle = "Suppliers"

	return result
}

func GetStoreProducts() storesettings.Products {
	var result storesettings.Products

	result.Active = "products"
	result.Title = "Products"
	result.MenuTitle = "Products"

	return result
}

func GetCampaignGoogle() Salon {
	var result Salon

	result.Active = "campaigngoogle"
	result.Title = "Google campaigns"
	result.MenuTitle = "Google campaigns"

	return result
}

func GetCampaignFacebook() Salon {
	var result Salon

	result.Active = "campaignfacebook"
	result.Title = "Facebook campaigns"
	result.MenuTitle = "Facebook campaigns"

	return result
}

func GetCampaignSms() Salon {
	var result Salon

	result.Active = "campaignsms"
	result.Title = "Sms campaigns"
	result.MenuTitle = "Sms campaigns"

	return result
}

func GetCampaignEmail() Salon {
	var result Salon

	result.Active = "campaignemail"
	result.Title = "Email Καμπάνιες"
	result.MenuTitle = "Email Καμπάνιες"

	return result
}

func GetCalendar() bookings.Calendar {
	var result bookings.Calendar

	result.Active = "calendar"
	result.Title = "Calendar"
	result.MenuTitle = "Calendar"

	return result
}
func GetCheckoutCalendar() bookings.CheckOutCalendar {
	var result bookings.CheckOutCalendar

	result.Active = "calendar"
	result.Title = "Checkout"
	result.MenuTitle = "Calendar"

	return result
}

func GetRewards() Salon {
	var result Salon

	result.Active = "rewards"
	result.Title = "Rewards"
	result.MenuTitle = "Rewards"

	return result
}
func GetPromotions() promotions.Promotions {
	var result promotions.Promotions

	result.Active = "promotions"
	result.Title = "Promotions"
	result.MenuTitle = "Promotions"

	return result
}
func GetRecommendations() Salon {
	var result Salon

	result.Active = "recommendations"
	result.Title = "Recommendations"
	result.MenuTitle = "Recommendations"

	return result
}
func GetExpenses() expenses.Expenses {
	var result expenses.Expenses

	result.Active = "Expenses"
	result.Title = "Expenses"
	result.MenuTitle = "Expenses"

	return result
}
func GetIncome() income.Income {
	var result income.Income

	result.Active = "Income"
	result.Title = "Income"
	result.MenuTitle = "Income"

	return result
}
func add(x, y int) int {
	return x + y
}
func GetHost(r *http.Request) string {
	if r.TLS == nil {
		return "http://" + r.Host
	}
	return "https://" + r.Host
}
func rowExists(query string, args ...interface{}) bool {
	db := mysqldb.Connect()
	defer db.Close()
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}
