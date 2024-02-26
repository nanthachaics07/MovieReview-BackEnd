package config

type AppConfig struct {
	Database DatabaseConfig
	// JWT      JWTConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// type JWTConfig struct {
// 	SecretKey string
// }

func LoadConfig() *AppConfig {
	return &AppConfig{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "myuser",
			Password: "mypassword",
			Name:     "mydatabase",
		},
	}
}
