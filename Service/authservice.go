package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/utils"
)

type AuthService struct {
	AuthRepo irepository.IAuthRepository
}

func NewAuthService(repo irepository.IAuthRepository) *AuthService {
	return &AuthService{AuthRepo: repo}
}

func (s *AuthService) CreateUser(req dataoject.Register) (dataoject.Register, error) {
	if req.Password != req.ConfirmPassword {
		return dataoject.Register{}, errors.New("Password and Password Confirm do not match")
	}
	hashpassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dataoject.Register{}, err
	}

	user := dataoject.Register{
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

func (s *AuthService) Login(req dataoject.LoginRequest) (string, string, error) {
	// Find user by username
	user, err := s.AuthRepo.FindUserByUsername(req.Username)
	if err != nil {
		return "", "", errors.New("Invalid username or password")
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return "", "", errors.New("Invalid username or password")
	}
	fmt.Println(user)

	uuid := strconv.Itoa(user.User_id)
	// Generate token
	accesstoken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshtoken, _, err := utils.GenerateRefreshToken(uuid)
	if err != nil {
		return "", "", err
	}
	utils.SaveRefreshToken(uuid, refreshtoken, 7*24*time.Hour)

	return refreshtoken, accesstoken, nil
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

	userID, err := strconv.Atoi(claims.StandardClaims.Subject)
	if err != nil {
		return "", errors.New("Invalid user ID")
	}

	// Tạo access token mới
	user, err := s.AuthRepo.FindUserById(userID)
	if err != nil {
		return "", errors.New("User not found")
	}

	newAccessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
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
