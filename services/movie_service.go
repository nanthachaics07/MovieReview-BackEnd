package services

import (
	"MovieReviewAPIs/model"
	"MovieReviewAPIs/repositories"
)

type MovieService struct {
	MovieRepository *repositories.MovieRepository
}

func NewMovieService(movieRepository *repositories.MovieRepository) *MovieService {
	return &MovieService{MovieRepository: movieRepository}
}

func (service *MovieService) GetAllMovies() []model.Movie {
	return service.MovieRepository.GetAllMovies()
}

func (service *MovieService) GetMovieById(id uint) (*model.Movie, error) {
	return service.MovieRepository.GetMovieById(id)
}
