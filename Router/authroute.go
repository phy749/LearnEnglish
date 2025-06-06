package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
)

func SetupRouterAuth(r *gin.Engine, authController *controller.AuthController) *gin.Engine {

	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/refresh-token", authController.RefreshToken)
		auth.POST("/logout", authController.Logout)
		auth.POST("/forgot-password", authController.ForgotPassword)
	}

	return r
}
