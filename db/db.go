package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Menggunakan configuration sering terjadi masalahh

// var db *sql.DB
// var err error

// func Init() {
// 	conf := config.GetConfig()

// 	conn := conf.USERNAME + ":" + conf.PASSWORD + "@tcp(" + conf.HOST + ":" + conf.PORT + ")/" + conf.NAME

// 	db, err := sql.Open("mysql", conn)
// 	if err != nil {
// 		panic("SERVER CONECTION ERROR")
// 	}

// 	if db == nil {
// 		panic("DNS INVALID")
// 	}

// }

// func CreateCon() *sql.DB {
// 	return db
// }

// Saya mencoba lebih mempersingkat db

func CreateCon() *sql.DB {
	driver := "mysql"
	user := "root"
	pass := ""
	dbName := "go_crud"

	db, _ := sql.Open(driver, user+":"+pass+"@/"+dbName)

	return db
}
