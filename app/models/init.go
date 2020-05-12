package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

var db_name = "techTest"
var db_user = "root"
var db_password = "password"

func InitDB() {
	var err error
	db, err = sql.Open("mysql", db_user+":"+db_password+"@tcp(127.0.0.1:3306)/"+db_name)
	if err != nil {
		log.Panic(err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Println("Failled to ping the DB")
			log.Panic(err)
		} else {
			log.Println("DB connected!")
		}
	}
}
