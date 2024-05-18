package services

import (
	"MovieReviewAPIs/models"
)

type MovieService interface {
	GetAllMovies(user *models.User) ([]models.Movies, error)
	GetMovieByID(user *models.User, id uint) (*models.Movies, error)
	GetMovieEachFieldForHomePage() ([]models.MovieOnHomePage, error)
	CreateMovie(user *models.User) error
	UpdateMovieByID(user *models.User, id uint) error
	DeleteMovieByID(user *models.User, id uint) error
}
