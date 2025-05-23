package main

import (
	controller "github.com/phy749/LearnEnglish/Controller"
	router "github.com/phy749/LearnEnglish/Router"
	service "github.com/phy749/LearnEnglish/Service"
	"github.com/phy749/LearnEnglish/config"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/repository"
)

func main() {
	config.ConnectDB() // Kết nối DB
	config.ConnectRedis()

	var userRepo irepository.IUserRepository = repository.CreateUserRepository(config.DB)
	var userService iservice.IUserService = service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := router.SetupRouterUser(userController)
	r.Run(":8000")
}
