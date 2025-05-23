package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
)

func SetupRouterUser(userController *controller.UserController) *gin.Engine {
	r := gin.Default()

	user := r.Group("/users")
	{
		user.GET("", userController.GetAllUser)
		user.POST("", userController.CreateUser)
		user.PUT("", userController.UpdateUser)
		user.GET(":id", userController.FindUserById)
		user.PUT(":id/deactivate", userController.DeactivateUser)
	}

	return r
}
