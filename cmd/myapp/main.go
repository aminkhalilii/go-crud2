package main

import (
	"crud/config"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.InitMysql()
	config.InitRedis()

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")

	defer config.DB.Close()

}
