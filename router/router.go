package router

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, handler *handler.MovieHandler) {
	auth := app.Group("/api", authentication.DeserializeRequiresAuth)

	app.Get("/movies", handler.GetAllMovies)

	auth.Get("/movies/:id", handler.GetMovieByID)
	auth.Post("/movies", handler.CreateMovie)
	auth.Put("/movies/:id", handler.UpdateMovie)
	auth.Delete("/movies/:id", handler.DeleteMovie)

	// app.Get("/api/movies", handler.GetAllMovies)
	// app.Get("/api/movies/:id", handler.GetMovieByID)
	// app.Post("/api/movies", handler.CreateMovie)
	// app.Put("/api/movies/:id", handler.UpdateMovie)
	// app.Delete("/api/movies/:id", handler.DeleteMovie)
}
