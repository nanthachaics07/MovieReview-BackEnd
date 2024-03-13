package handler

import (
	"MovieReviewAPIs/services"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (h *AccountHandler) UserAccountHandler(c *fiber.Ctx) error {
	payload := h.AccountService.UserAccount(c)

	if payload != nil {
		return c.JSON(payload)
	}

	return nil
}

func (h *AccountHandler) UsersAccountAllHandler(c *fiber.Ctx) error {
	payload := h.AccountService.UsersAccountAll(c)

	if payload != nil {
		return c.JSON(payload)
	}

	return nil
}
