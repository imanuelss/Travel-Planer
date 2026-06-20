package routes

import (
	"individual-project/config"
	"individual-project/handlers"
	"individual-project/middleware"
	"individual-project/repositories"
	"individual-project/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	tripRepo := repositories.NewTripRepository(config.DB)
	tripService := services.NewTripService(tripRepo)
	tripHandler := handlers.NewTripHandler(tripService)

	destRepo := repositories.NewDestinationRepository(config.DB)
	destService := services.NewDestinationService(destRepo)
	destHandler := handlers.NewDestinationHandler(destService)

	itineraryRepo := repositories.NewItineraryRepository(config.DB)
	itineraryService := services.NewItineraryService(itineraryRepo)
	itineraryHandler := handlers.NewItineraryHandler(itineraryService)
	aiRepo := repositories.NewAIGenerationRepository(config.DB)
	aiService := services.NewAIGenerationService(aiRepo)
	aiHandler := handlers.NewAIGenerationHandler(aiService)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Wonderlog API running",
		})
	})

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware())

	auth.GET("/me", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "authorized",
		})
	})
	auth.POST("/trips", tripHandler.CreateTrip)
	auth.GET("/trips", tripHandler.GetTrips)
	auth.GET("/trips/:id", tripHandler.GetTripByID)
	auth.PUT("/trips/:id", tripHandler.UpdateTrip)
	auth.DELETE("/trips/:id", tripHandler.DeleteTrip)
	auth.POST("/destinations", destHandler.CreateDestination)
	auth.GET("/destinations", destHandler.GetDestinations)
	auth.GET("/destinations/:id", destHandler.GetDestinationByID)
	auth.PUT("/destinations/:id", destHandler.UpdateDestination)
	auth.DELETE("/destinations/:id", destHandler.DeleteDestination)
	auth.POST("/itineraries", itineraryHandler.CreateItinerary)
	auth.GET("/itineraries", itineraryHandler.GetItineraries)
	auth.GET("/itineraries/:id", itineraryHandler.GetItineraryByID)
	auth.PUT("/itineraries/:id", itineraryHandler.UpdateItinerary)
	auth.DELETE("/itineraries/:id", itineraryHandler.DeleteItinerary)
	auth.POST("/ai-generations", aiHandler.Create)
	auth.GET("/ai-generations", aiHandler.GetAll)
	auth.GET("/ai-generations/:id", aiHandler.GetByID)
}
