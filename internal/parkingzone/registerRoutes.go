package parkingzone

import (
	"spot_sync/internal/auth"
	"spot_sync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB,) {
	parkingRepository :=NewRepository(db)
	jwtService := auth.NewJWTService("")
	parkingService := NewService(parkingRepository, jwtService)
	parkingHandler := NewHandler(parkingService)

	api := e.Group("/api/v1/parking-zones")
	api.POST("", parkingHandler.CreateParkingZone, middlewares.AuthMiddleware(jwtService), middlewares.AdminMiddleware())
	api.GET("",parkingHandler.GetAllParkingZones)
	api.GET("/:id",parkingHandler.GetParkingZoneByID)
}
