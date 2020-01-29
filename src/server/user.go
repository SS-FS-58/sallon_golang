package main

import (
	"context"
	"crypto/rand"
	"customers"
	"encoding/base64"
	"encoding/json"
	"fmt"
	tr "gettext"
	"html/template"
	"io"
	"log"
	"mysqldb"
	"net/http"
	"os"
	"storesettings"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Users struct {
	Active               string
	Title                string
	Username             string
	ImageName            int
	VatNumber            string
	UserProfile          string
	CompanyName          string
	Role                 string
	MenuTitle            string
	MessageError         string
	Users                []User
	User                 User
	UserID               int
	NameDays             Giortes
	BrthdayCustomers     []customers.Customer
	LenBrthdayCustomers  int
	LenNameDaysCustomers int
	LentTotal            int
	TimeNow              time.Time
	UserEmail            string
	TotalStore           int
	TotalAppoi           int
}

//User holds the data for each user
type User struct {
	ID                  int       `json:"id"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	Forename            string    `json:"forename"`
	Surname             string    `json:"surname"`
	CompanyName         string    `json:"company_name"`
	CompanyAddress      string    `json:"company_address"`
	CompanyStreetNumber int       `json:"company_street_number"`
	CompanyCity         string    `json:"company_city"`
	CompanyState        string    `json:"company_state"`
	CompanyZipCode      string    `json:"company_zip_code"`
	CompanyCountry      string    `json:"company_country"`
	HomeTelephone       string    `json:"work_telephone"`
	MobileTelephone     string    `json:"mobile_telephone"`
	Role                string    `json:"role"`
	IsActive            bool      `json:"is_active"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	LastLogin           time.Time `json:"last_login"`
}

var loginTemplate = template.Must(template.New("login.html").Funcs(templateFuncMap).ParseFiles("views/login.html"))
var registerTemplate = template.Must(template.New("register.html").Funcs(templateFuncMap).ParseFiles("views/register.html"))
var packageTemplate = template.Must(template.New("registerpackages.html").Funcs(templateFuncMap).ParseFiles("views/registerpackages.html"))

func Auth(users *User, password string, w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()
	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	context := GetRegister()
	if err != nil {
		println(err.Error())
		http.Redirect(w, r, r.URL.Path, 301)
		return
	}
	if users.IsActive == false {
		fmt.Println(users.ID)

		token := string(genRandByte())

		stmt, _ := db.Prepare(`INSERT INTO users_activation (user_id,
			username,
			email,
			activation_code,
			creation_date) VALUES (?,?,?,?, NOW())`)

		res, err := stmt.Exec(users.ID, users.Username, users.Email, token)

		if err != nil {
			context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Failed to create activation URL</strong></div>"
			registerTemplate.Execute(w, context)
			return
		}
		idint, err := res.LastInsertId()
		if err != nil {
			println(err.Error())
			return
		}
		if idint > 0 {

			registerTemplate.Execute(w, context)

		}

	}

}
func renderRegister(w http.ResponseWriter, r *http.Request) {
	context := GetRegister()
	registerTemplate.Execute(w, context)
}

// Register function create a user
func Register(w http.ResponseWriter, r *http.Request) {

	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	if err := r.ParseMultipartForm(10000000); err != nil {
		println(err.Error())
	}
	context := GetRegister()
	db := mysqldb.Connect()
	defer db.Close()
	if r.Method != "POST" {

		return
	}
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	forename := r.FormValue("forename")
	surname := r.FormValue("surname")
	companyName := r.FormValue("company_name")
	homeTelephone := r.FormValue("work_telephone")
	mobileTelephone := r.FormValue("mobile_telephone")

	roleAdmin := "administrator"
	isActive := true
	role := QueryAdministrator()
	println(role.Role)
	if len(role.Role) != 0 {
		roleAdmin = "storeadmin"
		isActive = false
	}

	users := QueryUser(username)
	if len(users.Username) != 0 {
		context.MessageError = "<div class='alert alert-danger show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>The username  already exist.</div>"
		registerTemplate.Execute(w, context)
		return
	}

	if len(users.Email) != 0 {
		context.MessageError = "<div class='alert alert-danger show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>The email  already exist.</div>"
		registerTemplate.Execute(w, context)
		return
	}

	if (User{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			print(err.Error())
		}

		if len(hashedPassword) != 0 && err == nil {
			stmt, err := db.Prepare("INSERT users SET  username = ?, forename = ?, surname = ?, email = ?, password = ?,company_name = ?,work_telephone = ?, mobile_telephone= ?,role = ?,is_active = ?, created_at = NOW(), updated_at = NOW(),last_login = NOW()")

			if err != nil {
				print(err.Error())
				return
			}
			row, err := stmt.Exec(&username, &forename, &surname, &email, &hashedPassword, &companyName, &homeTelephone, &mobileTelephone, &roleAdmin, &isActive)

			if err != nil {
				print(err.Error())
				return
			}
			id, _ := row.LastInsertId()
			if userProfileImg, ok := r.MultipartForm.File["logo"]; ok && len(userProfileImg) == 1 {
				if f, err := userProfileImg[0].Open(); err == nil {
					userImageName := "stor_admin_" + strconv.Itoa(int(id)) + ".jpg"
					out, err := os.Create("static/images/" + userImageName)
					if err != nil {
						println(err.Error())
						fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
						return
					}

					defer out.Close()

					_, err = io.Copy(out, f)

					if err != nil {
						println(err.Error())
						fmt.Fprintln(w, err)
					}
					stmt, err := db.Prepare("UPDATE users SET logo_image = ? WHERE id = ?")
					_, err = stmt.Exec(userImageName, int(id))
					if err != nil {
						println(err.Error())
					}
					fmt.Println("File uploaded successfully: ", userImageName)
				}
			} else {
				if openFile, err := os.Open("static/images/faces/face-0.jpg"); err == nil {

					userImageName := "stor_admin_" + username + ".jpg"
					out, err := os.Create("static/images/" + userImageName)
					if err != nil {
						println(err.Error())
						fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
						return
					}

					defer out.Close()

					_, err = io.Copy(out, openFile)

					if err != nil {
						println(err.Error())
						fmt.Fprintln(w, err)
					}
					stmt, err := db.Prepare("UPDATE users SET logo_image = ? WHERE id = ?")
					_, err = stmt.Exec(userImageName, int(id))
					if err != nil {
						println(err.Error())
					}
					fmt.Println("File uploaded successfully: ", userImageName)

				}
			}

			Auth(&User{
				ID:              int(id),
				Username:        string(username),
				Forename:        string(forename),
				Surname:         string(surname),
				Email:           string(email),
				Password:        string(hashedPassword),
				CompanyName:     string(companyName),
				HomeTelephone:   string(homeTelephone),
				MobileTelephone: string(mobileTelephone),
				Role:            string(roleAdmin),
				IsActive:        bool(isActive),
			}, password, w, r)

		}
	}
}

func Admin(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	context := GetIndex()

	username := getUserName(r)
	user := QueryUser(username)

	if len(username) != 0 {
		context.ImageName = user.ID

		context.Username = user.CompanyName + " " + "<br>" + "(" + user.Username + ")"
		totalStores := storesettings.GetAllShopsPerUser(user.ID)
		context.TotalStore = len(totalStores)

		context.UserID = user.ID
		context.UserEmail = user.Email

		todayNames := GetStatheresGiortesFromDB()

		giortes := strings.Split(todayNames.StatusMessage, ",")
		var notonous []string
		for _, g := range giortes {
			t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
			result, _, _ := transform.String(t, g)
			notonous = append(notonous, strings.ToUpper(result))
		}
		olesoigiortes := strings.Join(notonous, ",")
		todayNames.StatusMessage = olesoigiortes
		// arrayNameDays := strings.Split(todayNames.StatusMessage, ",")
		context.NameDays = todayNames

		allcustomers := customers.GetAllCustomersPerUser(user.ID)

		var bthdayCustomers []customers.Customer

		toras := time.Now()
		var allonomata []string
		var totalb []int
		var totaln []int

		for _, g := range allcustomers {
			if g.DateOfBirth.Format("02-01") == toras.Format("02-01") {

				g.TitleCelebration = "Birthday"
				g.CelebreationType = "b"
				bthdayCustomers = append(bthdayCustomers, g)
				totalb = append(totalb, g.ID)
			}
			t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
			result, _, _ := transform.String(t, g.CustomerName)
			allonomata = append(allonomata, strings.ToUpper(result))
			for _, k := range notonous {
				if k == strings.ToUpper(result) {
					g.TitleCelebration = "Name Day"
					g.CelebreationType = "n"
					bthdayCustomers = append(bthdayCustomers, g)
					totaln = append(totaln, g.ID)
				}
			}
		}
		context.LenBrthdayCustomers = len(totalb)
		context.LentTotal = len(bthdayCustomers)
		context.BrthdayCustomers = bthdayCustomers
		context.LenNameDaysCustomers = len(totaln)
		context.TimeNow = time.Now()
		renderView("index.html", w, context)

		return
	}

	http.Redirect(w, r, "/admin/login", 301)

}

func Index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/admin", 301)
}

