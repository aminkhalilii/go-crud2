package controllers

import (
	"crud/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	c.JSON(http.StatusOK, books)
}
