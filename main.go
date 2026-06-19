package main

import (
	"log"
	"os"

	"individual-project/config"
	"individual-project/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	e := echo.New()

	config.ConnectDB()
	routes.InitRoutes(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
