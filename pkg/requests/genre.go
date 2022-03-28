package requests

type CreateGenreInput struct {
	Name string `json:"name"`
}
type UpdateGenreInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
