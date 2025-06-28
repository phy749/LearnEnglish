package model

import (
	"gorm.io/datatypes"
)

type Exercise struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID      int64          `gorm:"column:lesson_id" json:"lesson_id"`
	Type          string         `gorm:"type:varchar(50)" json:"type"`
	Question      string         `gorm:"type:text" json:"question"`
	CorrectAnswer string         `gorm:"type:text" json:"correct_answer"`
	Instructions  string         `gorm:"type:text" json:"instructions"`
	MediaURL      string         `gorm:"type:varchar(255)" json:"media_url"`
	ExtraData     datatypes.JSON `gorm:"type:json" json:"extra_data"`
}
