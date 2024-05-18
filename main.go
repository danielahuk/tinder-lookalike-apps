package main

import (
	"log"
	"tinder-apps/internal/database"
	"tinder-apps/internal/middleware"
	"tinder-apps/internal/routes"
	"tinder-apps/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	//Initialize the database
	db, err := database.InitDB(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	//Initialize Fiber
	app := fiber.New()

	//Register routes
	routes.RegisterRoutes(app, db, cfg.JWT.Secret)

	app.Use(middleware.JWTMiddleware())

	//Member routes
	routes.MemberRoutes(app, db)

	//Start the server
	log.Fatal(app.Listen(cfg.Server.Address))
}
