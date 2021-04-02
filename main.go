package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fiber/web_app/book"
	"fiber/web_app/database"

	
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Put("/api/book/:id", book.UpdateBook)
	app.Delete("/api/book", book.DeleteBook)
}


func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Databse")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Dtabase Migrated")
}


func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")

}



