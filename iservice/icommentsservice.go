package iservice

import "github.com/phy749/LearnEnglish/model"

type ICommentsService interface {
	AddComment(comment *model.Comment) error
	RemoveComment(id int64) error
	UpdateComment(comment *model.Comment) error
	GetCommentByID(id int64) (*model.Comment, error)
	GetCommentsByLessonID(lessonID int64) ([]model.Comment, error)
	GetCommentsByExerciseID(exerciseID int64) ([]model.Comment, error)
	GetReplies(parentCommentID int64) ([]model.Comment, error)
}
