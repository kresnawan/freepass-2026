package mariadb

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func DbInit() *sql.DB {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Environment variables loaded")
	config := mysql.NewConfig()
	var db *sql.DB
	var err error

	config.Addr = os.Getenv("DB_HOST")
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.DBName = os.Getenv("DB_NAME")
	config.ParseTime = true
	config.Loc = time.UTC

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", config.FormatDSN())
		err = db.Ping()

		if err != nil {
			log.Println("Connection failed, retrying..")
		} else {
			log.Println("Connected to MariaDB")
			break
		}

		time.Sleep(3 * time.Second)
	}

	return db
}

var Db = DbInit()
var NoRows error = sql.ErrNoRows
