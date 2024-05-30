package main

import (
	"log"
	"os"
	"supabase-fiber-SupaDB-project/internal/handlers"
	"supabase-fiber-SupaDB-project/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := godotenv.Load("Environment.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("Supabase URL and Key must be set in .env file")
	}

	app := fiber.New()

	userRepo := repositories.NewBookRepository(supabaseURL, supabaseKey)
	userHandler := handlers.NewUserHandler(userRepo)

	app.Get("/books", userHandler.GetAllBooks)
	app.Get("/books/:name", userHandler.GetBookByName)
	app.Post("/books", userHandler.CreateBook)
	app.Put("/books/:name", userHandler.UpdateBook)
	app.Delete("/books/:name", userHandler.DeleteBook)

	log.Fatal(app.Listen(":3000"))
}
