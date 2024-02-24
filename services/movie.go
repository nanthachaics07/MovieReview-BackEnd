package services

import "MovieReviewAPIs/models"

type MovieService interface {
	GetAllMovies() ([]models.Movies, error)
	GetMovieByID(id uint) (*models.Movies, error)
	CreateMovie(movie *models.Movies) error
	UpdateMovie(movie *models.Movies) error
	DeleteMovieByID(id uint) error
}
