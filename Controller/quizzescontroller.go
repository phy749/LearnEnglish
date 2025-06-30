package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type QuizzesController struct {
	QuizzesService iservice.IQuizzesService
}

func NewQuizzesController(service iservice.IQuizzesService) *QuizzesController {
	return &QuizzesController{QuizzesService: service}
}

func (qc *QuizzesController) AddQuiz(c *gin.Context) {
	var quiz model.Quizzes
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := qc.QuizzesService.AddQuiz(&quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

func (qc *QuizzesController) RemoveQuiz(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz id"})
		return
	}
	if err := qc.QuizzesService.RemoveQuiz(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted"})
}

func (qc *QuizzesController) UpdateQuiz(c *gin.Context) {
	var quiz model.Quizzes
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := qc.QuizzesService.UpdateQuiz(&quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

func (qc *QuizzesController) GetQuizByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz id"})
		return
	}
	quiz, err := qc.QuizzesService.GetQuizByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

func (qc *QuizzesController) GetQuizzesByLessonID(c *gin.Context) {
	lessonIDStr := c.Param("lesson_id")
	lessonID, err := strconv.ParseInt(lessonIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson id"})
		return
	}
	quizzes, err := qc.QuizzesService.GetQuizzesByLessonID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}
