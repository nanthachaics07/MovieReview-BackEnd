package router

import (
	"MovieReviewAPIs/handler"
	// "MovieReviewAPIs/middlewares"

	"MovieReviewAPIs/utility"
	"log"

	"MovieReviewAPIs/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	// app.Use(middleware.Logger())

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     fURL.FrontendURL,
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Specify allowed headers for CORS //, Authorization
			AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
			AllowCredentials: true, // Specify if credentials are allowed
		},
	))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/metrics", middlewares.AuthMiddleware(), monitor.New(monitor.Config{
		Title: "Movie Review API Metrics",
	}))
}

func RouterControl(app *fiber.App, Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler, Ahandler *handler.AccountHandler) {
	// Create value Sub Router Group

	// app.Use("/auth/logout", middlewares.AuthenticationRequired)
	sub := app.Group("/auth")
	sub.Post("/singup", Uhandler.RegisterUserHandler)
	sub.Post("/singin", Uhandler.LoginUserHandler)
	sub.Post("/singout", middlewares.UserTokenMiddleware(), Uhandler.LogoutUserHandler)
	// sub.Post("/singout", Uhandler.LogoutUserHandler)

	// Admin Setting Movie Review Router Group
	admin := app.Group("/admin")
	admin.Post("/createmovie", Mhandler.CreateMovie)
	admin.Put("/updatemovie/:id", Mhandler.UpdateMovieByID)
	admin.Delete("/deletemovie/:id", Mhandler.DeleteMovie)
	admin.Get("/user/:id", middlewares.MiddlewareDeserializeRout, Ahandler.GetUserByIDHandler)
	admin.Get("/users", middlewares.MiddlewareDeserializeRout, Ahandler.UsersAccountAllHandler)

	acc := app.Group("/account")
	acc.Get("/user", Ahandler.UserAccountHandler)

	// acc.Patch("/", Uhandler.UpdateUserHandler)
	// acc.Delete("/", Uhandler.DeleteUserHandler)

	// Main User Movie Router Group
	app.Get("/", Mhandler.GetMovieForHomePage)
	app.Get("/allmovies", Mhandler.GetAllMovies)
	app.Get("/movie/:id", Mhandler.GetMovieByID)

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
