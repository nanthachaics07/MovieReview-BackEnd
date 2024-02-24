package handler

import (
	"MovieReviewAPIs/services"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DefineRoutes(app *fiber.App, movieService *services.MovieService, userService *services.UserService) error {
	// Middleware: Authentication required for movie routes
	app.Use("/movie/:id", authenticationRequired)

	// Define movie routes
	app.Get("/movies", func(c *fiber.Ctx) error {
		movies := movieService.GetAllMovies()
		return c.JSON(movies)
	})
	app.Get("/movie/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			// Handle the error
			return c.Status(500).SendString(err.Error())
		}
		movie, err := movieService.GetMovieById(uint(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(movie)
	})
	return nil
}

func authenticationRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	fmt.Println(claims)
	return c.Next()
}
