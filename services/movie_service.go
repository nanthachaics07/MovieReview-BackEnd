package services

import (
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/repositories"
	"errors"

	"gorm.io/gorm"
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

	// var movieByidRes []models.Movies

	// for _, movie := range ifndMovie {
	// 	if movie.ID == id {
	// 		movieByID := models.Movies{
	// 			ID:          movie.ID,
	// 			Title:       movie.Title,
	// 			ReleaseDate: movie.ReleaseDate,
	// 			Runtime:     movie.Runtime,
	// 			MPAA:        movie.MPAA,
	// 			Description: movie.Description,
	// 			ImageURL:    movie.ImageURL,
	// 		}
	// 		movieByidRes = append(movieByidRes, movieByID)
	// 	}
	// }

	return ifndMovie, nil
}

func (s *movieService) CreateMovie(user *models.User, movie *models.Movies) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	// movie := new(models.Movies)

	if movie.Title == "" || movie.ReleaseDate == "" || movie.Runtime == "" || movie.MPAA == "" || movie.Description == "" || movie.ImageURL == "" {
		return errs.NewBadRequestError("invalid movie data")
	}

	createMovieErr := s.MovieRepository.CreateMovie(movie)
	return createMovieErr
}

func (s *movieService) UpdateMovieByID(user *models.User, id uint, movie *models.Movies) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	// movie := new(models.Movies)

	var updateMovie models.Movies
	if _, err := s.MovieRepository.FindMovieByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("movie not found")
		}
		return err
	}

	// updateMovie.ID = id
	updateMovie.Title = movie.Title
	updateMovie.ReleaseDate = movie.ReleaseDate
	updateMovie.Runtime = movie.Runtime
	updateMovie.MPAA = movie.MPAA
	updateMovie.Description = movie.Description
	updateMovie.ImageURL = movie.ImageURL

	if err := s.MovieRepository.UpdateMovieByID(&updateMovie, id); err != nil {
		return err
	}
	return nil
}

func (s *movieService) DeleteMovieByID(user *models.User, id uint) error {
	if user.Role != nil && *user.Role != "admin" {
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	movie := new(models.Movies)
	deleteMovie := s.MovieRepository.DeleteMovieByID(movie, id)
	return deleteMovie
}
