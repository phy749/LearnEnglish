package controller

import "github.com/phy749/LearnEnglish/iservice"

type LevelController struct {
	LevelService iservice.ILevelService
}

func NewLevelController(levelService iservice.ILevelService) *LevelController {
	return &LevelController{LevelService: levelService}
}
