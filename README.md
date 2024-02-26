# DURING DEVELOPMENT

### `Project Movie Reviews`

This project is Use Hexagonal Architecture structure. (จะได้แก้ง่ายๆ)

**Note: this is a GoAPI project - Backend only**

// Backend GolangAPIs `SOON`

Release: On branch `Develop`

Backend—Movie-hexagonal/
|-- Admin/  `Force Controller Delete, Fix, Insert Data #Backend Only`
|   |-- config/
|   |   |-- config.go
|   |-- etc/
|   |   |-- json/
|   |   |-- text/
|   |-- main.go
|   |-- DB_Strut.go
|-- cmd/
|   |-- main.go
|-- database/
|   |-- db_connection.go
|-- handler/
|   |-- errs/
|   |   |-- errs.go
|   |-- auth_handler.go
|   |-- movie_handler.go
|-- middleware/
|   |-- authMiddleRout.go
|-- models/
|   |-- log.go
|   |-- movie.go
|   |-- user.go
|-- repositories/
|   |-- auth_repository.go
|   |-- auth.go
|   |-- movie_repository.go
|   |-- movie.go
|-- router/
|   |-- router_control.go
|   |-- router.go  `//# Empty File`
|-- services/
|   |-- auth_service.go
|   |-- auth.go
|   |-- movie_service.go
|   |-- movie.go
|-- utility/
|   |-- loadConfig.go
| 
| -- (etc file)
