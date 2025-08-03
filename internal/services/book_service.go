package services

import (
	"crud/internal/models"
	"crud/internal/repositories"
)

func GetBooks() ([]models.Book, error) {
	return repositories.GetAllBooks()
}
