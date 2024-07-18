package main

import (
	"hotel/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Get("/api/locations", handlers.GetLocations)
	app.Get("/api/hotels", handlers.GetHotels)
	app.Get("/api/rooms", handlers.GetRooms)
	app.Post("/api/proposal", handlers.SaveProposal)
	app.Static("/", "./build")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	app.Listen(":" + port)
}
