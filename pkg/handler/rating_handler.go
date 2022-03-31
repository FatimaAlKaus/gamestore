package handler

import (
	"Lab1/pkg/models"
	"Lab1/pkg/repository"
	"Lab1/pkg/requests"
	"Lab1/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RatingHandler struct {
	rep *repository.Repository
}

func (r *RatingHandler) CreateRating(ctx *gin.Context) {
	var input requests.CreateRatingInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.BadRequest(ctx, "invalid request")
		return
	}
	rating := &models.Rating{Name: input.Name}
	result, err := r.rep.Ratings.Create(rating)
	if err != nil {
		response.Conflict(ctx, "rating with such name already exists")
		return
	}
	response.Created(ctx, result)
}

func (r *RatingHandler) GetAllRatings(ctx *gin.Context) {
	var ratings []*models.RatingDto
	for _, rating := range r.rep.Ratings.GetAll() {
		ratings = append(ratings, rating.ToDTO())
	}
	if len(ratings) == 0 {
		response.NotFound(ctx, "")
		return
	}
	response.OK(ctx, ratings)
}

func (r *RatingHandler) GetRatingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Specified incorrect index"})
		return
	}
	data, err := r.rep.Ratings.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data.ToDTO()})
}

func newRatingHandler(rep *repository.Repository) *RatingHandler {
	return &RatingHandler{rep: rep}
}
