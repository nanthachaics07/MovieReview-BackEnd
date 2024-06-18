package test

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/models"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAdd(t *testing.T) {
	result := 5
	expected := 5
	if result != expected {
		t.Errorf("Add Fun Return %d, expected %d", result, expected)
	}
}

// func TestSubtract(t *testing.T) {

// }

func MockFindMovieByID(id uint, c *fiber.Ctx) (*models.Movies, error) {
	_, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("FindMovieByID", err.Error())
		return nil, err
	}

	// database.GetUserFromToken(token)

	var movie models.Movies
	if err := database.DB.Db.First(&movie, id).Error; err != nil {
		database.LogInfoErr("FindMovieByID", err.Error())
		return nil, err
	}

	return &movie, nil
}
