package models

type Rating struct {
	id    int
	name  string
	games []*Game
}
