package services

import (
	"MovieReviewAPIs/repositories"

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

func (s *accountService) UsersAccountAll(c *fiber.Ctx) error {
	return s.AccountRepository.UsersAccountAll(c)
}
