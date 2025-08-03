package controllers

import (
	"crud/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books, err := services.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت کتاب‌ها"})
		return
	}
	c.JSON(http.StatusOK, books)

}