// AdminCheck check render the static content for the logged in user
func AdminCheck(next http.Handler, redirect bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := getUserName(r)
		if len(username) != 0 {
			ctx := context.WithValue(r.Context(), "username", username)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			if redirect {
				http.Redirect(w, r, "/admin/login", http.StatusTemporaryRedirect)
				return
			}
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)
	db := mysqldb.Connect()
	defer db.Close()
	context := GetLogin()

	if r.Method != "POST" {
		err := loginTemplate.Execute(w, context)

		if err != nil {
			println(err.Error())
		}
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")

	if role == "shop" {
		shop := storesettings.QueryShop(username)
		usersPerID := QueryUserID(shop.UserID)
		err := bcrypt.CompareHashAndPassword([]byte(shop.Password), []byte(password))
		if len(shop.VatNumber) < 1 {

			context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>The vat numbeer does not exist.</div>"
			// http.Error(w, "Your account has been deactivated", http.StatusForbidden)
			loginTemplate.Execute(w, context)
			return
		}

		if err != nil {
			context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Username or Password is incorrect. Please try again.</div>"
			print(err.Error())
			loginTemplate.Execute(w, context)
			return
		}
		if shop.IsActive == false {
			context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Your account has been deactivated</div>"
			// http.Error(w, "Your account has been deactivated", http.StatusForbidden)
			loginTemplate.Execute(w, context)
			return
		}

		setSession(usersPerID.Username, w)
		setVatNumberSession(shop.VatNumber, w)
		http.Redirect(w, r, "/admin/calendar", 302)
	}

	users := QueryUser(username)

	stmt, err := db.Prepare("UPDATE users SET last_login = NOW() WHERE username = ?")
	_, err = stmt.Exec(username)
	if err != nil {
		println(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if len(users.Username) < 1 {

		context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>The username does not exist.</div>"
		// http.Error(w, "Your account has been deactivated", http.StatusForbidden)
		loginTemplate.Execute(w, context)
		return
	}

	if err != nil {
		context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Username or Password is incorrect. Please try again.</div>"
		print(err.Error())
		loginTemplate.Execute(w, context)
		return
	}

	if users.IsActive == false {
		context.MessageError = "<div class='alert alert-danger  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Your account has been deactivated</div>"
		// http.Error(w, "Your account has been deactivated", http.StatusForbidden)
		loginTemplate.Execute(w, context)
		return
	}

	setSession(users.Username, w)
	http.Redirect(w, r, "/admin", 302)
}
func logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	clearVatnumberSession(w)
	http.Redirect(w, r, "/admin", 302)
}

func QueryUser(username string) User {
	db := mysqldb.Connect()
	defer db.Close()

	var user = User{}
	err := db.QueryRow("SELECT id, username,password, forename, surname, email,company_name,work_telephone, mobile_telephone,role,is_active, created_at, updated_at,last_login FROM users where username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Forename, &user.Surname, &user.Email, &user.CompanyName, &user.HomeTelephone, &user.MobileTelephone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.LastLogin)
	if err != nil {
		log.Println(err.Error())
	}
	return user
}
func QueryUserEmailGive(email string) User {
	db := mysqldb.Connect()
	defer db.Close()

	var user = User{}
	err := db.QueryRow("SELECT id, username,password, forename, surname, email,company_name,work_telephone, mobile_telephone,role,is_active, created_at, updated_at,last_login FROM users where email = ?", email).Scan(&user.ID, &user.Username, &user.Password, &user.Forename, &user.Surname, &user.Email, &user.CompanyName, &user.HomeTelephone, &user.MobileTelephone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.LastLogin)
	if err != nil {
		log.Println(err.Error())
	}
	return user
}

func QueryUserID(id int) User {
	db := mysqldb.Connect()
	defer db.Close()

	var user = User{}
	err := db.QueryRow("SELECT id, username,password, forename, surname, email,company_name,work_telephone, mobile_telephone,role,is_active, created_at, updated_at,last_login FROM users where id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Forename, &user.Surname, &user.Email, &user.CompanyName, &user.HomeTelephone, &user.MobileTelephone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.LastLogin)
	if err != nil {
		log.Println(err.Error())
	}
	return user
}

//QueryAdministrator retrun from database the administator role
func QueryAdministrator() User {
	db := mysqldb.Connect()
	defer db.Close()

	var users = User{}
	db.QueryRow(`SELECT role FROM users WHERE role  = 'administrator'`).Scan(&users.Role)

	return users
}

//CompanyNameToUpperCase CompanyNameToUpperCase
func CompanyNameToUpperCase(companyName string) string {
	return strings.ToLower(companyName)
}
func (s User) UserNameForenameSurnamePrefic() string {
	return s.Forename + "-" + s.Surname
}

//GetSingleUser return a single user from databse based on user id
func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	db := mysqldb.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		println(err.Error())
	}

	var user User
	err = db.QueryRow("SELECT id, username, forename, surname, email,company_name,work_telephone, mobile_telephone,role,is_active, created_at, updated_at,last_login  FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Forename, &user.Surname, &user.Email, &user.CompanyName, &user.HomeTelephone, &user.MobileTelephone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.LastLogin)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to select user from database"})
	}
	json.NewEncoder(w).Encode(user)
}

