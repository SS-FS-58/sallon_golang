// db_test.go
package mysqldb

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("root", "SailIn1985!", "salon_new")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role VARCHAR(120),
    forename VARCHAR(120),
    surname VARCHAR(120),
    work_telephone VARCHAR(120) NULL,
    mobile_telephone VARCHAR(120) NULL, 
    company_name VARCHAR(200) NOT NULL,
    username VARCHAR (50),
    email VARCHAR(50) NOT NULL ,
    password VARCHAR (120),
    is_active BOOL,
    logo_image VARCHAR(60),
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    last_login DATETIME,
    UNIQUE KEY username (username)
    ) ENGINE = MYISAM DEFAULT CHARSET = utf8;`
