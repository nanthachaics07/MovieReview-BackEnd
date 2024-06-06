package services

import (
	"MovieReviewAPIs/handler/errs"
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

func (s *movieService) GetAllMovies(users *models.User) ([]models.Movies, error) { //fix business logic
	// Check user role
	if users.Role != nil && *users.Role != "admin" {
		return nil, errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	getAllMovieRes, err := s.MovieRepository.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return getAllMovieRes, nil
}

func (s *movieService) GetMovieEachFieldForHomePage() ([]models.MovieOnHomePage, error) {
	homePage, err := s.MovieRepository.GetMovieEachFieldForHomePage()
	if err != nil {
		return nil, err
	}
	var movieHomeRes []models.MovieOnHomePage

	for _, movies := range homePage {
		movieHomePages := models.MovieOnHomePage{
			ID:          movies.ID,
			Title:       movies.Title,
			ReleaseDate: movies.ReleaseDate,
			// Runtime:     movies.Runtime,
			MPAA:     movies.MPAA,
			ImageURL: movies.ImageURL,
		}
		movieHomeRes = append(movieHomeRes, movieHomePages)
	}

	return movieHomeRes, nil
}

func (s *movieService) GetMovieByID(users *models.User, id uint) (*models.Movies, error) {
	// Check user role
	if users.Role != nil && *users.Role != "admin" && *users.Role != "user" {
		return nil, errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	ifndMovie, err := s.MovieRepository.FindMovieByID(id)
	if err != nil {
		return nil, err
	}
	return ifndMovie, nil
}

func (s *movieService) CreateMovie(user *models.User) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	movie := new(models.Movies)
	createMovieErr := s.MovieRepository.CreateMovie(movie)
	return createMovieErr
}

func (s *movieService) UpdateMovieByID(user *models.User, id uint) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	movie := new(models.Movies)
	updateMovie := s.MovieRepository.UpdateMovieByID(movie, id)
	return updateMovie
}

func (s *movieService) DeleteMovieByID(user *models.User, id uint) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	movie := new(models.Movies)
	deleteMovie := s.MovieRepository.DeleteMovieByID(movie, id)
	return deleteMovie
}
