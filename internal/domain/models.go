package domain

import "time"

type AuthorInput struct {
	Name string `json:"name"`
}

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}


type BookInput struct {
	AuthorID      string `json:"author_id"`
	Title         string `json:"title"`
	PublishedYear int    `json:"published_year"`
}

type Book struct {
	ID            string    `json:"id"`
	AuthorID      string    `json:"author_id"`
	Title         string    `json:"title"`
	PublishedYear int       `json:"published_year"`
	CreatedAt     time.Time `json:"created_at"`
}
