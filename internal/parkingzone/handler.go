package parkingzone

import (
	"errors"
	"net/http"
	"strconv"

	"spot_sync/internal/httpresponse"
	"spot_sync/internal/parkingzone/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *service
}

func parkingZoneErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrParkingZoneNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
			Success: false,
			Message: "Parking zone not found",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
		Success: false,
		Message: "Something went wrong",
		Errors:  err.Error(),
	})
}
func NewHandler(service *service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
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

	response, err := h.service.CreateParkingZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Success: false,
			Message: "Failed to create parking zone",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, httpresponse.SuccessResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    response,
	})
}
 func (h *Handler) GetAllParkingZones(c *echo.Context) error {
	zones, err := h.service.GetAllParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve parking zones",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    zones,
	})
}
 func (h *Handler) GetParkingZoneByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Success: false,
			Message: "Invalid parking zone id",
			Errors:  err.Error(),
		})
	}

	response, err := h.service.GetParkingZoneByID(uint(id)) // err => re-assign

	if err != nil {
		return parkingZoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data:    response,
	})
}
