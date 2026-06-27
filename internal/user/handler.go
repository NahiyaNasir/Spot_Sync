package user

import (
	"errors"
	"net/http"

	"spot_sync/internal/httpresponse"
	"spot_sync/internal/user/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *service
}

func NewHandler(service *service) *Handler {
	return &Handler{service: service}
}



func (h *Handler) CreateUser(c *echo.Context) error {
	var req dto.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
	}
	response, err := h.service.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response)
}
 func (h *Handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(&req)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Cannot login user",
				Details: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to login user",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
func (h *Handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Cannot get user information",
			Details: "missing user id in context",
		})
	}

	email, _ := c.Get("user_email").(string)
	name, _ := c.Get("user_name").(string)

	return c.JSON(http.StatusOK, dto.Response{
		ID:    userID,
		Name:  name,
		Email: email,
	})
}

 func (h *Handler) GetAllUsers(c *echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve users",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) Update(c *echo.Context) error {
	_ = c.Param("id")
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "not implemented"})
}

func (h *Handler) Delete(c *echo.Context) error {
	_ = c.Param("id")
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "not implemented"})
}
