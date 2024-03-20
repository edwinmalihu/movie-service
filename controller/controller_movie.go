package controller

import (
	"log"
	"movie-service/repository"
	"movie-service/request"
	"movie-service/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	AddMovie(*gin.Context)
	DetailMovie(*gin.Context)
	UpdateMovie(*gin.Context)
	ListMovie(*gin.Context)
	DeteleMovie(*gin.Context)
}

type movieController struct {
	MovieRepo repository.MovieRepo
}

// AddMovie implements MovieController.
func (m movieController) AddMovie(ctx *gin.Context) {
	var req request.MovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong field"})
		return
	}

	data, err := m.MovieRepo.AddMovie(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "list failed to response"})
		return
	}

	result := response.MovieResponse{
		Id:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, result)

}

// DetailMovie implements MovieController.
func (m movieController) DetailMovie(ctx *gin.Context) {
	var req request.MovieId

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong field"})
		return
	}

	id, _ := strconv.Atoi(req.Id)

	data, err := m.MovieRepo.DetailMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "list failed to response"})
		return
	}

	result := response.MovieResponse{
		Id:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, result)

}

// DeteleMovie implements MovieController.
func (m movieController) DeteleMovie(ctx *gin.Context) {
	var req request.MovieId

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong field"})
		return
	}

	id, _ := strconv.Atoi(req.Id)

	data, err := m.MovieRepo.DeleteMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "list failed to response"})
		return
	}

	result := "delete movie success"
	if data.ID != 0 {
		result = "delete movie failed"
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})

}

// ListMovie implements MovieController.
func (m movieController) ListMovie(ctx *gin.Context) {
	var result []response.MovieResponse
	data, err := m.MovieRepo.ListMovie()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "list movie failed to response"})
		return
	}

	for _, list := range data {
		row := response.MovieResponse{
			Id:          list.ID,
			Title:       list.Title,
			Description: list.Description,
			Rating:      list.Rating,
			Image:       list.Image,
			CreatedAt:   list.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   list.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		result = append(result, row)
	}

	ctx.JSON(http.StatusOK, result)

}

// UpdateMovie implements MovieController.
func (m movieController) UpdateMovie(ctx *gin.Context) {
	string_id := ctx.Param("id")
	id, _ := strconv.Atoi(string_id)
	var req request.MovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong field"})
		return
	}

	data, err := m.MovieRepo.UpdateMovie(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed update movie"})
		return
	}

	log.Println("[Update query] = ", data)

	res, err := m.MovieRepo.DetailMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed update movie"})
		return
	}

	result := response.MovieResponse{
		Id:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Rating:      res.Rating,
		Image:       res.Image,
		CreatedAt:   res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   res.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, result)

}

func NewMovieController(repo repository.MovieRepo) MovieController {
	return movieController{
		MovieRepo: repo,
	}
}
