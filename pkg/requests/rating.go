package requests

type CreateRatingInput struct {
	Name string `json:"name"`
}
type UpdateRatingInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
