package utility

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPass      string
	AppPort     string
	FrontendURL string

	JwtSecret    string
	JwtExpiresIn time.Duration
	JwtMaxAge    int

	ClientOrigin string
}

type Configs interface {
	GetConfig() (config Config, err error)
}

func GetConfig() (config Config, err error) {
	return LoadConfig(".env")
}

func LoadConfig(path string) (config Config, err error) {
	err = godotenv.Load(path)
	if err != nil {
		return
	}

	config = Config{
		DBPass:      os.Getenv("DB_prod"),
		AppPort:     os.Getenv("APP_port"),
		FrontendURL: os.Getenv("FRONTEND_URL"),

		JwtSecret:    os.Getenv("JWT_SECRET"),
		JwtExpiresIn: time.Duration(mustParseInt(os.Getenv("JWT_EXPIRED_IN"))),
		JwtMaxAge:    mustParseInt(os.Getenv("JWT_MAXAGE")),

		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
	}

	return config, nil
}

func mustParseInt(s string) int {
	i := 0
	fmt.Sscanf(s, "%d", &i)
	return i
}
