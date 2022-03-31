package models

type Game struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Genres []*Genre `json:"genres"`
	Rating *Rating  `json:"rating"`
}
