package service

import (
	"nomedoseuprojeto/internal/domain"
	"nomedoseuprojeto/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(r *repository.AuthorRepository) *AuthorService{
	return &AuthorService{repo: r}
}

func (s *AuthorService) CreateAuthor(name string) error{
	author := &domain.Author{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	return s.repo.Save(author)	
}