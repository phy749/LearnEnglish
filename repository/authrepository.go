package repository

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Register(user model.Useraccount) (dataoject.Register, error) {

	var existingUser model.Useraccount
	if err := r.db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return dataoject.Register{}, errors.New("username already exists")
	}

	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return dataoject.Register{}, errors.New("email already exists")
	}

	if err := r.db.Create(&user).Error; err != nil {
		return dataoject.Register{}, err
	}

	var result dataoject.Register
	if err := copier.Copy(&result, &user); err != nil {
		return dataoject.Register{}, err
	}

	result.Password = ""
	result.ConfirmPassword = ""

	return result, nil
}

func (r *AuthRepository) FindUserByUsername(username string) (model.Useraccount, error) {
	var user model.Useraccount
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}

func (r *AuthRepository) FindUserById(userid int) (model.Useraccount, error) {
	var user model.Useraccount
	if err := r.db.Where("user_id = ?", userid).First(&user).Error; err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}

func (r *AuthRepository) FindUserByEmail(email string) (model.Useraccount, error) {
	var user model.Useraccount
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}

func (r *AuthRepository) Update(user model.Useraccount) (model.Useraccount, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}

func (r *UserRepository) FindUserByEmail(email string) (model.Useraccount, error) {
	var user model.Useraccount
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.Useraccount{}, err
	}
	return user, nil
}
