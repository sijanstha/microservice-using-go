package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	DB_USER_NAME     = "db_user_name"
	DB_USER_PASSWORD = "db_user_password"
	DB_HOST          = "db_host"
	SCHEMA           = "schema"
)

var (
	Client *sql.DB

	username = os.Getenv(DB_USER_NAME)
	password = os.Getenv(DB_USER_PASSWORD)
	host     = os.Getenv(DB_HOST)
	schema   = os.Getenv(SCHEMA)
)

// this function gets invoked if this package is imported
func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connection successful")
}
