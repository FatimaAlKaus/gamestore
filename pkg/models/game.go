package models

type Game struct {
	id     int
	name   string
	genres []*Genre
	rating *Rating
}
