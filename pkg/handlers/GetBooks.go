package handlers

import (
	"net/http"
	"postgresql-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	if err := r.repo.DB.Find(bookModels).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "could not get books data",
			},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "books fetched successfully",
			"data":   bookModels,
		},
	)
	return nil
}
