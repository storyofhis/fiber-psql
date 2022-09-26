package handlers

import (
	"net/http"
	"postgresql-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
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
	err := r.repo.DB.Where("id = ?", id).First(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "could not get book",
			},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "book id fetched successfully",
			"data":   bookModel,
		},
	)
	return nil
}
