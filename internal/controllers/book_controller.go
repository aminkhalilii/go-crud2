package controllers

import (
	"crud/internal/models"
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
func SaveBook(c *gin.Context) {

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := services.SaveBook(book)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save book"})
		return
	}

	c.JSON(200, gin.H{"message": "Book saved successfully"})

}
