package repository

import (
	"database/sql"
	"nomedoseuprojeto/internal/domain"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository{
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Save(author *domain.Author) error{
	stmt, err := r.db.Prepare(`
		INSERT INTO authors (id, name, created_at)
		VALUES ($1, $2, $3)
	`)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		author.ID, 
		author.Name, 
		author.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}