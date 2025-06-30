package irepository

import "github.com/phy749/LearnEnglish/model"

type IQuizzesRepository interface {
	AddQuiz(quiz *model.Quizzes) error
	RemoveQuiz(id int64) error
	UpdateQuiz(quiz *model.Quizzes) error
	GetQuizByID(id int64) (*model.Quizzes, error)
	GetQuizzesByLessonID(lessonID int64) ([]model.Quizzes, error)
}
