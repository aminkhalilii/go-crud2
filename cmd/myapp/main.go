package main

import (
	"crud/config"
	"crud/internal/controllers"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.InitMysql()

	createBooksTable()

	defer config.DB.Close()

	router := gin.Default()

	router.GET("/books", controllers.GetBooks)

	router.Run(":8080")
}

func createBooksTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100),
		Author VARCHAR(100)
	);
	`
	_, err := config.DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
