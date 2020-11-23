package database

import (
	"database/sql"
	// mysql blank import is used to init driver
	_ "github.com/go-sql-driver/mysql"
)

// DBConn is used to connect to mysql database
func DBConn() *sql.DB {
	sqlDriver := "mysql"
	userName := "root"
	passWord := "Bikhungha1!"
	host := "127.0.0.1"
	port := "3306"
	dbName := "gin_demo"
	db, err := sql.Open(sqlDriver, userName+":"+passWord+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
