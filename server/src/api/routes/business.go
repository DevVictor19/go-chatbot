package routes

import (
	"server/src/api/controllers"
	"server/src/api/database"
	"server/src/api/middlewares"
	"server/src/api/repositories"

	"github.com/gin-gonic/gin"
)

func initBusinessRoutes(r *gin.Engine) {
	db := database.GetDatabase()
	businessRepository := repositories.NewBusinessRepository(db)
	customerRepository := repositories.NewCustomerRepository(db)
	businessController := controllers.NewBusinessController(
		businessRepository,
		customerRepository,
	)

	api := r.Group("/business")
	api.Use(middlewares.AuthMiddleware())

	api.POST("", businessController.Create)
	api.GET("", businessController.FindAllPaginated)
	api.GET("/:id", businessController.FindById)
	api.PUT("/:id", businessController.Update)
	api.DELETE("/:id", businessController.Delete)
}
