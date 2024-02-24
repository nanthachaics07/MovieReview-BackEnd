package repositories

import (
	"MovieReviewAPIs/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) CreateMovie(movie *models.Movies) error {
	if err := r.db.Create(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) GetAllMovies() ([]models.Movies, error) {
	var movies []models.Movies
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) FindMovieByID(id uint) (*models.Movies, error) {
	var movie models.Movies
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) UpdateMovieByID(movie *models.Movies) error {
	if err := r.db.Save(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) DeleteMovieByID(id uint) error {
	if err := r.db.Delete(&models.Movies{}, id).Error; err != nil {
		return err
	}
	return nil
}
