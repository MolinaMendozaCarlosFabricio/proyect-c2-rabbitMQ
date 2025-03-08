package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
)

type GetProductsOfARequestController struct {
	service application.GetProductsOfARequestUC
}

func NewGetProductsOfARequestController(uc application.GetProductsOfARequestUC)*GetProductsOfARequestController{
	return&GetProductsOfARequestController{service: uc}
}

func(controller *GetProductsOfARequestController)Execute(c *gin.Context){
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

	results, err := controller.service.Execute(id_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al iniciar servidor",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Productos obtenidos:",
		"Results": results,
	})
}