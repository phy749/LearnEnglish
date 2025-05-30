package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
	router "github.com/phy749/LearnEnglish/Router"
	service "github.com/phy749/LearnEnglish/Service"
	_ "github.com/phy749/LearnEnglish/cmd/docs" // <-- cần import để load swagger docs
	"github.com/phy749/LearnEnglish/config"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ConnectDB() // Kết nối DB
	config.ConnectRedis()

	var userRepo irepository.IUserRepository = repository.CreateUserRepository(config.DB)
	var userService iservice.IUserService = service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	var AuthRepo irepository.IAuthRepository = repository.NewAuthRepository(config.DB)
	var AuthService iservice.IAuthService = service.NewAuthService(AuthRepo)
	authController := controller.NewAuthController(AuthService)
	r := gin.Default()
	router.SetupRouterUser(r, userController)
	router.SetupRouterAuth(r, authController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}
