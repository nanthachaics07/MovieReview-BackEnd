package repositories

import (
	"MovieReviewAPIs/models"

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
		return err
	}
	return nil
}

func (r *movieRepository) GetAllMovies() ([]models.Movies, error) {
	var movies []models.Movies
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) GetMovieEachFieldForHomePage() ([]models.Movies, error) {
	var movies []models.Movies
	if err := r.db.Table("movies").Select("id, title, release_date, mpaa, image_url").Find(&movies).Error; err != nil {
		return movies, err
	}
	return movies, nil
}

func (r *movieRepository) FindMovieByID(id uint) (*models.Movies, error) {
	var movie models.Movies
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepository) UpdateMovieByID(movie *models.Movies) error {
	if err := r.db.Save(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *movieRepository) DeleteMovieByID(id uint) error {
	if err := r.db.Delete(&models.Movies{}, id).Error; err != nil {
		return err
	}
	return nil
}
