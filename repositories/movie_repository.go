package repositories

import (
	"MovieReviewAPIs/model"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (repo *MovieRepository) GetAllMovies() []model.Movie {
	var movies []model.Movie
	repo.DB.Find(&movies)
	return movies
}

func (repo *MovieRepository) GetMovieById(id uint) (*model.Movie, error) {
	var movie model.Movie
	err := repo.DB.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}
