package router

import (
	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/middleware"
	"MovieReviewAPIs/utility"
	"log"

	_ "MovieReviewAPIs/middleware"

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
	// Create value Sub Router
	// app.Mount("/api", sub)

	// sub.Route("/auth", func(r fiber.Router) {
	// 	r.Post("/login", Uhandler.RegisterUserHandler)
	// 	r.Post("/register", Uhandler.RegisterUserHandler)
	// 	r.Post("/logout", Uhandler.LogoutUserHandler)
	// })
	// sub.Route("/movies", func(r fiber.Router) {
	// 	r.Get("/", Mhandler.GetMovieForHomePage)
	// 	r.Get("/:id", Mhandler.GetMovieByID)
	// 	r.Post("/", Mhandler.CreateMovie)
	// 	r.Put("/:id", Mhandler.UpdateMovie)
	// 	r.Delete("/:id", Mhandler.DeleteMovie)
	// })
	app.Get("/movies", Mhandler.GetAllMovies)

	app.Post("/register", Uhandler.RegisterUserHandler)
	app.Post("/login", Uhandler.LoginUserHandler)
	app.Post("/logout", middleware.MiddlewareDeserializeRout, Uhandler.LogoutUserHandler)
}

func RouterPortListener() {

}
