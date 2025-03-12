package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
)

type GetOneOfMyRequestsMethodController struct {
	controller application.GetOneOfMyRequestsMethodUC
}

func NewGetOneOfMyRequestsMethodController(c application.GetOneOfMyRequestsMethodUC)*GetOneOfMyRequestsMethodController{
	return&GetOneOfMyRequestsMethodController{controller: c}
}

func(controller *GetOneOfMyRequestsMethodController)Execute(c *gin.Context){
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

	result, err := controller.controller.Execute(id_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Parámetro incorrecto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Pedido enviado",
		"Results": result,
	})
}