package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type UserVocabularyRepository struct {
	DB *gorm.DB
}

func CreateUserVocabularyRepository(db *gorm.DB) irepository.IUserVocabularyRepository {
	return &UserVocabularyRepository{DB: db}
}

func (r *UserVocabularyRepository) AddUserVocabulary(userVocab *model.UserVocabulary) error {
	return r.DB.Create(&userVocab).Error
}

func (r *UserVocabularyRepository) RemoveUserVocabulary(id int64) error {
	return r.DB.Delete(&model.UserVocabulary{}, id).Error
}

func (r *UserVocabularyRepository) UpdateUserVocabulary(userVocab *model.UserVocabulary) error {
	return r.DB.Save(userVocab).Error
}

func (r *UserVocabularyRepository) GetUserVocabularyByID(id int64) (*model.UserVocabulary, error) {
	var userVocab model.UserVocabulary
	err := r.DB.First(&userVocab, id).Error
	if err != nil {
		return nil, err
	}
	return &userVocab, nil
}

func (r *UserVocabularyRepository) GetUserVocabulariesByUserID(userID int64) ([]model.UserVocabulary, error) {
	var userVocabs []model.UserVocabulary
	err := r.DB.Where("user_id = ?", userID).Find(&userVocabs).Error
	return userVocabs, err
}
