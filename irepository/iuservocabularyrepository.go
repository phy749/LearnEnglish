package irepository

import (
	"github.com/phy749/LearnEnglish/model"
)

type IUserVocabularyRepository interface {
	AddUserVocabulary(userVocab *model.UserVocabulary) error
	RemoveUserVocabulary(id int64) error
	UpdateUserVocabulary(userVocab *model.UserVocabulary) error
	GetUserVocabularyByID(id int64) (*model.UserVocabulary, error)
	GetUserVocabulariesByUserID(userID int64) ([]model.UserVocabulary, error)
}
