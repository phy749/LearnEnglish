package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type ExercrisesRepository struct {
	DB *gorm.DB
}

func CreateExercrisesRepository(db *gorm.DB) irepository.IExercisesRepository {
	return &ExercrisesRepository{DB: db}
}

func (r *ExercrisesRepository) AddExercise(ex *model.Exercise) error {
	return r.DB.Create(ex).Error
}

func (r *ExercrisesRepository) RemoveExercise(id int64) error {
	return r.DB.Delete(&model.Exercise{}, id).Error
}

func (r *ExercrisesRepository) UpdateExercise(ex *model.Exercise) error {
	return r.DB.Save(ex).Error
}

func (r *ExercrisesRepository) GetExerciseByID(id int64) (*model.Exercise, error) {
	var ex model.Exercise
	err := r.DB.First(&ex, id).Error
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

func (r *ExercrisesRepository) GetExercisesByLessonID(lessonID int64) ([]model.Exercise, error) {
	var exercises []model.Exercise
	err := r.DB.Where("lesson_id = ?", lessonID).Find(&exercises).Error
	return exercises, err
}
