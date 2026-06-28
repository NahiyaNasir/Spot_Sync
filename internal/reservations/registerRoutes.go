package reservations

import (
	"spot_sync/internal/auth"
	"spot_sync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
    reservationRepository := NewRepository(db)
	jwtService := auth.NewJWTService("")
    reservationService := NewService(reservationRepository)
    reservationHandler := NewHandler(reservationService, jwtService)

    api := e.Group("/api/v1/reservations")
    api.POST("", reservationHandler.CreateReservation , middlewares.AuthMiddleware(jwtService), middlewares.RoleMiddleware("admin","driver"))
    api.GET("", reservationHandler.GetAllReservations, middlewares.AuthMiddleware(jwtService), middlewares.RoleMiddleware("admin"))
    api.GET("/my-reservations", reservationHandler.GetMyReservations, middlewares.AuthMiddleware(jwtService), middlewares.RoleMiddleware("driver"))
    api.DELETE("/:id", reservationHandler.CancelReservation, middlewares.AuthMiddleware(jwtService), middlewares.RoleMiddleware("driver"))
}
