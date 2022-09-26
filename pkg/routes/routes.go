package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.repo.CreateBook)
	api.Get("/get_books/:id", r.repo.GetBookByID)
	api.Get("/get_books", r.repo.GetBooks)
	api.Get("/delete_books/:id", r.repo.DeleteBook)
}
