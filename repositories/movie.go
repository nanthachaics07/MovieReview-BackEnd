package repositories

import "MovieReviewAPIs/models"

type MovieRepository interface {
	CreateMovie(movie *models.Movies) error
	GetAllMovies() ([]models.Movies, error)
	GetMovieEachFieldForHomePage() ([]models.Movies, error)
	FindMovieByID(id uint) (*models.Movies, error)
	UpdateMovieByID(movie *models.Movies, id uint) error
	DeleteMovieByID(id uint) error
}
