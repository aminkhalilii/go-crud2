package config

import "database/sql"

var DB *sql.DB
var err error

func InitMysql() {
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-crud")
	if err != nil {
		panic(err)
	}
}
