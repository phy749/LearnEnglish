package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type UserProgressRepository struct {
	DB *gorm.DB
}

func CreateUserProgressRepository(db *gorm.DB) irepository.IUserProgressRepository {
	return &UserProgressRepository{DB: db}
}

func (r *UserProgressRepository) AddUserProgress(progress *model.UserProgress) error {
	return r.DB.Create(progress).Error
}

func (r *UserProgressRepository) RemoveUserProgress(id int64) error {
	return r.DB.Delete(&model.UserProgress{}, id).Error
}

func (r *UserProgressRepository) UpdateUserProgress(progress *model.UserProgress) error {
	return r.DB.Save(progress).Error
}

func (r *UserProgressRepository) GetUserProgressByID(id int64) (*model.UserProgress, error) {
	var progress model.UserProgress
	err := r.DB.First(&progress, id).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *UserProgressRepository) GetProgressByUserID(userID int64) ([]model.UserProgress, error) {
	var progresses []model.UserProgress
	err := r.DB.Where("user_id = ?", userID).Find(&progresses).Error
	return progresses, err
}
