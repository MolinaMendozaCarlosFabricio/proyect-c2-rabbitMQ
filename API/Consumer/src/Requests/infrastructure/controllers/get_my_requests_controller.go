package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
)

type GetAllMyRequestsController struct {
	service application.GetAllMyRequestsMethodUC
}

func NewGetAllMyRequestsController(uc application.GetAllMyRequestsMethodUC)*GetAllMyRequestsController{
	return&GetAllMyRequestsController{service: uc}
}

func(controller *GetAllMyRequestsController)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No se pudo mapear el parámetro",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro incorrecto",
		})
		return
	}

	result, err := controller.service.Execute(id_number)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al comprobar inventar",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Pedidos conseguidos",
		"Results": result,
	})
}