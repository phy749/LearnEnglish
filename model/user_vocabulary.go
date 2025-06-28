package model

import (
	"time"
)

type UserVocabulary struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int64     `gorm:"column:user_id" json:"user_id"`
	VocabularyID  int64     `gorm:"column:vocabulary_id" json:"vocabulary_id"`
	CustomWord    string    `gorm:"type:varchar(100)" json:"custom_word"`
	CustomMeaning string    `gorm:"type:text" json:"custom_meaning"`
	AddedAt       time.Time `gorm:"column:added_at;autoCreateTime" json:"added_at"`
	MasteryLevel  int       `gorm:"column:mastery_level" json:"mastery_level"`
}
