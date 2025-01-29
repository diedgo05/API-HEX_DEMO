package routes

import (
	"mundo/src/products/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	routes := router.Group("/products")

	createProduct := dependencies.CreateProductController()
	getAllProduct := dependencies.GetAllProductController()
	updateProduct := dependencies.UpdateProductController()
	deleteProduct := dependencies.DeleteProductController()

	routes.POST("/add", createProduct.Run)
	routes.GET("/", getAllProduct.Run)
	routes.PUT("/update/:id", updateProduct.Run)
	routes.DELETE("/delete/:id", deleteProduct.Run)
}