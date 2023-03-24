package main

import (
	"log"
	"os"

	"github.com/RianIhsan/blogbackend/database"
	"github.com/RianIhsan/blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
