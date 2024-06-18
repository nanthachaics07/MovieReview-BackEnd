package test

// import (
// 	"MovieReviewAPIs/models"
// 	"errors"

// 	"gorm.io/gorm"
// )

// type mockMovieRepository struct {
// 	db *gorm.DB
// }

// func (m *mockMovieRepository) CreateMovie(movie *models.Movies) error {

// 	return nil
// }

// // GetAllMovies mocks the GetAllMovies method of MovieRepository
// func (m *mockMovieRepository) GetAllMovies() ([]models.Movies, error) {

// 	return []models.Movies{
// 		{
// 			// Model:         models.Movies{ID: 1},
// 			ID:            1,
// 			Title:         "Movie 1",
// 			ReleaseDate:   "2023-01-01",
// 			Runtime:       "120 mins",
// 			Rating:        "8.5",
// 			Category:      "Action",
// 			Popularity:    "High",
// 			Budget:        10000000,
// 			Revenue:       50000000,
// 			Director:      "Director 1",
// 			Casting:       "Actor 1, Actress 1",
// 			Writers:       "Writer 1",
// 			DistributedBy: "Company 1",
// 			MPAA:          "PG-13",
// 			Description:   "Description of Movie 1",
// 			ImageURL:      "http://example.com/movie1.jpg",
// 		},
// 		{
// 			// Model:         models.Model{ID: 2},
// 			ID:            2,
// 			Title:         "Movie 2",
// 			ReleaseDate:   "2023-02-01",
// 			Runtime:       "110 mins",
// 			Rating:        "7.9",
// 			Category:      "Comedy",
// 			Popularity:    "Medium",
// 			Budget:        8000000,
// 			Revenue:       30000000,
// 			Director:      "Director 2",
// 			Casting:       "Actor 2, Actress 2",
// 			Writers:       "Writer 2",
// 			DistributedBy: "Company 2",
// 			MPAA:          "PG",
// 			Description:   "Description of Movie 2",
// 			ImageURL:      "http://example.com/movie2.jpg",
// 		},
// 	}, nil
// }

// // GetMovieEachFieldForHomePage mocks the GetMovieEachFieldForHomePage method of MovieRepository
// // func (m *mockMovieRepository) GetMovieEachFieldForHomePage() ([]models.MovieOnHomePage, error) {
// // 	//TODO!! fix json type #############
// // 	return []models.MovieOnHomePage{
// // 		{ID: 1, Title: "Movie 1", Genre: "Action"},
// // 		{ID: 2, Title: "Movie 2", Genre: "Comedy"},
// // 	}, nil
// // }

// // FindMovieByID mocks the FindMovieByID method of MovieRepository
// func (m *mockMovieRepository) FindMovieByID(id uint) (*models.Movies, error) {

// 	if id == 1 {
// 		return &models.Movies{
// 			// Model:         models.Model{ID: 1},
// 			ID:            1,
// 			Title:         "Movie 1",
// 			ReleaseDate:   "2023-01-01",
// 			Runtime:       "120 mins",
// 			Rating:        "8.5",
// 			Category:      "Action",
// 			Popularity:    "High",
// 			Budget:        10000000,
// 			Revenue:       50000000,
// 			Director:      "Director 1",
// 			Casting:       "Actor 1, Actress 1",
// 			Writers:       "Writer 1",
// 			DistributedBy: "Company 1",
// 			MPAA:          "PG-13",
// 			Description:   "Description of Movie 1",
// 			ImageURL:      "http://example.com/movie1.jpg",
// 		}, nil
// 	} else if id == 2 {
// 		return &models.Movies{
// 			// Model:         models.Model{ID: 2},
// 			ID:            2,
// 			Title:         "Movie 2",
// 			ReleaseDate:   "2023-02-01",
// 			Runtime:       "110 mins",
// 			Rating:        "7.9",
// 			Category:      "Comedy",
// 			Popularity:    "Medium",
// 			Budget:        8000000,
// 			Revenue:       30000000,
// 			Director:      "Director 2",
// 			Casting:       "Actor 2, Actress 2",
// 			Writers:       "Writer 2",
// 			DistributedBy: "Company 2",
// 			MPAA:          "PG",
// 			Description:   "Description of Movie 2",
// 			ImageURL:      "http://example.com/movie2.jpg",
// 		}, nil
// 	}
// 	return nil, errors.New("movie not found")
// }

// // UpdateMovieByID mocks the UpdateMovieByID method of MovieRepository
// func (m *mockMovieRepository) UpdateMovieByID(movie *models.Movies, id uint) error {

// 	if id != movie.ID {
// 		return errors.New("IDs do not match")
// 	}
// 	// Update logic...
// 	return nil
// }

// // DeleteMovieByID mocks the DeleteMovieByID method of MovieRepository
// func (m *mockMovieRepository) DeleteMovieByID(id uint) error {
// 	if id == 1 || id == 2 {
// 		return nil
// 	}
// 	return errors.New("movie not found")
// }

// // func Test_movieRepository_GetMovieEachFieldForHomePage(t *testing.T) {
// // 	type fields struct {
// // 		db *gorm.DB
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		want    []models.MovieOnHomePage
// // 		wantErr bool
// // 	}{
// // 		// Add test cases here
// // 		{
// // 			name: "Test case 1",
// // 			fields: fields{
// // 				// Initialize with a GORM.DB instance
// // 				db: nil,
// // 			},
// // 			want:    nil,
// // 			wantErr: false,
// // 		},
// // 		// Add more test

// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			r := &mockMovieRepository{
// // 				db: tt.fields.db,
// // 			}
// // 			got, err := r.GetMovieEachFieldForHomePage()
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("movieRepository.GetMovieEachFieldForHomePage() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("movieRepository.GetMovieEachFieldForHomePage() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }
