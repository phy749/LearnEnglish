package repository

import "gorm.io/gorm"

type LevelRepository struct {
	DB *gorm.DB
}
