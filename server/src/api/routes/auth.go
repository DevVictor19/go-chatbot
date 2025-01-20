package routes

import (
	"server/src/api/controllers"
	"server/src/api/database"
	"server/src/api/repositories"

	"github.com/gin-gonic/gin"
)

func initAuthRoutes(r *gin.Engine) {
	db := database.GetDatabase()
	customerRepository := repositories.NewCustomerRepository(db)
	authController := controllers.NewAuthController(customerRepository)

	api := r.Group("/auth")

	api.POST("/signup", authController.Signup)
}
