package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"request_api.com/r/src/users/application"
)

type RegisterUserController struct {
	Services application.RegisterUserUC
}

func NewRegisterUserController(uc application.RegisterUserUC)*RegisterUserController{
	return&RegisterUserController{Services: uc}
}

func(controller *RegisterUserController)Execute(c *gin.Context){
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
		input.Name, input.Last_name, input.Email, input.Cellphone, input.Password,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario registrado",
	})
}