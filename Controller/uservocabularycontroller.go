package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type UserVocabularyController struct {
	UserVocabularyService iservice.IUserVocabularyService
}

func NewUserVocabularyController(service iservice.IUserVocabularyService) *UserVocabularyController {
	return &UserVocabularyController{UserVocabularyService: service}
}

func (uvc *UserVocabularyController) AddUserVocabulary(c *gin.Context) {
	var userVocab model.UserVocabulary
	if err := c.ShouldBindJSON(&userVocab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uvc.UserVocabularyService.AddUserVocabulary(&userVocab); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userVocab)
}

func (uvc *UserVocabularyController) RemoveUserVocabulary(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user vocabulary id"})
		return
	}
	if err := uvc.UserVocabularyService.RemoveUserVocabulary(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User vocabulary deleted"})
}

func (uvc *UserVocabularyController) UpdateUserVocabulary(c *gin.Context) {
	var userVocab model.UserVocabulary
	if err := c.ShouldBindJSON(&userVocab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uvc.UserVocabularyService.UpdateUserVocabulary(&userVocab); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userVocab)
}

func (uvc *UserVocabularyController) GetUserVocabularyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user vocabulary id"})
		return
	}
	userVocab, err := uvc.UserVocabularyService.GetUserVocabularyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userVocab)
}

func (uvc *UserVocabularyController) GetUserVocabulariesByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	userVocabs, err := uvc.UserVocabularyService.GetUserVocabulariesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userVocabs)
}
