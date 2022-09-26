package main

import (
	"log"
	"net/http"
	"os"
	"postgresql-gorm/models"
	"postgresql-gorm/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "request failed gan",
			},
		)
		return err
	}
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
			"status": "book has been added mantap",
		},
	)

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	if err := r.DB.Find(bookModels).Error; err != nil {
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

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"stastus": "id cannot be empty",
			},
		)
		return nil
	}
	err := r.DB.Delete(bookModel, id)
	if err.Error != nil {
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
	err := r.DB.Where("id = ?", id).First(&bookModel).Error
	// err := r.DB.Find(bookModel).Error
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
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/get_books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the databases")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := &Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
