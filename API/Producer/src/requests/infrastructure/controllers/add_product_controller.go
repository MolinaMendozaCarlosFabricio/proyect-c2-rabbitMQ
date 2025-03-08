package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/requests/application"
)

type AddProductController struct {
	Service application.AddProductUC
}

func NewAddProductController(uc application.AddProductUC)*AddProductController{
	return&AddProductController{Service: uc}
}

func(controller *AddProductController)Execute(c *gin.Context){
	var input struct {
		Id_request int `json:"id_request"`
		Id_product int `json:"id_product"`
		Quantity int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.Service.Execute(input.Id_request, input.Id_product, input.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al añadir producto al pedido",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto añadido al envío",
	})
}