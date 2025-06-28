package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
	middleware "github.com/phy749/LearnEnglish/Middleware"
)

func SetupRouterUser(r *gin.Engine, userController *controller.UserController) *gin.Engine {
	userProtected := r.Group("/users", middleware.AuthMiddleware())
	{
		userProtected.GET("/GetAllUser", userController.GetAllUser)
		userProtected.POST("/CreateUser", userController.CreateUser)
		userProtected.PUT("/UpdateImformationUser", userController.UpdateUser)
		userProtected.GET("/GetUserById/:id", middleware.RoleMiddleware(1), userController.FindUserById)
		userProtected.PUT("/deactivate/:id", userController.DeactivateUser)
	}

	return r
}
