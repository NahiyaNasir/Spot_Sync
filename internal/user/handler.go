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
			Success: false,
			Message: "Invalid request body",
			Errors:  err.Error(),
		})
	}
	response, err := h.service.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, httpresponse.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    response,
	})
}
 func (h *Handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Success: false,
			Message: "Invalid request payload",
			Errors:  err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  err.Error(),
		})
	}

	response, err := h.service.LoginUser(&req)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
				Success: false,
				Message: "Cannot login user",
				Errors:  err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Success: false,
			Message: "Failed to login user",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
		Success: true,
		Message: "User logged in successfully",
		Data:    response,
	})
}
func (h *Handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Success: false,
			Message: "Cannot get user information",
			Errors:  "missing user id in context",
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
			Success: false,
			Message: "Failed to retrieve users",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (h *Handler) Update(c *echo.Context) error {
	_ = c.Param("id")
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "not implemented"})
}

func (h *Handler) Delete(c *echo.Context) error {
	_ = c.Param("id")
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "not implemented"})
}
