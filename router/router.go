package router

import (
	"github.com/Alperen10/booksAPI/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes function
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/book")
	// routes
	v1.Get("/", handler.GetAllBooks)
	v1.Get("/single", handler.GetSingleBook)
	v1.Post("/", handler.CreateBook)
	v1.Put("/", handler.UpdateBook)
	v1.Delete("/", handler.DeleteBookByID)
}
