package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
	"products_api.com/p/src/Requests/application/services"
)

type ReduceStockOfAProductController struct {
	service application.ReduceStockOfAProductUC
	another_service services.ConfirmValidationRequestService
}

func NewReduceStockOfAProductController(
	uc application.ReduceStockOfAProductUC,
	s services.ConfirmValidationRequestService,
	)*ReduceStockOfAProductController{
	return&ReduceStockOfAProductController{service: uc, another_service: s}
}

func(controller *ReduceStockOfAProductController)Execute(c *gin.Context){
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

	fmt.Println("Entrada: ", input)

	id_status, err := controller.service.Execute(input.Id_request, input.Id_product, input.Quantity);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar product",
		})
		return
	}

	if err := controller.another_service.Execute(input.Id_request, id_status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al retornar resultado del pedido",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Existencias modificadas, estado del perdido cambiado",
	})
}