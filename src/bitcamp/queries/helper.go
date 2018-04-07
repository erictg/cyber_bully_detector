package queries

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"os"
	"fmt"
	"log"
)

var db *sql.DB = nil

func GetDB() (*sql.DB, error){

	if db == nil{
		host := ""
		if os.Getenv("MYSQL_HOST") == ""{
			host = "localhost"
		}else{
			host = os.Getenv("MYSQL_HOST")
		}

		uname := ""
		if os.Getenv("MYSQL_USER") == ""{
			uname = "dbuser"
		}else{
			uname = os.Getenv("MYSQL_USER")
		}

		pwd := ""
		if os.Getenv("MYSQL_PASSWORD") == ""{
			pwd = "userpass"
		}else {
			pwd = os.Getenv("MYSQL_PASSWORD")
		}



		connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/psych?parseTime=true", uname, pwd, host)

		log.Println(connectionString)

		dbConn, err := sql.Open("mysql", connectionString)
		if err != nil{
			return nil, err
		}

		pingErr := dbConn.Ping()
		if pingErr != nil{
			return nil, pingErr
		}


		//db is the package level db
		db = dbConn
		db.SetMaxIdleConns(0)
	}

	return db, nil
}