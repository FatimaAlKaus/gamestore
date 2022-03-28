package repository

import (
	"Lab1/pkg/models"
	"database/sql"
)

type GenreRepository interface {
	GetById(id int) (*models.Genre, error)
	GetAll() []models.Genre
	Create(p *models.Genre) (*models.Genre, error)
	Delete(id int) error
	Update(genre *models.Genre) (*models.Genre, error)
}
type Repository struct {
	Genres GenreRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Genres: newGenreRepository(db),
	}
}
