package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
)

type GetAllProductsController struct {
	service application.GetAllProductsUC
}

func NewGetAllProductsController(uc application.GetAllProductsUC)*GetAllProductsController{
	return&GetAllProductsController{service: uc}
}

func(controller *GetAllProductsController)Execute(c *gin.Context){
	results, err := controller.service.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener productos",
			"Log": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Productos obtenidos:",
		"Results": results,
	})
}