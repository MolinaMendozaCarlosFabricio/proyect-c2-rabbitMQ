package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/users/application"
)

type DeleteUserController struct {
	Services application.DeleteUserUC
}

func NewDeleteUserController(uc application.DeleteUserUC)*DeleteUserController{
	return&DeleteUserController{Services: uc}
}

func(controller *DeleteUserController)Execute(c *gin.Context){
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

	if err := controller.Services.Execute(id_number); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al eliminar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario eliminado",
	})
}