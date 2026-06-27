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
			Code:    http.StatusNotFound,
			Message: "Parking zone not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}
func NewHandler(service *service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
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

	response, err := h.service.CreateParkingZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}
 func (h *Handler) GetAllParkingZones(c *echo.Context) error {
	zones, err := h.service.GetAllParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve parking zones",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, zones)
}
 func (h *Handler) GetParkingZoneByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid parking zone id",
			Details: err.Error(),
		})
	}

	response, err := h.service.GetParkingZoneByID(uint(id)) // err => re-assign

	if err != nil {
		return parkingZoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)

	}	
