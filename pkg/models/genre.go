package models

type Genre struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Games []*Game `json:"games"`
}
