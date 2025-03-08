package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
)

type UpdateRequestsStatusController struct {
	service application.UpdateRequestStatusUC
}

func NewUpdateRequestStatusController(uc application.UpdateRequestStatusUC)*UpdateRequestsStatusController{
	return&UpdateRequestsStatusController{service: uc}
}

func(controller *UpdateRequestsStatusController)Execute(c *gin.Context){
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

	message := "estado"
	if result {
		message = "Envío aprobado"
	} else {
		message = "Envío rechaszada"
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"Message": "Estatus actualizado, no infa",
		"Results": message,
	})
}