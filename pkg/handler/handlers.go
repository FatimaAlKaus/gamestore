package handler

import (
	"Lab1/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Genres
}
type Genres interface {
	GetAllGenres(ctx *gin.Context)
	GetGenreById(ctx *gin.Context)
	CreateGenre(ctx *gin.Context)
	DeleteGenre(ctx *gin.Context)
	UpdateGenre(ctx *gin.Context)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		genres := api.Group("/genres")
		{
			genres.GET("/", h.GetAllGenres)
			genres.GET("/:id", h.GetGenreById)
			genres.POST("/", h.CreateGenre)
			genres.DELETE("/:id", h.DeleteGenre)
			genres.PUT("/", h.UpdateGenre)
		}
	}

	return router
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{
		Genres: newGenreHandler(repos),
	}
}
