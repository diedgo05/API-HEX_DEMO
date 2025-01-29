package controllers

import (
	"mundo/src/products/application"
	"mundo/src/products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)


// var validate *validator.Validate

type CreateProductController struct {
	uc *application.CreateProduct
}

//Constructor --> permite inyectar la dependencia del caso de uso (application.CreateProduct) al controlador. 	
func NewCreateProductController(uc *application.CreateProduct) *CreateProductController {
	return &CreateProductController{uc: uc}
}

func (ctrl *CreateProductController) Run(c *gin.Context) {
	var products domain.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Todos los campos son requeridos"})
		return 
	}

	err := ctrl.uc.Run(products)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"mensaje":"Producto creado"})
		c.JSON(http.StatusOK, products)
	}

		// c.JSON(http.StatusOK, products)

	
	// c.JSON(http.StatusOK, products)
}
