package iservice

import "github.com/phy749/LearnEnglish/model"

type IExecrisesService interface {
	AddExercise(ex *model.Exercise) error
	RemoveExercise(id int64) error
	UpdateExercise(ex *model.Exercise) error
	GetExerciseByID(id int64) (*model.Exercise, error)
	GetExercisesByLessonID(lessonID int64) ([]model.Exercise, error)
}
