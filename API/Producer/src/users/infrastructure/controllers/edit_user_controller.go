package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/users/application"
)

type EditUserController struct {
	Services application.EditUserUC
}

func NewEditUserController(uc application.EditUserUC)*EditUserController{
	return&EditUserController{Services: uc}
}

func(controller *EditUserController)Execute(c *gin.Context){
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

	var input struct {
		Name string `json:"name"`
		Last_name string `json:"last_name"`
		Email string `json:"email"`
		Cellphone string `json:"cellphone"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.Services.Execute(
		id_number, input.Name, input.Last_name, input.Email, input.Cellphone, input.Password,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario editado",
	})
}