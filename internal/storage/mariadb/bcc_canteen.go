package mariadb

import (
	"canteen/internal/env"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func DbInit() *sql.DB {
	env.InitEnv()
	config := mysql.NewConfig()

	config.Addr = "localhost:3306"
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.DBName = os.Getenv("DB_NAME")
	config.ParseTime = true
	config.Loc = time.UTC

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Printf("Connection with database %s established\n", os.Getenv("DB_NAME"))
	return db
}

var Db = DbInit()
var NoRows error = sql.ErrNoRows
