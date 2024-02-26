package router

import (
	"MovieReviewAPIs/handler"

	"MovieReviewAPIs/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
	Example Router Path if used Authentication

	"localhost:3000/api/movies/:id"

*/

var app *fiber.App
var sub *fiber.App

// func Router(Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler) {
// 	MovieRouter(Mhandler)
// 	UserAuthRouter(Uhandler)
// }

func RouterControl(Mhandler *handler.MovieHandler, Uhandler *handler.UserHandler) {
	// Create value Sub Router
	app.Mount("/api", sub)

	sub.Route("/auth", func(r fiber.Router) {
		r.Post("/login", Uhandler.RegisterUserHandler)
		r.Post("/register", Uhandler.RegisterUserHandler)
		r.Post("/logout", middleware.MiddlewareDeserializeRout, Uhandler.LogoutUserHandler)
	})
	sub.Route("/movies", func(r fiber.Router) {
		r.Get("/", Mhandler.GetMovieForHomePage)
		r.Get("/:id", Mhandler.GetMovieByID)
		r.Post("/", Mhandler.CreateMovie)
		r.Put("/:id", Mhandler.UpdateMovie)
		r.Delete("/:id", Mhandler.DeleteMovie)
	})
}
