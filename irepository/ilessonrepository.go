package irepository

import "github.com/phy749/LearnEnglish/model"

type ILessonRepository interface {
	AddLesson(lesson *model.Lesson) error
	RemoveLesson(id int64) error
	UpdateLesson(lesson *model.Lesson) error
	GetLessonByID(id int64) (*model.Lesson, error)
	GetLessonsByLevelID(levelID int64) ([]model.Lesson, error)
}
