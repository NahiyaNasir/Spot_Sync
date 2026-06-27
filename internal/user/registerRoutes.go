package user

import (
	"spot_sync/internal/auth"
	"spot_sync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	// fmt.Println("RegisterRoutes called")
	userRepository :=NewRepository(db)
	jwtService := auth.NewJWTService("")
	userService := NewService(userRepository,jwtService)
	userHandler := NewHandler(userService)
		api := e.Group("/api/v1/auth")
		api.POST("/register", userHandler.CreateUser)
		api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService)) 
	api.GET("/allUsers",userHandler.GetAllUsers, middlewares.AuthMiddleware(jwtService), middlewares.AdminMiddleware()	)
	
}