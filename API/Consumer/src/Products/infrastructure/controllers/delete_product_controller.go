package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
)

type DeleteProductController struct {
	service application.DeleteProductUC
}

func NewDeleteProductController(uc application.DeleteProductUC)*DeleteProductController{
	return&DeleteProductController{service: uc}
}

func(controller *DeleteProductController)Execute(c *gin.Context){
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

	if err := controller.service.Execute(id_number); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al eliminar producto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto eliminado",
	})
}