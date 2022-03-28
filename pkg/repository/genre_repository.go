package repository

import (
	"Lab1/pkg/models"
	"database/sql"
	"errors"
)

type Genre struct {
	db *sql.DB
}

func newGenreRepository(db *sql.DB) *Genre {
	return &Genre{db: db}
}
func (g *Genre) GetById(id int) (*models.Genre, error) {
	genre := models.Genre{}
	if err := g.db.QueryRow("SELECT id, name FROM genres WHERE id= $1", id).
		Scan(&genre.Id, &genre.Name); err != nil {
		return nil, err
	}
	return &genre, nil
}

func (g *Genre) GetAll() []models.Genre {
	rows, err := g.db.Query("select * from genres")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var genres []models.Genre

	for rows.Next() {
		genre := models.Genre{}
		_ = rows.Scan(&genre.Id, &genre.Name)
		genres = append(genres, genre)
	}
	return genres
}
func (g *Genre) Create(genre *models.Genre) (*models.Genre, error) {
	var id int
	if err := g.db.QueryRow("INSERT INTO genres (name) VALUES ($1) RETURNING id", genre.Name).
		Scan(&id); err != nil {
		return nil, err
	}
	model, err := g.GetById(id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (g *Genre) Delete(id int) error {
	result, err := g.db.Exec("DELETE FROM genres WHERE id=$1", id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows != 1 {
		return errors.New("specified id wasn't found")
	}
	return nil
}

func (g *Genre) Update(genre *models.Genre) (*models.Genre, error) {
	result, err := g.db.Exec("UPDATE  genres SET name = $1 WHERE id = $2", genre.Name, genre.Id)
	if err != nil {
		return nil, err
	}
	if rows, _ := result.RowsAffected(); rows != 1 {
		return nil, errors.New("specified id wasn't found")
	}
	model, _ := g.GetById(genre.Id)
	return model, err
}
