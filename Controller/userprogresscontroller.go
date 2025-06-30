package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type UserProgressController struct {
	UserProgressService iservice.IUserProgressService
}

func NewUserProgressController(service iservice.IUserProgressService) *UserProgressController {
	return &UserProgressController{UserProgressService: service}
}

func (upc *UserProgressController) AddUserProgress(c *gin.Context) {
	var progress model.UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := upc.UserProgressService.AddUserProgress(&progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

func (upc *UserProgressController) RemoveUserProgress(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user progress id"})
		return
	}
	if err := upc.UserProgressService.RemoveUserProgress(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User progress deleted"})
}

func (upc *UserProgressController) UpdateUserProgress(c *gin.Context) {
	var progress model.UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := upc.UserProgressService.UpdateUserProgress(&progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

func (upc *UserProgressController) GetUserProgressByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user progress id"})
		return
	}
	progress, err := upc.UserProgressService.GetUserProgressByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

func (upc *UserProgressController) GetProgressByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	progresses, err := upc.UserProgressService.GetProgressByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progresses)
}
