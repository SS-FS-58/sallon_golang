package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Connect() *sql.DB {
	connString := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASS") + "@(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DB") + "?parseTime=true"

	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to database" + err.Error())
	}

	return db
}

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}
