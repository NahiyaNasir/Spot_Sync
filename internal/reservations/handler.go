package reservations

import (
	"errors"
	"net/http"
	"strconv"

	"spot_sync/internal/auth"
	"spot_sync/internal/httpresponse"
	"spot_sync/internal/reservations/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
    service *service
	jwtService auth.JWTService
}

func NewHandler(service *service, jwtService auth.JWTService) *Handler {
    return &Handler{service: service, jwtService: jwtService}
}
func getCurrentUserID(c *echo.Context) (uint, bool) {
	userId, ok := c.Get("user_id").(uint)
	return userId, ok
}
func reservationErrorResponse(c *echo.Context, err error) error {
    if errors.Is(err, ErrReservationNotFound) {
        return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
            Success: false,
            Message: "Reservation not found",
            Errors:  "Reservation not found",
        })
    }

    return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
        Success: false,
        Message: "Something went wrong",
        Errors:  err.Error(),
    })
}

func (h *Handler) CreateReservation(c *echo.Context) error {
    var req dto.CreateReservationRequest
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

    raw := c.Get("user_id")
    if raw == nil {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Success: false,
            Message: "Unauthorized",
            Errors:  "User not authenticated",
        })
    }

    userID, ok := raw.(uint)
    if !ok {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Success: false,
            Message: "Unauthorized",
            Errors:  "User not authenticated",
        })
            
        
    }
// fmt.Printf("DEBUG handler user_id: %v (type: %T)\n", raw, raw)
    response, err := h.service.CreateReservation(userID, &req)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Success: false,
            Message: "Failed to create reservation",
            Errors:  err.Error(),
        })
    }

    return c.JSON(http.StatusCreated, httpresponse.SuccessResponse{
        Success: true,
        Message: "Reservation created successfully",
        Data:    response,
    })
}

func (h *Handler) GetAllReservations(c *echo.Context) error {
    responses, err := h.service.GetAllReservations()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Success: false,
            Message: "Failed to retrieve reservations",
            Errors:  err.Error(),
        })
    }

    return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
        Success: true,
        Message: "Reservations retrieved successfully",
        Data:    responses,
    })
}

func (h *Handler) GetMyReservations(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Success: false,
			Message: "Unauthorized",
			Errors:  "User not authenticated",
		})
	}

	bookings, err := h.service.GetMyReservations(userId)
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
		Success: true,
		Message: "Reservations retrieved successfully",
		Data:    bookings,
	})
}


func (h *Handler) CancelReservation(c *echo.Context) error {
    raw := c.Get("user_id")
    if raw == nil {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Success: false,
            Message: "Unauthorized",
            Errors:  "User not authenticated",
        })
    }
    userID := raw.(uint)

    reservationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
            Success: false,
            Message: "Invalid reservation ID",
            Errors:  err.Error(),
        })
    }

    if err := h.service.CancelReservation(userID, uint(reservationID)); err != nil {
        if errors.Is(err, ErrReservationNotFound) {
            return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
                Success: false,
                Message: "Reservation not found",
                Errors:  err.Error(),
            })
        }
        if errors.Is(err, ErrUnauthorized) {
            return c.JSON(http.StatusForbidden, httpresponse.ErrorResponse{
                Success: false,
                Message: "You can only cancel your own reservations",
                Errors:  err.Error(),
            })
        }
        if errors.Is(err, ErrAlreadyCancelled) {
            return c.JSON(http.StatusConflict, httpresponse.ErrorResponse{
                Success: false,
                Message: "Reservation is already cancelled",
                Errors:  err.Error(),
            })
        }
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Success: false,
            Message: "Failed to cancel reservation",
            Errors:  err.Error(),
        })
    }

    return c.JSON(http.StatusOK, httpresponse.SuccessResponse{
        Success: true,
        Message: "Reservation cancelled successfully",
        Data:    nil,
    })
}
