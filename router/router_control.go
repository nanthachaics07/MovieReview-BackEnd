package router

import (
	"MovieReviewAPIs/handler"

	"MovieReviewAPIs/utility"
	"log"

	"MovieReviewAPIs/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*
	Example Router Path if used Authentication

	"localhost:3000/api/movies/:id"

*/

// var sub *fiber.App

// func Router(Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler) {
// 	MovieRouter(Mhandler)
// 	UserAuthRouter(Uhandler)
// }

func InitRouterHeaderConfig(app *fiber.App) {
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

func RouterControl(app *fiber.App, Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler) {
	// Create value Sub Router Group
	sub := app.Group("/auth")
	sub.Post("/register", Uhandler.RegisterUserHandler)
	sub.Post("/login", Uhandler.LoginUserHandler)
	sub.Post("/logout", middlewares.MiddlewareDeserializeRout, Uhandler.LogoutUserHandler)

	// Main User Movie Router Group
	app.Get("/movies", Mhandler.GetMovieForHomePage)
	app.Get("/allmovies", Mhandler.GetAllMovies)
	app.Get("/movie/:id", Mhandler.GetMovieByID)

	// Admin Setting Movie Review Router Group
	admin := app.Group("/admin")
	admin.Post("/movie", Mhandler.CreateMovie)
	admin.Put("/movie/:id", Mhandler.UpdateMovie)
	admin.Delete("/movie/:id", Mhandler.DeleteMovie)

	/*
		Force Delete movie or Delete All
		## Pls Contract Developer
			- Use Only for testing นะจ๊ะ
	*/
}

func RouterPortListener(app *fiber.App) {
	// Load configuration
	config, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Start Fiber server
	err = app.Listen(config.AppPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
