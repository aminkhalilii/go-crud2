package services

import (
	"crud/config"
	"crud/internal/models"
	"crud/internal/repositories"
	"encoding/json"
	"fmt"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	cached, err := config.RedisClient.Get(config.Ctx, "books").Result()
	if err == nil && cached != "" {
		json.Unmarshal([]byte(cached), &books)
		return books, nil
	}

	books, err = repositories.GetAllBooks()
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(books)
	config.RedisClient.Set(config.Ctx, "books", data, 0)

	return books, nil
}
func SaveBook(book models.Book) error {
	data, err := repositories.SaveBook(book)
	if err != nil {
		return err
	}
	config.RedisClient.HSet(config.Ctx, "book", fmt.Sprintf("%d", book.ID), data)
	return nil
}
func GetBook(id int) (*models.Book, error) {
	cached, err := config.RedisClient.HGet(config.Ctx, "book", fmt.Sprintf("%d", id)).Result()
	if err == nil && cached != "" {
		var book models.Book
		if err := json.Unmarshal([]byte(cached), &book); err == nil {
			return &book, nil
		}
	}

	book, err := repositories.FindBook(id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(book)
	config.RedisClient.HSet(config.Ctx, "book", fmt.Sprintf("%d", book.ID), data)

	return book, nil
}
