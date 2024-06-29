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
			AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
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
	// sub.Post("/singout", middlewares.CookieTokenMiddleware(), Uhandler.LogoutUserHandler)
	sub.Post("/singout", middlewares.CookieTokenMiddleware(), Uhandler.LogoutUserHandler)

	// Admin Setting Movie Review Router Group
	admin := app.Group("/admin")
	//Movie admin router group
	admin.Post("/createmovie", middlewares.CookieTokenMiddleware(), Mhandler.CreateMovie)
	admin.Put("/updatemovie/:id", middlewares.CookieTokenMiddleware(), Mhandler.UpdateMovieByID)
	admin.Delete("/deletemovie/:id", middlewares.CookieTokenMiddleware(), Mhandler.DeleteMovie)
	// // User Auth router group
	// admin.Get("/user/:id", Ahandler.GetUserByIDHandler)
	admin.Get("/users", middlewares.CookieTokenMiddleware(), Ahandler.UsersAccountAllHandler)
	admin.Delete("/deleteuser/:id", middlewares.CookieTokenMiddleware(), Ahandler.DeleteUserHandler)

	//User Account router group
	acc := app.Group("/account")
	acc.Get("/user", middlewares.CookieTokenMiddleware(), Ahandler.UserAccountHandler)
	acc.Patch("/updateuser", middlewares.CookieTokenMiddleware(), Ahandler.UpdateUserHandler)
	// acc.Delete("/deleteuser/:id", middlewares.CookieTokenMiddleware(), Ahandler.DeleteUserHandler)

	// acc.Patch("/", Uhandler.UpdateUserHandler)
	// acc.Delete("/", Uhandler.DeleteUserHandler)

	// Main User Movie Router Group
	app.Get("/", Mhandler.GetMovieForHomePage)
	app.Get("/allmovies", middlewares.CookieTokenMiddleware(), Mhandler.GetAllMovies)
	app.Get("/movie/:id", middlewares.CookieTokenMiddleware(), Mhandler.GetMovieByID)

	/*
		Force Delete with gorm or Delete All
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
