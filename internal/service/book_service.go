package service

import (
	"nomedoseuprojeto/internal/domain"
	"nomedoseuprojeto/internal/repository"
	"time"

	"github.com/google/uuid"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService{//n fiz ainda
	return &BookService{repo: r}
}

func (s *BookService) CreateBook(input domain.BookInput) error{
	book := &domain.Book{
		ID:            uuid.NewString(),
		AuthorID:      input.AuthorID,
		Title:         input.Title,
		PublishedYear: input.PublishedYear,
		CreatedAt:     time.Now(),
	}
	return s.repo.Save(book)	
}