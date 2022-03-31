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
type RatingRepository interface {
	GetById(id int) (*models.Rating, error)
	GetAll() []models.Rating
	Create(p *models.Rating) (*models.Rating, error)
	Delete(id int) error
	Update(genre *models.Rating) (*models.Rating, error)
}
type GameRepository interface {
	GetById(id int) (*models.Game, error)
	GetAll() []models.Game
	Create(p *models.Game) (*models.Game, error)
	Delete(id int) error
	Update(genre *models.Game) (*models.Game, error)
}
type Repository struct {
	Genres         GenreRepository
	Ratings        RatingRepository
	GameRepository GameRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Genres:  newGenreRepository(db),
		Ratings: newRatingRepository(db),
	}
}
