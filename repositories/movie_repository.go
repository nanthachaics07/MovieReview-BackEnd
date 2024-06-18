package repositories

import (
	"MovieReviewAPIs/database"

	// "MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"

	// "reflect"

	// "github.com/dgrijalva/jwt-go"
	// "github.com/gofiber/fiber/v2"

	// jwt "github.com/golang-jwt/jwt"

	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) CreateMovie(movie *models.Movies) error {
	if err := r.db.Create(movie).Error; err != nil {
		database.LogInfoErr("CreateMovie", err.Error())
		return err
	}
	return nil
}

func (r *movieRepository) GetAllMovies() ([]models.Movies, error) {
	var movies []models.Movies
	// var c *fiber.Ctx
	// _, err := middlewares.VerifyAuth(c)
	// if err != nil {
	// 	database.LogInfoErr("GetAllMovies", err.Error())
	// 	return nil, err
	// }

	if err := r.db.Find(&movies).Error; err != nil {
		database.LogInfoErr("GetAllMovies", err.Error())
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) GetMovieEachFieldForHomePage() ([]models.MovieOnHomePage, error) {
	var movies []models.MovieOnHomePage
	if err := r.db.Table("movies").Select("id, title, release_date, mpaa, image_url").Find(&movies).Error; err != nil {
		database.LogInfoErr("GetMovieEachFieldForHomePage", err.Error())
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) FindMovieByID(id uint) (*models.Movies, error) {

	// var movie []models.Movies
	// if err := r.db.Table("movies").Select("id, title, release_date, runtime, description, mpaa, image_url").Find(&movie, id).Error; err != nil {
	// 	database.LogInfoErr("FindMovieByID", err.Error())
	// 	return nil, err
	// }
	var movie models.Movies
	if err := r.db.First(&movie, id).Error; err != nil {
		database.LogInfoErr("FindMovieByID", err.Error())
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) UpdateMovieByID(movie *models.Movies, id uint) error {
	movie.ID = id
	// Use Model
	if err := r.db.Model(&models.Movies{}).Where("id = ?", id).Updates(movie).Error; err != nil {
		database.LogInfoErr("UpdateMovieByID", err.Error())
		return err
	}
	return nil
}

func (r *movieRepository) DeleteMovieByID(movie *models.Movies, id uint) error {
	movie.ID = id
	if err := r.db.Delete(&models.Movies{}).Where("id = ?", id).Error; err != nil {
		database.LogInfoErr("DeleteMovieByID", err.Error())
		return err
	}
	return nil
}
