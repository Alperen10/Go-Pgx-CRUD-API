package handler

import (
	"context"
	"fmt"

	"github.com/Alperen10/booksAPI/database"
	"github.com/Alperen10/booksAPI/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	db := database.Connection
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	insertBook := `insert into books (book_name, author, description) values($1, $2, $3)`

	_, err = db.Exec(context.Background(), insertBook, book.BookName, book.Author, book.Description)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	// return the created book
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Book has created", "data": book})

}

func GetAllBooks(c *fiber.Ctx) error {
	db := database.Connection
	books := []models.Book{}

	// find all books in the database
	selectAllBooks := `select * from "books"`
	rows, err := db.Query(context.Background(), selectAllBooks)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})

	}

	for rows.Next() {
		currentBook := models.Book{}
		err = rows.Scan(&currentBook.Id, &currentBook.BookName, &currentBook.Author, &currentBook.Description)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not read book", "data": err})
		}
		fmt.Printf("Id: %d - %s - %s - %s\n", currentBook.Id, currentBook.BookName, currentBook.Author, currentBook.Description)
		books = append(books, currentBook)

	}

	// return books
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "All books reading", "data": books})

}

func GetSingleBook(c *fiber.Ctx) error {
	db := database.Connection
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// find single book in the database by id
	selectBook := `select book_name, author, description from books where id = $1`
	err = db.QueryRow(context.Background(), selectBook, book.Id).Scan(&book.BookName, &book.Author, &book.Description)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to getting book", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Getting a book", "data": book})
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.Connection
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// find single book in the database by id
	updateBooks := `update books set book_name = $1 where id = $2`
	res, err := db.Exec(context.Background(), updateBooks, book.BookName, book.Id)

	if err == nil {
		count := res.RowsAffected()
		if count == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
		}
		// return the updated book
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Book deleted"})
	}

	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
}

func DeleteBookByID(c *fiber.Ctx) error {
	db := database.Connection
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// find book in the database by id
	deleteBooks := `delete from books where id = $1`
	res, err := db.Exec(context.Background(), deleteBooks, book.Id)

	if err == nil {
		count := res.RowsAffected()
		if count == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
		}

		// if deleted successfully it will return this message
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Book deleted"})

	}

	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
}
