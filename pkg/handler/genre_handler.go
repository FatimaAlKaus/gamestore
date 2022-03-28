package handler

import (
	"Lab1/pkg/models"
	"Lab1/pkg/repository"
	"Lab1/pkg/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GenreHandler struct {
	rep *repository.Repository
}

func newGenreHandler(rep *repository.Repository) *GenreHandler {
	return &GenreHandler{rep: rep}
}

func (g GenreHandler) GetAllGenres(ctx *gin.Context) {
	genres := g.rep.Genres.GetAll()
	ctx.JSON(http.StatusOK, gin.H{"data": genres})
}

func (g GenreHandler) GetGenreById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Specified incorrect index"})
		return
	}
	genre, err := g.rep.Genres.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Genre with such is wasn't found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": genre})
}

func (g *GenreHandler) CreateGenre(ctx *gin.Context) {
	var input requests.CreateGenreInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre := &models.Genre{Name: input.Name}
	result, err := g.rep.Genres.Create(genre)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Same genre already exists"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": result})
}

func (g *GenreHandler) DeleteGenre(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Specified incorrect index"})
		return
	}
	if err := g.rep.Genres.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Genre with such is wasn't found"})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
func (g *GenreHandler) UpdateGenre(ctx *gin.Context) {
	var input requests.UpdateGenreInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre := &models.Genre{Id: input.Id, Name: input.Name}
	result, err := g.rep.Genres.Update(genre)
	// TODO: we need to handle err to return associated response
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": result})
}
