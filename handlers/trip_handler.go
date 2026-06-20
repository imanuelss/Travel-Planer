package handlers

import (
	"individual-project/models"
	"individual-project/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TripHandler struct {
	service *services.TripService
}

func NewTripHandler(service *services.TripService) *TripHandler {
	return &TripHandler{service}
}

func (h *TripHandler) CreateTrip(c echo.Context) error {
	var trip models.Trip

	if err := c.Bind(&trip); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	err := h.service.CreateTrip(&trip)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(201, trip)
}

func (h *TripHandler) GetTrips(c echo.Context) error {
	trips, err := h.service.GetTrips()
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, trips)
}

func (h *TripHandler) GetTripByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.service.GetTripByID(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": "trip not found",
		})
	}

	return c.JSON(200, trip)
}

func (h *TripHandler) UpdateTrip(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req models.Trip
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	trip, err := h.service.UpdateTrip(
		uint(id),
		req.Title,
		req.DurationDays,
		req.Budget,
	)

	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, trip)
}

func (h *TripHandler) DeleteTrip(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.DeleteTrip(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "trip deleted",
	})
}
