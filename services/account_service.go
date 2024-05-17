package services

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/repositories"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type accountService struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountService(accountRepositoty repositories.AccountRepository) *accountService {
	return &accountService{AccountRepository: accountRepositoty}
}

func (s *accountService) UserAccount(c *fiber.Ctx) error {
	return s.AccountRepository.UserAccount(c)
}

func (s *accountService) UsersAccountAll(c *fiber.Ctx, user *models.User) ([]models.User, error) {
	if *user.Role != "admin" {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UsersAccountAll", "unauthorized")
		return nil, errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}
	allUser, err := s.AccountRepository.UsersAccountAll(c)
	if err != nil {
		return nil, err
	}
	return allUser, nil
}

func (s *accountService) GetUserByID(c *fiber.Ctx, user *models.User, id uint) (*models.User, error) {
	if user.Role != nil && *user.Role != "admin" {
		return nil, errors.New("unauthorized user role!! WHO ARE U?")
	}
	getUserbyid, err := s.AccountRepository.GetuserByID(c, id)
	if err != nil {
		return nil, err
	}

	return getUserbyid, nil
}
