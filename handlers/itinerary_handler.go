package handlers

import (
	"individual-project/models"
	"individual-project/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItineraryHandler struct {
	service *services.ItineraryService
}

func NewItineraryHandler(service *services.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{service}
}

// CreateItinerary handles the creation of a new itinerary
func (h *ItineraryHandler) CreateItinerary(c echo.Context) error {
	var itinerary models.Itinerary

	if err := c.Bind(&itinerary); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	err := h.service.CreateItinerary(&itinerary)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(201, itinerary)
}

// get all itineraries
func (h *ItineraryHandler) GetItineraries(c echo.Context) error {
	itineraries, err := h.service.GetItineraries()
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, itineraries)
}

// get itinerary by id
func (h *ItineraryHandler) GetItineraryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	itinerary, err := h.service.GetItineraryByID(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": "itinerary not found",
		})
	}

	return c.JSON(200, itinerary)
}

// UpdateItinerary handles the update of an existing itinerary by ID
func (h *ItineraryHandler) UpdateItinerary(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req models.Itinerary
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	itinerary, err := h.service.UpdateItinerary(
		uint(id),
		req.DayNumber,
		req.Activity,
		req.Location,
		req.EstimatedCost,
	)

	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, itinerary)
}

// DeleteItinerary handles the deletion of an itinerary by ID
func (h *ItineraryHandler) DeleteItinerary(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.DeleteItinerary(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "itinerary deleted",
	})
}
