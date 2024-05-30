package repositories

import (
	"context"
	"fmt"
	"supabase-fiber-SupaDB-project/internal/models"

	"github.com/nedpals/supabase-go"
)

type UserRepository struct {
	client *supabase.Client
}

func NewBookRepository(url, key string) *UserRepository {
	client := supabase.CreateClient(url, key)
	return &UserRepository{client: client}
}

func (r *UserRepository) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var users []models.Book
	err := r.client.DB.From("books").Select("*").Execute(&users)
	fmt.Println(users)

	return users, err
}

func (r *UserRepository) GetBookByName(ctx context.Context, name string) (models.Book, error) {
	var book models.Book
	err := r.client.DB.From("books").Select("*").Single().Eq("name", name).Execute(&book)

	fmt.Println("book at repo", book)

	return book, err
}

func (r *UserRepository) CreateBook(ctx context.Context, book *models.Book) error {
	// var results map[string]interface{}
	fmt.Println("book at createBook", book)
	err := r.client.DB.From("books").Insert(book).Execute(nil)
	return err
}

func (r *UserRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	fmt.Println("update book func in repo", book)

	err := r.client.DB.From("books").Update(book).Eq("uid", book.UID.String()).Execute(nil)
	return err
}

func (r *UserRepository) DeleteBook(ctx context.Context, name string) error {
	err := r.client.DB.From("books").Delete().Eq("name", name).Execute(nil)
	return err
}
