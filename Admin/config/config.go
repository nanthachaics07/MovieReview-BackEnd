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
			Host:     "192.168.1.39",
			Port:     5432,
			Username: "piuser",
			Password: "pipassword",
			Name:     "pidatabase",
		},
	}
}
