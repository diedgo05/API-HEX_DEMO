package controllers

import (
	"mundo/src/products/application"
	"mundo/src/products/domain"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type UpdateProductByIDController struct {
	uc *application.UpdateProduct
}

func NewUpdateProductByIDController(uc *application.UpdateProduct) *UpdateProductByIDController {
	return &UpdateProductByIDController{uc: uc}
}

func (ctrl *UpdateProductByIDController) Run(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := ctrl.uc.Run(idInt, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}