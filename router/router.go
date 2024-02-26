package router

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/utility"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*
	Example Router Path if used Authentication

	"localhost:3000/api/movies/:id"

*/

func InitRouterConfig(app *fiber.App) {
	// app.Use(authentication.DeserializeRequiresAuth)

	fURL, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     fURL.FrontendURL,
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Specify allowed headers for CORS
			AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
			AllowCredentials: true, // Specify if credentials are allowed
		},
	))
}

func Router(app *fiber.App, Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler) {
	MovieRouter(app, Mhandler)
	UserRouter(app, Uhandler)
}

func MovieRouter(app *fiber.App, Mhandler *handler.MovieHandler) {

	app.Get("/movies", Mhandler.GetAllMovies)
	app.Get("/", Mhandler.GetMovieForHomePage)

	auth := app.Group("/api", authentication.DeserializeRequiresAuth)
	app.Get("/movies/:id", Mhandler.GetMovieByID)
	auth.Post("/movies", Mhandler.CreateMovie)
	auth.Put("/movies/:id", Mhandler.UpdateMovie)
	auth.Delete("/movies/:id", Mhandler.DeleteMovie)

}

func UserRouter(app *fiber.App, Uhandler *handler.UserHandler) {

	app.Post("/register", Uhandler.RegisterUser)
	app.Post("/login", Uhandler.LoginUser)

	// auth := app.Group("/api", authentication.DeserializeRequiresAuth)
	app.Post("/logout", Uhandler.LogoutUser)
	// auth.Get("/user-account", Uhandler.GetUserAccount)

}

// func AdminRouter(app *fiber.App, Uhandler *handler.AdminHandler) {

// }
