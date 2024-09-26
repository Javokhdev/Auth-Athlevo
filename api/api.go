package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-athlevo/api/docs"
	"auth-athlevo/api/handlers"
	"auth-athlevo/api/middleware"
)

// @title Authentication Service API
// @version 1.0
// @description API for Authentication Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Engine(handler *handlers.Handlers) *gin.Engine {
	router := gin.Default()
	ca := middleware.CasbinEnforcer()
	err := ca.LoadPolicy()
	if err != nil {
		panic(err)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	protected := router.Group("/")
	protected.Use(middleware.JWTMiddleware(), middleware.CasbinMiddleware(ca))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.LoginUser)
	router.POST("/forgot-password", handler.ForgotPassword)
	router.PUT("/reset-password", handler.ResetPassword)
	router.GET("/refresh-token/:email", handler.RefreshToken)
	protected.PUT("/change-role", handler.ChangeRole)
	router.GET("/validate", handler.Validate)
	user := router.Group("/user").Use(middleware.JWTMiddleware(), middleware.CasbinMiddleware(ca))
	{
		user.GET("/get-profiles", handler.GetProfile)
		user.PUT("/profiles", handler.EditProfile)
		user.PUT("/passwords", handler.ChangePassword)
		user.GET("/setting", handler.GetSetting)
		user.PUT("/setting", handler.EditSetting)
		user.DELETE("/delete/:id", handler.DeleteUser)
	}

	dashboard := router.Group("/dashboard").Use(middleware.JWTMiddleware(), middleware.CasbinMiddleware(ca))
	{
		dashboard.GET("/access-count", handler.GetPersonalAccessCount)
        dashboard.GET("/booking-revenue", handler.GetTotalPersonalBookingRevenue)
		dashboard.GET("/subscription/access-count", handler.GetAccessCountBySubscriptionID)
		dashboard.GET("/subscription/booking-revenue", handler.GetBookingRevenueBySubscriptionID)
		dashboard.GET("/total-amount", handler.TotalAmount)
		dashboard.GET("/total-members", handler.TotalMembers)
		dashboard.GET("/total-men", handler.TotalMen)
		dashboard.GET("/total-women", handler.TotalWomen)
	}

	return router
}
