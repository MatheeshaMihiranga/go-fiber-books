package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/go-fiber-books/database"
	"github.com/yourusername/go-fiber-books/handlers"
)

func main() {
    // Initialize the database
    database.Connect()

    // Create a new Fiber app
    app := fiber.New()

    // Routes
    app.Post("/books", handlers.CreateBook)
    app.Get("/books", handlers.GetBooks)
    app.Get("/books/:id", handlers.GetBook)
    app.Put("/books/:id", handlers.UpdateBook)
    app.Delete("/books/:id", handlers.DeleteBook)

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
