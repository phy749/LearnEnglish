package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type ExercisersController struct {
	ExercisesService iservice.IExercisesService
}

func NewExercisersController(service iservice.IExercisesService) *ExercisersController {
	return &ExercisersController{ExercisesService: service}
}

func (ec *ExercisersController) AddExercise(c *gin.Context) {
	var ex model.Exercise
	if err := c.ShouldBindJSON(&ex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.ExercisesService.AddExercise(&ex); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ex)
}

func (ec *ExercisersController) RemoveExercise(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exercise id"})
		return
	}
	if err := ec.ExercisesService.RemoveExercise(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted"})
}

func (ec *ExercisersController) UpdateExercise(c *gin.Context) {
	var ex model.Exercise
	if err := c.ShouldBindJSON(&ex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.ExercisesService.UpdateExercise(&ex); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ex)
}

func (ec *ExercisersController) GetExerciseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exercise id"})
		return
	}
	ex, err := ec.ExercisesService.GetExerciseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ex)
}

func (ec *ExercisersController) GetExercisesByLessonID(c *gin.Context) {
	lessonIDStr := c.Param("lesson_id")
	lessonID, err := strconv.ParseInt(lessonIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson id"})
		return
	}
	exercises, err := ec.ExercisesService.GetExercisesByLessonID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exercises)
}
