package repository

import (
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) irepository.IUserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]model.Useraccount, error) {
	var users []model.Useraccount
	err := r.DB.Preload("Role").Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id int) (model.Useraccount, error) {
	var user model.Useraccount
	err := r.DB.Preload("Role").First(&user, id).Error
	return user, err
}

func (r *UserRepository) Create(user model.Useraccount) (model.Useraccount, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *UserRepository) Update(user model.Useraccount) (model.Useraccount, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *UserRepository) Delete(id int) error {
	return r.DB.Delete(&model.Useraccount{}, id).Error
}

func (r *UserRepository) FindUserByUsername(username string) (model.Useraccount, error) {
	var user model.Useraccount
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}

// func (r *UserRepository) FindUserByEmail(email string) (model.Useraccount, error) {
// 	var user model.Useraccount
// 	err := r.DB.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		return model.Useraccount{}, err
// 	}
// 	return user, nil
// }
