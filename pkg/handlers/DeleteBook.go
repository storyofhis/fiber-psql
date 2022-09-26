package handlers

import (
	"net/http"
	"postgresql-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"status": "id cannot be empty",
			},
		)
		return nil
	}

	if err := r.repo.DB.Delete(bookModel, id); err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "could not delete book",
			},
		)
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "books delete successfully",
			"data":   bookModel,
		},
	)
	return nil
}
