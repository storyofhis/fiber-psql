package handlers

import (
	"net/http"
	"postgresql-gorm/pkg/helper"
	"postgresql-gorm/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func (r *routes.Repository) CreateBook(context *fiber.Ctx) error {
	book := helper.Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "request failed gan kwkw",
			},
		)
		return err
	}

	// err = r.repo.DB.Create(&book).Error
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "could not create book gan",
			},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "book has been added mantap gan",
			"data":   book,
		},
	)
	return nil
}
