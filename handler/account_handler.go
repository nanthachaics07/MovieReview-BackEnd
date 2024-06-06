package handler

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
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

func (h *AccountHandler) UpdateUserHandler(c *fiber.Ctx) error {

	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("UpdateUserHandler", "unauthenticated")
		return errs.NewUnauthorizedError("unauthenticated")
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		database.LogInfoErr("UpdateUserHandler", "failed to get user from token")
		return errs.NewInternalServerError("failed to get user from token")
	}

	payload := new(models.UserUpdate)
	if err := c.BodyParser(payload); err != nil {
		database.LogInfoErr("UpdateUserHandler", "failed to parse request body")
		return errs.NewBadRequestError("failed to parse request body")
	}

	if err := h.AccountService.UpdateUserByID(c, user, payload); err != nil {
		database.LogInfoErr("UpdateUserHandler", "failed to update user")
		return errs.NewInternalServerError("failed to update user")
	}

	return c.JSON(fiber.Map{"status": "success", "message": "user updated successfully"})
}

func (h *AccountHandler) DeleteUserHandler(c *fiber.Ctx) error {
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

	if err := h.AccountService.DeleteUserByID(c, user, id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success"})
}
