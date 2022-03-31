package repository

import (
	"Lab1/pkg/models"
	"database/sql"
	"errors"
)

type Rating struct {
	db *sql.DB
}

func (r *Rating) GetById(id int) (*models.Rating, error) {
	query := ` 
SELECT r.id, r.name, g.id, g.name
    FROM rating AS r
    LEFT JOIN games AS g ON r.id = g.rating_id
    WHERE r.id = $1
    `
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	rating := &models.Rating{}
	for rows.Next() {
		game := &models.Game{}
		err = rows.Scan(
			&rating.Id,
			&rating.Name,
			&game.Id,
			&game.Name,
		)
		if err != nil {
			continue
		}
		rating.Games = append(rating.Games, game)
	}
	if rating.Id == 0 {
		return nil, errors.New("not found")
	}
	return rating, nil
}

func (r *Rating) GetAll() []models.Rating {
	query := `SELECT id, name FROM rating`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil
	}
	var ratings []models.Rating
	for rows.Next() {
		ra := models.Rating{}
		rows.Scan(&ra.Id, &ra.Name)
		ratings = append(ratings, ra)
	}
	return ratings
}

func (r *Rating) Create(input *models.Rating) (*models.Rating, error) {
	var id int
	if err := r.db.QueryRow("INSERT INTO rating (name) VALUES ($1) RETURNING id", input.Name).
		Scan(&id); err != nil {
		return nil, err
	}
	return r.GetById(id)
}

func (r *Rating) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *Rating) Update(input *models.Rating) (*models.Rating, error) {
	//TODO implement me
	panic("implement me")
}

func newRatingRepository(db *sql.DB) *Rating {
	return &Rating{db: db}
}
