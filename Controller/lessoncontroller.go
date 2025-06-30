package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type LessonController struct {
	LessonService iservice.ILessonService
}

func NewLessonController(service iservice.ILessonService) *LessonController {
	return &LessonController{LessonService: service}
}

func (lc *LessonController) AddLesson(c *gin.Context) {
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := lc.LessonService.AddLesson(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (lc *LessonController) RemoveLesson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson id"})
		return
	}
	if err := lc.LessonService.RemoveLesson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted"})
}

func (lc *LessonController) UpdateLesson(c *gin.Context) {
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := lc.LessonService.UpdateLesson(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (lc *LessonController) GetLessonByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson id"})
		return
	}
	lesson, err := lc.LessonService.GetLessonByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (lc *LessonController) GetLessonsByLevelID(c *gin.Context) {
	levelIDStr := c.Param("level_id")
	levelID, err := strconv.ParseInt(levelIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid level id"})
		return
	}
	lessons, err := lc.LessonService.GetLessonsByLevelID(levelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}
