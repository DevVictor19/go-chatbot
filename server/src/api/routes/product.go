package routes

import (
	"server/src/api/controllers"
	"server/src/api/database"
	"server/src/api/middlewares"
	"server/src/api/repositories"

	"github.com/gin-gonic/gin"
)

func initProductRoutes(r *gin.Engine) {
	db := database.GetDatabase()
	businessRepository := repositories.NewBusinessRepository(db)
	customerRepository := repositories.NewCustomerRepository(db)
	productRepository := repositories.NewProductRepository(db)
	productController := controllers.NewProductController(
		productRepository,
		customerRepository,
		businessRepository,
	)

	api := r.Group("/products")
	api.Use(middlewares.AuthMiddleware())

	postFix := "/business/:businessId"

	api.POST(postFix, productController.Create)
	api.GET(postFix, productController.FindAllPaginated)
	api.GET("/:productId"+postFix, productController.FindById)
	api.PUT("/:productId"+postFix, productController.Update)
	api.DELETE("/:productId"+postFix, productController.Delete)
}
