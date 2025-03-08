package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
)

type CreateProductController struct {
	service application.CreateProductUC
}

func NewCreateProductController(uc application.CreateProductUC)*CreateProductController{
	return&CreateProductController{service: uc}
}

func(c *CreateProductController)Execute(ctx *gin.Context){
	var input struct {
		Name_product string `json:"name_product"`
		Category string `json:"category"`
		Price float32 `json:"price"`
		Stock int `json:"stock"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := c.service.Execute(input.Name_product, input.Category, input.Price, input.Stock); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Producto registrado",
	})
}