package controllers

import (
	"mundo/src/products/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type DeleteProductByIDController struct {
	uc *application.DeleteProduct
}

func NewDeleteProductByIDController(uc *application.DeleteProduct) *DeleteProductByIDController {
	return &DeleteProductByIDController{uc: uc}
}

func (ctrl *DeleteProductByIDController) Run(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = ctrl.uc.Run(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encontrar el producto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})

}