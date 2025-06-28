package model

import (
	"time"
)

type Comment struct {
	ID              int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID        int64     `gorm:"column:lesson_id" json:"lesson_id"`
	ExerciseID      int64     `gorm:"column:exercise_id" json:"exercise_id"`
	UserID          int64     `gorm:"column:user_id" json:"user_id"`
	ParentCommentID int64     `gorm:"column:parent_comment_id" json:"parent_comment_id"`
	Content         string    `gorm:"type:text" json:"content"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
