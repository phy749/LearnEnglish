package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func CreateCommentRepository(db *gorm.DB) irepository.ICommentsRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) AddComment(comment *model.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentRepository) RemoveComment(id int64) error {
	return r.DB.Delete(&model.Comment{}, id).Error
}

func (r *CommentRepository) UpdateComment(comment *model.Comment) error {
	return r.DB.Save(comment).Error
}

func (r *CommentRepository) GetCommentByID(id int64) (*model.Comment, error) {
	var comment model.Comment
	err := r.DB.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetCommentsByLessonID(lessonID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Where("lesson_id = ?", lessonID).Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) GetCommentsByExerciseID(exerciseID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Where("exercise_id = ?", exerciseID).Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) GetReplies(parentCommentID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Where("parent_comment_id = ?", parentCommentID).Find(&comments).Error
	return comments, err
}
