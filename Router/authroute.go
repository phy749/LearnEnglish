package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
	middleware "github.com/phy749/LearnEnglish/Middleware"
)

func SetupRouterAuth(r *gin.Engine, authController *controller.AuthController) *gin.Engine {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/forgot-password", authController.ChangePassword)
	}
	authProtected := r.Group("/auth").Use(middleware.AuthMiddleware())
	{

		authProtected.POST("/logout", authController.Logout)
		authProtected.POST("/refresh-token", authController.RefreshToken)
	}

	return r
}
