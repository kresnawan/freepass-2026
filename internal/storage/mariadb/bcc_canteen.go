package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DbInit() {
	config := mysql.NewConfig()

	// Database config
	config.Addr = "localhost:3306"
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.DBName = ""

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Printf("Connection with database %s established\n", os.Getenv("DB_NAME"))
}
