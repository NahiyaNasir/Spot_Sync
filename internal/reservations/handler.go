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
            Code:    http.StatusNotFound,
            Message: "Reservation not found",
        })
    }

    return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
        Code:    http.StatusInternalServerError,
        Message: "Something went wrong",
        Details: err.Error(),
    })
}

func (h *Handler) CreateReservation(c *echo.Context) error {
    var req dto.CreateReservationRequest
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

    raw := c.Get("user_id")
    if raw == nil {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Code:    http.StatusUnauthorized,
            Message: "Unauthorized",
        })
    }

    userID, ok := raw.(uint)
    if !ok {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Code:    http.StatusUnauthorized,
            Message: "Unauthorized",
    
        })
            
        
    }
// fmt.Printf("DEBUG handler user_id: %v (type: %T)\n", raw, raw)
    response, err := h.service.CreateReservation(userID, &req)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to create reservation",
            Details: err.Error(),
        })
    }

    return c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetAllReservations(c *echo.Context) error {
    responses, err := h.service.GetAllReservations()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to retrieve reservations",
            Details: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, responses)
}

func (h *Handler) GetMyReservations(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	bookings, err := h.service.GetMyReservations(userId)
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, bookings)
}


func (h *Handler) CancelReservation(c *echo.Context) error {
    raw := c.Get("user_id")
    if raw == nil {
        return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
            Code:    http.StatusUnauthorized,
            Message: "Unauthorized",
        })
    }
    userID := raw.(uint)

    reservationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid reservation ID",
            Details: err.Error(),
        })
    }

    if err := h.service.CancelReservation(userID, uint(reservationID)); err != nil {
        if errors.Is(err, ErrReservationNotFound) {
            return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
                Code:    http.StatusNotFound,
                Message: "Reservation not found",
            })
        }
        if errors.Is(err, ErrUnauthorized) {
            return c.JSON(http.StatusForbidden, httpresponse.ErrorResponse{
                Code:    http.StatusForbidden,
                Message: "You can only cancel your own reservations",
            })
        }
        if errors.Is(err, ErrAlreadyCancelled) {
            return c.JSON(http.StatusConflict, httpresponse.ErrorResponse{
                Code:    http.StatusConflict,
                Message: "Reservation is already cancelled",
            })
        }
        return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to cancel reservation",
            Details: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Reservation cancelled successfully",
    })
}
