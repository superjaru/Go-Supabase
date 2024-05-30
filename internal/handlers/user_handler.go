package handlers

import (
	"context"
	"fmt"
	"log"
	"supabase-fiber-SupaDB-project/internal/models"
	"supabase-fiber-SupaDB-project/internal/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetAllBooks(c *fiber.Ctx) error {
	users, err := h.repo.GetAllBooks(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetBookByName(c *fiber.Ctx) error {
	name := c.Params("name")
	fmt.Println("name1", name)
	book, err := h.repo.GetBookByName(context.Background(), name)
	fmt.Println("book", book)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func (h *UserHandler) CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if book.UID == uuid.Nil {
		book.UID = uuid.New()
	}

	if err := h.repo.CreateBook(context.Background(), &book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error at CreateBook": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *UserHandler) UpdateBook(c *fiber.Ctx) error {
	var book models.Book
	bookName := c.Params("name")
	existingBook, err := h.repo.GetBookByName(context.Background(), bookName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if book.Name != "" {
		existingBook.Name = book.Name
	}
	if book.Price != 0 {
		existingBook.Price = book.Price
	}
	if book.Details != "" {
		existingBook.Details = book.Details
	}
	existingBook.CreatedAt = time.Now()
	if err := h.repo.UpdateBook(context.Background(), &existingBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}

func (h *UserHandler) DeleteBook(c *fiber.Ctx) error {
	name := c.Params("name")
	if err := h.repo.DeleteBook(context.Background(), name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
