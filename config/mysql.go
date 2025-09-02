package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB
var err error

func InitMysql() {
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-crud")
	if err != nil {
		panic(err)
	}
	createBooksTable()

}
func createBooksTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100),
		Author VARCHAR(100)
	);
	`
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
