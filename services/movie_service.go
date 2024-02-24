package services

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/repositories"
)

type MovieService struct {
	MovieRepository *repositories.MovieRepository
}

func NewMovieService(movieRepository *repositories.MovieRepository) *MovieService {
	return &MovieService{
		MovieRepository: movieRepository,
	}
}

func (s *MovieService) GetAllMovies() ([]models.Movies, error) {
	return s.MovieRepository.GetAllMovies()
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movies, error) {
	return s.MovieRepository.FindMovieByID(id)
}

func (s *MovieService) CreateMovie(movie *models.Movies) error {
	return s.MovieRepository.CreateMovie(movie)
}

func (s *MovieService) UpdateMovie(movie *models.Movies) error {
	return s.MovieRepository.UpdateMovieByID(movie)
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.MovieRepository.DeleteMovieByID(id)
}
