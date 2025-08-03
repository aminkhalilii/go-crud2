package main

import (
	"crud/config"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.InitMysql()

	defer config.DB.Close()

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
