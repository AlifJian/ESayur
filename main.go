package main

import (
	"ESayur/Databases"
	"ESayur/Handler"
	"ESayur/Repo"
	"ESayur/Services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// DB init
	db, err := Databases.InitDB()
	if err != nil {
		log.Fatalf("Failed To Connect to DB")
	}

	// Fiber Init
	app := fiber.New()
	app.Use(cors.New())

	pathApi := app.Group("/api/v1")

	// Repo Initialize
	mysqlRepo := Repo.NewMysqlRepo(db)

	// Services
	userServices := Services.NewUserService(mysqlRepo)

	//Handler
	userApi := pathApi.Group("/users")
	Handler.NewUserHandler(userApi, userServices)
	app.Listen("127.0.0.1:8080")
}
