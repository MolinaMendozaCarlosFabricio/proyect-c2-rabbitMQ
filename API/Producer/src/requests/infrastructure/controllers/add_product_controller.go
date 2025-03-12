package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/requests/application"
	"request_api.com/r/src/requests/application/services"
)

type AddProductController struct {
	Service application.AddProductUC
	another_service services.SendRequestToVerifyMethodService
}

func NewAddProductController(
	uc application.AddProductUC, 
	s services.SendRequestToVerifyMethodService,
	)*AddProductController{
	return&AddProductController{Service: uc, another_service: s}
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

	fmt.Println("ID del pedido: ", input.Id_request)

	if err := controller.Service.Execute(input.Id_request, input.Id_product, input.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al añadir producto al pedido",
		})
		return
	}

	if err := controller.another_service.Execute(input.Id_request, input.Id_product, input.Quantity); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al enviar producto a revisión",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto añadido al envío",
	})
}