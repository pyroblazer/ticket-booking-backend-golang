package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pyroblazer/ticket-booking-backend-golang/config"
	"github.com/pyroblazer/ticket-booking-backend-golang/db"
	"github.com/pyroblazer/ticket-booking-backend-golang/handlers"
	"github.com/pyroblazer/ticket-booking-backend-golang/middlewares"
	"github.com/pyroblazer/ticket-booking-backend-golang/repositories"
	"github.com/pyroblazer/ticket-booking-backend-golang/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: envConfig.AllowedOrigins,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders: "Content-Length",
		AllowCredentials: true,
	}))

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
