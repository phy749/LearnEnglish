package repository

import (
	"errors"

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

func (r *AuthRepository) Register(user dataoject.Register) (dataoject.Register, error) {
	// Check if username already exists
	var existingUser dataoject.Register
	if err := r.db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return dataoject.Register{}, errors.New("username already exists")
	}

	// Check if email already exists
	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return dataoject.Register{}, errors.New("email already exists")
	}

	// Create new user
	if err := r.db.Create(&user).Error; err != nil {
		return dataoject.Register{}, err
	}

	return user, nil
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
