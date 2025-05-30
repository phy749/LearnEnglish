package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
	middleware "github.com/phy749/LearnEnglish/Middleware"
)

func SetupRouterUser(r *gin.Engine, userController *controller.UserController) *gin.Engine {

	user := r.Group("/users", middleware.AuthMiddleware())
	{
		user.GET("/GetAllUser", userController.GetAllUser)
		user.POST("/CreateUser", userController.CreateUser)
		user.PUT("/UpdateImformationUser", userController.UpdateUser)
		user.GET("/GetUserById/:id", middleware.RoleMiddleware(1), userController.FindUserById)
		user.PUT("/deactivate/:id", userController.DeactivateUser)
	}

	return r
}
