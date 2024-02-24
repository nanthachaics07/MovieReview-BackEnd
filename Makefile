dev:
	docker-compose up -d
	
dev-down:
	docker-compose down

start-server:
	go run ./cmd

install-modules:
	go get github.com/gofiber/fiber/v2
	go get -u gorm.io/gorm
	go get gorm.io/gorm/logger
	go get gorm.io/driver/postgres
	go get github.com/golang-jwt/jwt

