package repositories

import (
	"crud/config"
	"crud/internal/models"
)

func GetAllBooks() ([]models.Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil

}
func SaveBook(book models.Book) (models.Book, error) {
	res, err := config.DB.Exec("insert into books (title,author) values (?,?)", book.Title, book.Author)

	id, _ := res.LastInsertId()
	book.ID = int(id)
	return book, err
}
func FindBook(id int) (*models.Book, error) {
	row := config.DB.QueryRow("SELECT * FROM books where id = ? ", id)

	var b models.Book
	if err := row.Scan(&b.ID, &b.Title, &b.Author); err != nil {
		return nil, err
	}
	return &b, nil
}
