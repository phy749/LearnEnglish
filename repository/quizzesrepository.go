package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type QuizzesRepository struct {
	DB *gorm.DB
}

func CreateQuizzesRepository(db *gorm.DB) irepository.IQuizzesRepository {
	return &QuizzesRepository{DB: db}
}

func (r *QuizzesRepository) AddQuiz(quiz *model.Quizzes) error {
	return r.DB.Create(quiz).Error
}

func (r *QuizzesRepository) RemoveQuiz(id int64) error {
	return r.DB.Delete(&model.Quizzes{}, id).Error
}

func (r *QuizzesRepository) UpdateQuiz(quiz *model.Quizzes) error {
	return r.DB.Save(quiz).Error
}

func (r *QuizzesRepository) GetQuizByID(id int64) (*model.Quizzes, error) {
	var quiz model.Quizzes
	err := r.DB.First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *QuizzesRepository) GetQuizzesByLessonID(lessonID int64) ([]model.Quizzes, error) {
	var quizzes []model.Quizzes
	err := r.DB.Where("lesson_id = ?", lessonID).Find(&quizzes).Error
	return quizzes, err
}
