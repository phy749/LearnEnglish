package model

import (
	"time"
)

type Lesson struct {
	LessonID         int64     `gorm:"primaryKey;column:lesson_id;autoIncrement" json:"lesson_id"`
	Title            string    `gorm:"type:varchar(150)" json:"title"`
	Content          string    `gorm:"type:text" json:"content"`
	OrderInLevel     int       `gorm:"column:order_in_level" json:"order_in_level"`
	EstimatedTimeMin int       `gorm:"column:estimated_time_min" json:"estimated_time_min"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	LevelID          int64     `gorm:"column:level_id" json:"level_id"`
}
