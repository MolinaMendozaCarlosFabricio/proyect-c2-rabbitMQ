package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/users/application"
)

type GetOneUserController struct {
	Service application.GetOneUserUC
}

func NewGetOneUserController(uc application.GetOneUserUC)*GetOneUserController{
	return&GetOneUserController{Service: uc}
}

func(controller *GetOneUserController)Execute(c *gin.Context){
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

	results, err := controller.Service.Execute(id_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":"Usuario obtenido",
		"Results": results,
	})
}