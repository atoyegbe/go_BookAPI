package book

import (
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"

	"fiber/web_app/database"
)

type Book struct {
	gorm.Model

	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.SendStatus(503)
	}

	db.Create(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	new_book := new(Book)

	var book Book
	db.Find(&book, id)

	if err := c.BodyParser(new_book); err != nil {
		return c.SendStatus(503)
	}

	book.Title = new_book.Title
	book.Author = new_book.Author
	book.Rating = new_book.Rating

	db.Updates(&book)

	return c.JSON(book)

}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.SendString("No Book Found with ID")
	}

	db.Delete(&book)
	return c.SendString("Book Successfully deleted")

}
