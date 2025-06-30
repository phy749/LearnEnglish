package iservice

import "github.com/phy749/LearnEnglish/model"

type IUserProgressService interface {
	AddUserProgress(progress *model.UserProgress) error
	RemoveUserProgress(id int64) error
	UpdateUserProgress(progress *model.UserProgress) error
	GetUserProgressByID(id int64) (*model.UserProgress, error)
	GetProgressByUserID(userID int64) ([]model.UserProgress, error)
}
