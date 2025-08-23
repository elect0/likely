package handler

import (
	"net/http"

	"github.com/elect0/likely/internal/service"
	"github.com/labstack/echo/v4"
)

type signUpRequst struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HTTPHandler struct {
	userService *service.UserService
}

func NewHTTPHandler(userService *service.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (h *HTTPHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/signup", h.SignUp)
	e.POST("/signin", h.SignIn)
}

func (h *HTTPHandler) SignUp(c echo.Context) error {
	var req signUpRequst
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, token, err := h.userService.SignUp(c.Request().Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *HTTPHandler) SignIn(c echo.Context) error {
	var req signInRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, token, err := h.userService.SignIn(c.Request().Context(), req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})

}
