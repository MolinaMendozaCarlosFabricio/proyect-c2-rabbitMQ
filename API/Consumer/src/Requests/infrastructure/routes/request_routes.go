package routes

import (
	"github.com/gin-gonic/gin"
	"products_api.com/p/src/Requests/application"
	"products_api.com/p/src/Requests/application/services"
	"products_api.com/p/src/Requests/infrastructure"
	"products_api.com/p/src/Requests/infrastructure/adapters"
	"products_api.com/p/src/Requests/infrastructure/controllers"
)

func RequestRoutes(r *gin.Engine){
	rs := infrastructure.NewRequestRepoMySQL()
	rq := adapters.NewRequestRepoRabbitMQ()
	urs := application.NewUpdateRequestStatusUC(rs)
	urs_controller := controllers.NewUpdateRequestStatusController(*urs)
	gmr := application.NewGetAllMyRequestsMethodUC(rs)
	gmr_controller := controllers.NewGetAllMyRequestsController(*gmr)
	rsp := application.NewReduceStockOfAProductUC(rs)
	cv := services.NewConfirmValidationRequestService(rq)
	rsp_controller := controllers.NewReduceStockOfAProductController(*rsp, *cv)
	gor := application.NewGetOneOfMyRequestsMethodUC(rs)
	gor_controller := controllers.NewGetOneOfMyRequestsMethodController(*gor)

	request := r.Group("/requests")
	{
		request.PUT("/:id", urs_controller.Execute)
		request.GET("/:id", gmr_controller.Execute)
		request.PUT("/product", rsp_controller.Execute)
		request.GET("/one/:id", gor_controller.Execute)
	}
}