package handler

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (h *AccountHandler) UserAccountHandler(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("GetuserByID", "unauthenticated")
		return err
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}

	payload, err := h.AccountService.UserAccount(c, user)
	if err != nil {
		return err
	}

	return c.JSON(payload)
}

func (h *AccountHandler) UsersAccountAllHandler(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UsersAccountAll", "unauthenticated")
		return err
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	fmt.Println(user.Role)

	payload, err := h.AccountService.UsersAccountAll(c, user)
	if err != nil {
		return err
	}

	return c.JSON(payload)
}

func (h *AccountHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("GetuserByID", "unauthenticated")
		return err
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}

	idStr := c.Params("id")
	newID, newErr := strconv.ParseUint(idStr, 10, 0)
	if newErr != nil {
		return errs.NewUnexpectedError(newErr.Error())
	}
	id := uint(newID)

	payload, err := h.AccountService.GetUserByID(c, user, id)
	if err != nil {
		return err
	}
	return c.JSON(payload)
}
