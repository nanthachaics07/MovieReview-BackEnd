package handler

import (
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
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

// func (h *MovieHandler) GetAllMovies(c *fiber.Ctx) error {
// 	movies, err := h.MovieService.GetAllMovies()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(movies)
// }

func (h *MovieHandler) GetAllMovies(c *fiber.Ctx) error {
	movies, err := h.MovieService.GetAllMovies()
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movies)
}

func (h *MovieHandler) GetMovieByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	movie, err := h.MovieService.GetMovieByID(uint(id))
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movie)
}

func (h *MovieHandler) GetMovieForHomePage(c *fiber.Ctx) error {
	movies, err := h.MovieService.GetMovieEachFieldForHomePage()
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movies)
}

func (h *MovieHandler) CreateMovie(c *fiber.Ctx) error {
	movie := new(models.Movies)
	if err := c.BodyParser(movie); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	if err := h.MovieService.CreateMovie(movie); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movie)
}

func (h *MovieHandler) UpdateMovie(c *fiber.Ctx) error {
	movie := new(models.Movies)
	if err := c.BodyParser(movie); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	if err := h.MovieService.UpdateMovie(movie); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movie)
}

func (h *MovieHandler) DeleteMovie(c *fiber.Ctx) error {
	movie := new(models.Movies)
	if err := c.BodyParser(movie); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	if err := h.MovieService.DeleteMovieByID(movie.ID); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return c.JSON(movie)
}
