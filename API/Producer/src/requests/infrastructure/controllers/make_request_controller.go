package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/requests/application"
)

type MakeRequestController struct {
	Service application.MakeRequestUC
}

func NewMakeRequestController(uc application.MakeRequestUC)*MakeRequestController{
	return&MakeRequestController{Service: uc}
}

func(controller *MakeRequestController)Execute(c *gin.Context){
	var input struct {
		Id_user int `json:"id_user"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.Service.Execute(input.Id_user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al realizar pedido",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Pedido realizado",
	})
}