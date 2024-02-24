package services

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/repositories"
)

type movieService struct {
	MovieRepository repositories.MovieRepository
}

func NewMovieService(movieRepository repositories.MovieRepository) *movieService {
	return &movieService{
		MovieRepository: movieRepository,
	}
}

func (s *movieService) GetAllMovies() ([]models.Movies, error) {
	return s.MovieRepository.GetAllMovies()
}

func (s *movieService) GetMovieByID(id uint) (*models.Movies, error) {
	return s.MovieRepository.FindMovieByID(id)
}

func (s *movieService) CreateMovie(movie *models.Movies) error {
	return s.MovieRepository.CreateMovie(movie)
}

func (s *movieService) UpdateMovie(movie *models.Movies) error {
	return s.MovieRepository.UpdateMovieByID(movie)
}

func (s *movieService) DeleteMovieByID(id uint) error {
	return s.MovieRepository.DeleteMovieByID(id)
}
