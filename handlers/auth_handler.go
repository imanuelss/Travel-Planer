package handlers

import (
	"individual-project/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

// request struct for register
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// register endpoint
func (h *AuthHandler) Register(c echo.Context) error {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "invalid request")
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(201, user)
}

// request struct for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid request",
		})
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(401, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "login success",
		"token":   token,
	})
}
