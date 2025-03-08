package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
)

type ReduceStockOfAProductController struct {
	service application.ReduceStockOfAProductUC
}

func NewReduceStockOfAProductController(uc application.ReduceStockOfAProductUC)*ReduceStockOfAProductController{
	return&ReduceStockOfAProductController{service: uc}
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

	if err := controller.service.Execute(input.Id_request, input.Id_product, input.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Existencias modificadas, estado del perdido cambiado",
	})
}