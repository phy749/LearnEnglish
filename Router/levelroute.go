package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/phy749/LearnEnglish/Controller"
)

func SetupRouterLevel(r *gin.Engine, levelController *controller.LevelController)
