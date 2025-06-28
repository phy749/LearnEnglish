package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
	router "github.com/phy749/LearnEnglish/Router"
	Service "github.com/phy749/LearnEnglish/Service"
	_ "github.com/phy749/LearnEnglish/cmd/docs" // <-- cần import để load swagger docs
	"github.com/phy749/LearnEnglish/config"
	dataoject "github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ConnectDB() // Kết nối DB
	config.ConnectRedis()

	hub := dataoject.NewHub()
	go hub.Run()

	mux := http.NewServeMux()
	router.InitWebSocketRoute(mux, hub)

	var userRepo irepository.IUserRepository = repository.CreateUserRepository(config.DB)
	var userService iservice.IUserService = Service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	var AuthRepo irepository.IAuthRepository = repository.NewAuthRepository(config.DB)
	var AuthService iservice.IAuthService = Service.NewAuthService(AuthRepo)
	authController := controller.NewAuthController(AuthService)
	r := gin.Default()
	// r.Use(cors.Default())

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.SetupRouterUser(r, userController)
	router.SetupRouterAuth(r, authController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}
