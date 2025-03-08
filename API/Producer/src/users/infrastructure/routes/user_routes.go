package routes

import (
	"github.com/gin-gonic/gin"
	"request_api.com/r/src/users/application"
	"request_api.com/r/src/users/infrastructure"
	"request_api.com/r/src/users/infrastructure/controllers"
)

func UserRoutes(r *gin.Engine){
	us := infrastructure.NewUserRepoMySQL()
	ru := application.NewRegisterUserUC(us)
	ru_controller := controllers.NewRegisterUserController(*ru)
	gou := application.NewGetOneUserUC(us)
	gou_controller := controllers.NewGetOneUserController(*gou)
	eu := application.NewEditUserUC(us)
	eu_controller := controllers.NewEditUserController(*eu)
	du := application.NewDeleteUserUC(us)
	du_controller := controllers.NewDeleteUserController(*du)

	users := r.Group("/users")
	{
		users.POST("/", ru_controller.Execute)
		users.GET("/:id", gou_controller.Execute)
		users.PUT("/:id", eu_controller.Execute)
		users.DELETE("/:id", du_controller.Execute)
	}
}