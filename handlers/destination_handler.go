package handlers

import (
	"individual-project/models"
	"individual-project/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DestinationHandler struct {
	service *services.DestinationService
}

func NewDestinationHandler(service *services.DestinationService) *DestinationHandler {
	return &DestinationHandler{service}
}

// CreateDestination handles the creation of a new destination
func (h *DestinationHandler) CreateDestination(c echo.Context) error {
	var destination models.Destination

	if err := c.Bind(&destination); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	err := h.service.CreateDestination(&destination)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(201, destination)
}

// GetDestinations retrieves all destinations
func (h *DestinationHandler) GetDestinations(c echo.Context) error {
	destinations, err := h.service.GetDestinations()
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, destinations)
}

// GetDestinationByID retrieves a destination by its ID
func (h *DestinationHandler) GetDestinationByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	destination, err := h.service.GetDestinationByID(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": "destination not found",
		})
	}

	return c.JSON(200, destination)
}

// UpdateDestination handles the update of an existing destination
func (h *DestinationHandler) UpdateDestination(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var req models.Destination
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	dest, err := h.service.UpdateDestination(
		uint(id),
		req.Name,
		req.Country,
		req.Description,
		req.EstimatedDailyCost,
	)

	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, dest)
}

// DeleteDestination handles the deletion of a destination by its ID
func (h *DestinationHandler) DeleteDestination(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteDestination(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "destination deleted",
	})
}