//UpdateLender update the lender when its calls
func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10000000); err != nil {
		println("err1", err.Error())
	}
	db := mysqldb.Connect()
	defer db.Close()

	println("update request")

	var user User
	vars := mux.Vars(r)
	userID := vars["id"]
	ID, _ := strconv.Atoi(userID)

	userJSON := r.MultipartForm.Value["userJSON"][0]

	err := json.Unmarshal([]byte(userJSON), &user)

	if err != nil {
		println("err3", err.Error())
	}
	var userImageName string
	if userlBkImg, ok := r.MultipartForm.File["userLogoImageFile"]; ok && len(userlBkImg) == 1 {
		// do something with carousel background image
		if f, err := userlBkImg[0].Open(); err == nil {
			userImageName = "stor_admin_" + strconv.Itoa(ID) + ".jpg"
			out, err := os.Create("static/images/" + userImageName)
			if err != nil {
				println(err.Error())
				fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
				return
			}

			defer out.Close()

			_, err = io.Copy(out, f)

			if err != nil {
				println(err.Error())
				fmt.Fprintln(w, err)
			}

			fmt.Println("File uploaded successfully: ", userImageName)
		}
	}
	if len(user.Password) != 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			print(err.Error())
		}
		stmt, _ := db.Prepare("UPDATE users SET username = ?, password= ?, forename = ?, surname = ?, email = ?,company_name = ?,work_telephone = ?, mobile_telephone = ?,logo_image = ?, updated_at = NOW() where id = ?")

		_, err = stmt.Exec(user.Username, hashedPassword, user.Forename, user.Surname, user.Email, user.CompanyName, user.HomeTelephone, user.MobileTelephone, userImageName, ID)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to Update user in the Database"})
		}

		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update user in the Database"})

	} else {
		stmt, _ := db.Prepare("UPDATE users SET username = ?, forename = ?, surname = ?, email = ?,company_name = ?,work_telephone = ?, mobile_telephone = ?, updated_at = NOW() where id = ?")

		_, err = stmt.Exec(user.Username, user.Forename, user.Surname, user.Email, user.CompanyName, user.HomeTelephone, user.MobileTelephone, ID)

		if err != nil {
			println(err.Error())
			json.NewEncoder(w).Encode(HTTPResp{Status: 500, Description: "Failed to Update Post in the Database"})
		}

		json.NewEncoder(w).Encode(HTTPResp{Status: 200, Description: "Successfully Update Post in the Database"})
	}

}

func packageRegister(w http.ResponseWriter, r *http.Request) {
	langParam := r.URL.Query().Get("lang")

	pathtolocale := os.Getenv("LOCALE_LANG")

	tr.Setup(langParam, langParam, pathtolocale)

	tr.ChangeLocale(langParam)

	db := mysqldb.Connect()
	defer db.Close()

	println("update request")

	vars := mux.Vars(r)
	email := vars["email"]
	token := vars["token"]
	context := GetRegister()

	if rowExists("SELECT id FROM users_activation WHERE  email = ? and activation_code = ?", email, token) {
		context.User = QueryUserEmailGive(email)
		context.MessageError = "<div class='alert alert-success  show' role='alert'> <button type='button' class='close' data-dismiss='alert' aria-label='Close'> <span aria-hidden='true'>&times;</span></button><strong>Thank you for you validation on more step before activate you account. </div>"

		packageTemplate.Execute(w, context)
		return
	}

	http.Redirect(w, r, "/admin/login", http.StatusTemporaryRedirect)
	return

}
func genRandByte() []byte {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return []byte(base64.URLEncoding.EncodeToString(b))
}
