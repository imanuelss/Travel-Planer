package handlers

import (
	"individual-project/helpers"
	"individual-project/models"
	"individual-project/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AIGenerationHandler struct {
	service *services.AIGenerationService
}

func NewAIGenerationHandler(service *services.AIGenerationService) *AIGenerationHandler {
	return &AIGenerationHandler{service}
}

func (h *AIGenerationHandler) Create(c echo.Context) error {
	var ai models.AIGeneration

	if err := c.Bind(&ai); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	result, err := helpers.GenerateWithGemini(ai.Prompt)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	ai.GeneratedResult = result

	err = h.service.Create(&ai)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(201, ai)
}

func (h *AIGenerationHandler) GetAll(c echo.Context) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, data)
}

func (h *AIGenerationHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(404, map[string]string{
			"error": "data not found",
		})
	}

	return c.JSON(200, data)
}
