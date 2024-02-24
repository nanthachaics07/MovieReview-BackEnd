package router

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/handler"

	"github.com/gofiber/fiber/v2"
)

/*
	Example Router Path if used Authentication

	"localhost:3000/api/movies/:id"

*/

func MovieRouter(app *fiber.App, Mhandler *handler.MovieHandler) {

	app.Get("/movies", Mhandler.GetAllMovies)

	auth := app.Group("/api", authentication.DeserializeRequiresAuth)
	auth.Get("/movies/:id", Mhandler.GetMovieByID)
	auth.Post("/movies", Mhandler.CreateMovie)
	auth.Put("/movies/:id", Mhandler.UpdateMovie)
	auth.Delete("/movies/:id", Mhandler.DeleteMovie)

	// app.Get("/api/movies", handler.GetAllMovies)
	// app.Get("/api/movies/:id", handler.GetMovieByID)
	// app.Post("/api/movies", handler.CreateMovie)
	// app.Put("/api/movies/:id", handler.UpdateMovie)
	// app.Delete("/api/movies/:id", handler.DeleteMovie)
}

func UserRouter(app *fiber.App, Uhandler *handler.UserHandler) {

	app.Post("/register", Uhandler.RegisterUser)
	app.Post("/login", Uhandler.LoginUser)

	auth := app.Group("/api", authentication.DeserializeRequiresAuth)
	auth.Post("/logout", Uhandler.LogoutUser)

}
