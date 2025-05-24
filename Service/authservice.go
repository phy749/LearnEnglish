package service

import (
	"errors"

	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"github.com/phy749/LearnEnglish/utils"
)

type AuthService struct {
	AuthRepo irepository.IAuthRepository
}

func NewAuthService(repo irepository.IAuthRepository) *AuthService {
	return &AuthService{AuthRepo: repo}
}

func (s *AuthService) CreateUser(req dataoject.Register) (model.Useraccount, error) {
	if req.Password != req.ConfirmPassword {
		return model.Useraccount{}, errors.New("Password and Password Confirm do not match")
	}
	hashpassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return model.Useraccount{}, err
	}
	user := model.Useraccount{
		Username:  req.Username,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Password:  hashpassword,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
		Gender:    req.Gender,
	}
	return s.AuthRepo.Register(user)
}

func (s *AuthService) Login(req dataoject.LoginRequest) (string, error) {
	// Find user by username
	user, err := s.AuthRepo.FindUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return "", errors.New("Invalid username or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ChangePassword(req dataoject.ChangePasswordRequest) (string, error) {
	// Find user by username
	user, err := s.AuthRepo.FindUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("User not found")
	}

	// Verify password and confirm password match
	if req.Password != req.ConfirmPassword {
		return "", errors.New("Password and confirm password do not match")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	_, err = s.AuthRepo.Update(user)
	if err != nil {
		return "", err
	}

	return "Change password complete", nil
}

// sai
func (s *AuthService) RefreshToken(req dataoject.RefreshToken) (string, error) {

	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		return "", errors.New("Invalid refresh token")
	}

	user, err := s.AuthRepo.FindUserByUsername(claims.Username)
	if err != nil {
		return "", errors.New("User not found")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// sai
func (s *AuthService) Logout(userID string) (string, error) {

	return "Logged out successfully", nil
}

// sai
func (s *AuthService) SendResetLink(email string) (string, error) {
	// Find user by email
	_, err := s.AuthRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("User not found")
	}
	return "Send mail complete", nil
}
