package router

import (
	_ "MovieReviewAPIs/authentication"
	"MovieReviewAPIs/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/api/movies", handler.GetAllMovies)
	app.Get("/api/movies/:id", handler.GetMovieByID)
	app.Post("/api/movies", handler.CreateMovie)
	app.Put("/api/movies/:id", handler.UpdateMovie)
	app.Delete("/api/movies/:id", handler.DeleteMovie)
}
