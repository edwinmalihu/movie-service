package repository

import (
	"log"
	"movie-service/model"
	"movie-service/request"

	"gorm.io/gorm"
)

type MovieRepo interface {
	Migrate() error
	AddMovie(request.MovieRequest) (model.Movie, error)
	DetailMovie(int) (model.Movie, error)
	UpdateMovie(request.UpdateMovieRequest) (model.Movie, error)
	ListMovie() ([]model.Movie, error)
	DeleteMovie(int) (model.Movie, error)
}

type movierepo struct {
	DB *gorm.DB
}

// DeleteMovie implements MovieRepo.
func (m movierepo) DeleteMovie(req int) (data model.Movie, err error) {
	return data, m.DB.Where("id = ?", req).Delete(&data).Error
}

// AddMovie implements MovieRepo.
func (m movierepo) AddMovie(req request.MovieRequest) (data model.Movie, err error) {
	data = model.Movie{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}

	return data, m.DB.Create(&data).Error
}

// DetailMovie implements MovieRepo.
func (m movierepo) DetailMovie(id int) (data model.Movie, err error) {
	return data, m.DB.First(&data, "id = ?", id).Error
}

// ListMovie implements MovieRepo.
func (m movierepo) ListMovie() (data []model.Movie, err error) {
	return data, m.DB.Find(&data).Error
}

// UpdateMovie implements MovieRepo.
func (m movierepo) UpdateMovie(req request.UpdateMovieRequest) (data model.Movie, err error) {
	return data, m.DB.Model(&data).Where("id = ?", req.Id).Updates(model.Movie{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}).Error
}

// Migrate implements MovieRepo.
func (m movierepo) Migrate() error {
	log.Print("[movie db migrate]")
	return m.DB.AutoMigrate(&model.Movie{})
}

func NewMovieRepo(db *gorm.DB) MovieRepo {
	return movierepo{
		DB: db,
	}
}
