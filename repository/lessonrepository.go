package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type LessonRepository struct {
	DB *gorm.DB
}

func CreateLessonRepository(db *gorm.DB) irepository.ILessonRepository {
	return &LessonRepository{DB: db}
}

func (r *LessonRepository) AddLesson(lesson *model.Lesson) error {
	return r.DB.Create(lesson).Error
}

func (r *LessonRepository) RemoveLesson(id int64) error {
	return r.DB.Delete(&model.Lesson{}, id).Error
}

func (r *LessonRepository) UpdateLesson(lesson *model.Lesson) error {
	return r.DB.Save(lesson).Error
}

func (r *LessonRepository) GetLessonByID(id int64) (*model.Lesson, error) {
	var lesson model.Lesson
	err := r.DB.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *LessonRepository) GetLessonsByLevelID(levelID int64) ([]model.Lesson, error) {
	var lessons []model.Lesson
	err := r.DB.Where("level_id = ?", levelID).Find(&lessons).Error
	return lessons, err
}
