package model

import (
	"time"
)

type CompletionStatus string

const (
	Completed  CompletionStatus = "completed"
	InProgress CompletionStatus = "in_progress"
	NotStarted CompletionStatus = "not_started"
)

type UserProgress struct {
	ID               int64            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           int64            `gorm:"column:user_id" json:"user_id"`
	LessonID         int64            `gorm:"column:lesson_id" json:"lesson_id"`
	LevelID          int64            `gorm:"column:level_id" json:"level_id"`
	CompletionStatus CompletionStatus `gorm:"type:enum('completed','in_progress','not_started');default:'not_started'" json:"completion_status"`
	Score            int              `gorm:"column:score" json:"score"`
	LastAccessedAt   time.Time        `gorm:"column:last_accessed_at;autoCreateTime" json:"last_accessed_at"`
	CompletedAt      *time.Time       `gorm:"column:completed_at" json:"completed_at"`
}
