package models

type Rating struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Games []*Game `json:"games"`
}
type RatingDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Games []int  `json:"games"`
}

func (r *Rating) ToDTO() *RatingDto {
	var games []int
	for _, game := range r.Games {
		games = append(games, game.Id)
	}
	return &RatingDto{
		Id:    r.Id,
		Name:  r.Name,
		Games: games,
	}
}
