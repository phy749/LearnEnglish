package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/iservice"
	"github.com/phy749/LearnEnglish/model"
)

type CommentsController struct {
	CommentsService iservice.ICommentsService
}

func NewCommentsController(service iservice.ICommentsService) *CommentsController {
	return &CommentsController{CommentsService: service}
}

func (cc *CommentsController) AddComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.CommentsService.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (cc *CommentsController) RemoveComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment id"})
		return
	}
	if err := cc.CommentsService.RemoveComment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}

func (cc *CommentsController) UpdateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.CommentsService.UpdateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (cc *CommentsController) GetCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment id"})
		return
	}
	comment, err := cc.CommentsService.GetCommentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (cc *CommentsController) GetCommentsByLessonID(c *gin.Context) {
	lessonIDStr := c.Param("lesson_id")
	lessonID, err := strconv.ParseInt(lessonIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson id"})
		return
	}
	comments, err := cc.CommentsService.GetCommentsByLessonID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (cc *CommentsController) GetCommentsByExerciseID(c *gin.Context) {
	exerciseIDStr := c.Param("exercise_id")
	exerciseID, err := strconv.ParseInt(exerciseIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exercise id"})
		return
	}
	comments, err := cc.CommentsService.GetCommentsByExerciseID(exerciseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (cc *CommentsController) GetReplies(c *gin.Context) {
	parentIDStr := c.Param("parent_id")
	parentID, err := strconv.ParseInt(parentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent comment id"})
		return
	}
	replies, err := cc.CommentsService.GetReplies(parentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, replies)
}
