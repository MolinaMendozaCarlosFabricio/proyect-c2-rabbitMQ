package routes

import (
	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Products/application"
	"products_api.com/p/src/Products/infrastructure"
	"products_api.com/p/src/Products/infrastructure/controllers"
)

func ProductRoutes(r *gin.Engine){
	ps := infrastructure.NewProductRepoMySQL()
	cp := application.NewCreateProductUC(ps)
	cp_controller := controllers.NewCreateProductController(*cp)
	gap := application.NewGetAllProductsUC(ps)
	gap_controller := controllers.NewGetAllProductsController(*gap)
	ep := application.NewEditProductUC(ps)
	ep_controller := controllers.NewEditProductController(*ep)
	dp := application.NewDeleteProductUC(ps)
	dp_controller := controllers.NewDeleteProductController(*dp)
	gpr := application.NewGetProductsOfARequestUC(ps)
	gpr_controller := controllers.NewGetProductsOfARequestController(*gpr)

	products := r.Group("/products")
	{
		products.POST("/", cp_controller.Execute)
		products.GET("/", gap_controller.Execute)
		products.PUT("/:id", ep_controller.Execute)
		products.DELETE("/:id", dp_controller.Execute)
		products.GET("/request/:id", gpr_controller.Execute)
	}
}