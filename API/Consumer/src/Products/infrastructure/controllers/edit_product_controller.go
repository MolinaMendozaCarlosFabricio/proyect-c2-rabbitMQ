package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
)

type EditProductController struct {
	service application.EditProductUC
}

func NewEditProductController(uc application.EditProductUC)*EditProductController{
	return&EditProductController{service: uc}
}

func(controller *EditProductController)Execute(c *gin.Context){
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
		Name_product string `json:"name_product"`
		Category string `json:"category"`
		Price float32 `json:"price"`
		Stock int `json:"stock"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.service.Execute(
		id_number, input.Name_product, input.Category, input.Price, input.Stock,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar producto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto editado",
	})
}