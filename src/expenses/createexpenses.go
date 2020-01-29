package expenses

import (
	"encoding/json"
	"fmt"
	"log"
	"mysqldb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type InsertExpense struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	StoreID        int       `json:"store_id"`
	ExpensePrice   float64   `json:"expense_price"`
	CostCategoryid int       `json:"expense_cost_category"`
	IsPaid         bool      `json:"expense_paid_unpaid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type AllExpensesPerUser struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	StoreID          int       `json:"store_id"`
	StoreName        string    `json:"store_name"`
	ExpensePrice     float64   `json:"expense_price"`
	CostCategoryid   int       `json:"expense_cost_category"`
	IsPaid           bool      `json:"expense_paid_unpaid"`
	CostCategoryName string    `json:"cost_category"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ExpenseConfig struct {
	AllExpenses []AllExpensesPerUser `json:"data"`
}

func InsertExpenseToDB(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var expense InsertExpense

	err := decoder.Decode(&expense)
	if err != nil {
		print(err.Error())
	}

	stmt, _ := db.Prepare(`INSERT INTO expensestrans (user_id,
		store_id,
		expenses_list_id,
		expenses_price,
		payment_type,
		created_at,
		updated_at) VALUES (?,?,?,?,?, NOW(), NOW())`)
	res, err := stmt.Exec(expense.UserID, expense.StoreID, expense.CostCategoryid, expense.ExpensePrice, expense.IsPaid)

	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to insert expense into database"})
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Inserted expense Into the Database", Body: strconv.Itoa(int(id))})
}
func GetAllExpensesPerUser(userID int) []AllExpensesPerUser {
	db := mysqldb.Connect()
	defer db.Close()

	var expenses []AllExpensesPerUser

	results, err := db.Query(`SELECT expensestrans.id, 
		expensestrans.user_id,
		expensestrans.store_id,
		expensestrans.expenses_price,
		expensestrans.payment_type,
		shops.company_name,
		expensestrans.expenses_list_id,
		expenses_list.expenses_name,
		expensestrans.created_at,
		expensestrans.updated_at 
		FROM expensestrans 
		LEFT JOIN shops ON expensestrans.store_id = shops.id
		LEFT JOIN expenses_list on expensestrans.expenses_list_id = expenses_list.id
		WHERE expensestrans.user_id = ?`, userID)
	if err != nil {
		println(err.Error())
	}

	defer results.Close()
	for results.Next() {
		var expense AllExpensesPerUser
		err = results.Scan(&expense.ID, &expense.UserID, &expense.StoreID, &expense.ExpensePrice, &expense.IsPaid, &expense.StoreName, &expense.CostCategoryid, &expense.CostCategoryName, &expense.CreatedAt, &expense.UpdatedAt)
		if err != nil {
			println(err.Error())
		}
		expenses = append(expenses, expense)
	}
	return expenses
}

func AllExpensesJSONPerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(err.Error())
		fmt.Fprintf(w, "Not a valid id")
	}
	expenses := GetAllExpensesPerUser(userID)
	err = json.NewEncoder(w).Encode(ExpenseConfig{AllExpenses: expenses})
	if err != nil {
		println(err.Error())
	}
}
func UpdateExpenseJSON(w http.ResponseWriter, r *http.Request) {

	db := mysqldb.Connect()
	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	var expense InsertExpense

	err := decoder.Decode(&expense)
	if err != nil {
		print(err.Error())
	}
	vars := mux.Vars(r)

	expenseID := vars["id"]
	ID, err := strconv.Atoi(expenseID)

	if err != nil {
		print(err.Error())
	}

	stmt, err := db.Prepare(`UPDATE expensestrans SET
		user_id = ?,
		store_id = ?,
		expenses_list_id = ?,
		expenses_price = ?,
		payment_type = ?,
		updated_at = NOW() WHERE id = ?`)
	_, err = stmt.Exec(expense.UserID, expense.StoreID, expense.CostCategoryid, expense.ExpensePrice, expense.IsPaid, ID)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update  expense in the Database"})
	}

	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update expense in the Database"})
}

func PaidUnpaid(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var expense InsertExpense
	err := decoder.Decode(&expense)

	if err != nil {
		println(err.Error())
	}

	vars := mux.Vars(r)

	expenseID := vars["id"]
	ID, err := strconv.Atoi(expenseID)

	if err != nil {
		println(err.Error())
	}

	stmt, _ := db.Prepare("UPDATE expensestrans SET payment_type = ? WHERE id = ?")
	_, err = stmt.Exec(expense.IsPaid, ID)
	if err != nil {
		println(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to update expense in the database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update expense in the Database"})
}
func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	categoryID := vars["id"]

	ID, _ := strconv.Atoi(categoryID)

	stmt, err := db.Prepare("DELETE FROM expensestrans where id = ?")

	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to delete expense from database"})
	}
	json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully deleted expense from database "})
}
