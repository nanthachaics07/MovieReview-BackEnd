package handler

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	MovieService services.MovieService
}

func NewMovieHandler(movieService services.MovieService) *MovieHandler {
	return &MovieHandler{
		MovieService: movieService,
	}
}

func (h *MovieHandler) GetAllMovies(c *fiber.Ctx) error { //TODO: fix business logic
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("GetAllMovies", err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	// if *user.Role != "admin" {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	database.LogInfoErr("GetAllMovies", "unauthorized")
	// 	return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	// }

	movies, err := h.MovieService.GetAllMovies(user)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movies)
}

func (h *MovieHandler) GetMovieByID(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("GetMovieByID", err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}

	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	movie, err := h.MovieService.GetMovieByID(user, uint(id))
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movie)
}

func (h *MovieHandler) GetMovieForHomePage(c *fiber.Ctx) error {
	// payload := new(models.MovieOnHomePage)
	movies, err := h.MovieService.GetMovieEachFieldForHomePage()
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movies)
}

func (h *MovieHandler) CreateMovie(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("CreateMovie", err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	// if *user.Role != "admin" {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	database.LogInfoErr("CreateMovie", "unauthorized")
	// 	return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	// }

	// movie := new(models.Movies)
	// if err := c.BodyParser(movie); err != nil {
	// 	return errs.NewUnexpectedError(err.Error())
	// }
	if err := h.MovieService.CreateMovie(user); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(fiber.Map{"status": "success",
		"message": "Movie created successfully"})
}

func (h *MovieHandler) UpdateMovieByID(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("UpdateMovie", err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	// if *user.Role != "admin" {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	database.LogInfoErr("UpdateMovie", "unauthorized")
	// 	return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	// }

	// movie := new(models.Movies)
	// if err := c.BodyParser(movie); err != nil {
	// 	return errs.NewUnexpectedError(err.Error())
	// }
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	if err := h.MovieService.UpdateMovieByID(user, uint(id)); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(fiber.Map{"status": "success",
		"message": "Movie updated successfully"})
}

func (h *MovieHandler) DeleteMovie(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		database.LogInfoErr("DeleteMovie", err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	// if *user.Role != "admin" {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	database.LogInfoErr("DeleteMovie", "unauthorized")
	// 	return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	// }

	// movie := new(models.Movies)
	// if err := c.BodyParser(movie); err != nil {
	// 	return errs.NewUnexpectedError(err.Error())
	// }
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	if err := h.MovieService.DeleteMovieByID(user, uint(id)); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(fiber.Map{"status": "success",
		"message": "Movie deleted successfully"})
}
