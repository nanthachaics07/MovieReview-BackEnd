package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Movie struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	ReleaseDate   string `json:"release_date"`
	Runtime       string `json:"runtime"`
	Rating        string `json:"rating"`
	Category      string `json:"category"`
	Popularity    string `json:"popularity"`
	Budget        int    `json:"budget"`
	Revenue       int    `json:"revenue"`
	Director      string `json:"Director"`
	Casting       string `json:"casting"`
	Writers       string `json:"Writers"`
	DistributedBy string `json:"Distributed by"`
	MPAA          string `json:"mpaa_rating"`
	Description   string `json:"description"`
	ImageURL      string `json:"imageUrl"`
}

func enableCORS(h fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "https://localhost:3000")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return c.SendStatus(fiber.StatusOK)
		}

		return h(c)
	}
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

func main() {
	// Initialize database connection
	db := initializeDB()

	// Initialize Fiber app
	app := fiber.New()

	// Middleware: Enable CORS
	app.Use(enableCORS)

	// Define routes
	defineRoutes(app, db)

	// Start Fiber server
	app.Listen(":8080")
}

func initializeDB() *gorm.DB {
	// Connect to PostgreSQL database
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&Movie{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return db
}

func defineRoutes(app *fiber.App, db *gorm.DB) {

	// Middleware: Authentication required for movie routes
	app.Use("/movie/:id", authenticationRequired)

	// Define movie routes
	app.Get("/movies", func(c *fiber.Ctx) error {
		movies := GetAllMovies(db)
		return c.JSON(movies)
	})
	app.Get("/movie/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(GetMovieById(db, uint(id)))
	})

	// Define user routes
	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		err := RegisterUser(db, user.Email, user.Password)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fmt.Println("User created successfully")
		return c.JSON(user)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		token, err := LoginUser(db, user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})
		return c.JSON(map[string]string{
			"message": "User logged in successfully",
			"token":   "Your token is: " + token,
		})
	})

	app.Post("/logout", LogoutUser)
}

func GetAllMovies(db *gorm.DB) []Movie {
	var movies []Movie
	db.Find(&movies)
	return movies
}

func GetMovieById(db *gorm.DB, id uint) *Movie {
	var movie Movie
	err := db.First(&movie, id)
	if err != nil {
		log.Printf("Error getting movie: %v", err)
	}
	return &movie
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}

func LoginUser(db *gorm.DB, user *User) (string, error) {
	selectedUser := new(User)
	result := db.Where("email =?", user.Email).First(selectedUser)
	if result.Error != nil {
		return "", result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		log.Printf("Password does not match : %v", err)
		return "", err
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": selectedUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func RegisterUser(db *gorm.DB, email, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &User{
		Email:    email,
		Password: string(hashPass),
	}
	result := db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func LogoutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(map[string]string{
		"message": "User logged out successfully",
	})
}
