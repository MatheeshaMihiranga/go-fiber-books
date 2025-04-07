package services

import (
	"errors"

	"github.com/yourusername/go-fiber-books/database"
	"github.com/yourusername/go-fiber-books/models"
)

// CreateBook adds a new book to the database
func CreateBook(book *models.Book) error {
    result := database.DB.Create(book)
    return result.Error
}

// GetBooks retrieves all books
func GetBooks() ([]models.Book, error) {
    var books []models.Book
    result := database.DB.Find(&books)
    return books, result.Error
}

// GetBook retrieves a book by ID
func GetBook(id uint) (*models.Book, error) {
    var book models.Book
    result := database.DB.First(&book, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &book, nil
}

// UpdateBook updates an existing book
func UpdateBook(id uint, newBook *models.Book) (*models.Book, error) {
    book, err := GetBook(id)
    if err != nil {
        return nil, err
    }
    book.Title = newBook.Title
    book.Author = newBook.Author
    book.Year = newBook.Year
    result := database.DB.Save(book)
    return book, result.Error
}

// DeleteBook removes a book from the database
func DeleteBook(id uint) error {
    result := database.DB.Delete(&models.Book{}, id)
    if result.RowsAffected == 0 {
        return errors.New("book not found")
    }
    return result.Error
}
