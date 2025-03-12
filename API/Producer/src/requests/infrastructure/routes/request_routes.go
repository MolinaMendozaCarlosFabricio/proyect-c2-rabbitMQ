package routes

import (
	"github.com/gin-gonic/gin"
	"request_api.com/r/src/requests/application"
	"request_api.com/r/src/requests/infrastructure"
	"request_api.com/r/src/requests/application/services"
	"request_api.com/r/src/requests/infrastructure/adapters"
	"request_api.com/r/src/requests/infrastructure/controllers"
)

func RequestRoutes(r *gin.Engine){
	rs := infrastructure.NewRequestRepoMySQL()
	rr := adapters.NewRequestRepoRabbitMQ()
	mr := application.NewMakeRequestUC(rs)
	mr_controller := controllers.NewMakeRequestController(*mr)
	ap := application.NewAddProductUC(rs)
	sv := services.NewSendRequestToVerifyMethodService(rr)
	ap_controller := controllers.NewAddProductController(*ap, *sv)

	request := r.Group("/requests")
	{
		request.POST("/request", mr_controller.Execute)
		request.POST("/product", ap_controller.Execute)
	}
}