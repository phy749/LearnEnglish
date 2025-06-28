package service

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
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
func (s *AuthService) CreateUser(req dataoject.Register) (dataoject.Register, error) {
	err := ValidateEmail(req.Email)
	if err != nil {
		return dataoject.Register{}, err
	}
	if req.Password != req.ConfirmPassword {
		return dataoject.Register{}, errors.New("password and confirm password do not match")
	}

	if err := ValidatePassword(req.Password); err != nil {
		return dataoject.Register{}, err
	}

	if _, err := s.AuthRepo.FindUserByUsername(req.Username); err == nil {
		return dataoject.Register{}, errors.New("username already exists")
	}

	if _, err := s.AuthRepo.FindUserByEmail(req.Email); err == nil {
		return dataoject.Register{}, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dataoject.Register{}, err
	}

	var user model.Useraccount
	if err := copier.Copy(&user, &req); err != nil {
		return dataoject.Register{}, err
	}
	user.Password = hashedPassword
	user.RoleID = 1
	// fmt.Printf("User: %+v\n", "123", user)

	return s.AuthRepo.Register(user)
}

func (s *AuthService) Login(req dataoject.LoginRequest) (string, string, error) {
	// Find user by username
	fmt.Println("Attempting to find user:", req.Username)
	user, err := s.AuthRepo.FindUserByUsername(req.Username)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return "", "", errors.New("Invalid username or password")
	}

	fmt.Println("User found, checking password")
	if !utils.CheckPassword(user.Password, req.Password) {
		fmt.Println("Password check failed")
		return "", "", errors.New("Invalid username or password")
	}
	fmt.Println("Password check passed")

	uuid := strconv.Itoa(user.User_id)
	fmt.Println("Generating access token for user:", uuid)
	// Generate token
	accesstoken, err := utils.GenerateAccessToken(user)
	if err != nil {
		fmt.Println("Error generating access token:", err)
		return "", "", err
	}
	fmt.Println("Access token generated successfully")

	fmt.Println("Generating refresh token")
	refreshtoken, _, err := utils.GenerateRefreshToken(uuid)
	if err != nil {
		fmt.Println("Error generating refresh token:", err)
		return "", "", err
	}
	fmt.Println("Refresh token generated successfully")

	fmt.Println("Saving refresh token to Redis")
	err = utils.SaveRefreshToken(uuid, refreshtoken, 7*24*time.Hour)
	if err != nil {
		fmt.Println("Error saving refresh token:", err)
		return "", "", err
	}
	fmt.Println("Refresh token saved successfully")

	return refreshtoken, accesstoken, nil
}

func (s *AuthService) ChangePassword(req dataoject.ChangePasswordRequest) (string, error) {
	err := ValidateEmail(req.Email)
	if err != nil {
		return "", err
	}
	user, err := s.AuthRepo.FindUserByUsername(req.Username)

	if err != nil {
		return "", errors.New("User not found")
	}
	if user.Email != req.Email {
		return "", errors.New("User not found")
	}

	if req.Password != req.ConfirmPassword {
		return "", errors.New("Password and confirm password do not match")
	}
	if err := ValidatePassword(req.Password); err != nil {
		return "", errors.New("Password must have P8 characters: 1 uppercase, 1 lowercase, 1 number and 1 special letter")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	_, err = s.AuthRepo.Update(user)
	if err != nil {
		return "", err
	}

	return "Đổi password thành công", nil
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
func ValidatePassword(password string) error {
	// Kiểm tra độ dài
	if len(password) < 8 {
		return errors.New("Mật khẩu phải có ít nhất 8 ký tự")
	}

	// Regex kiểm tra từng điều kiện
	var (
		uppercase = regexp.MustCompile(`[A-Z]`)
		lowercase = regexp.MustCompile(`[a-z]`)
		number    = regexp.MustCompile(`[0-9]`)
		special   = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/?]`)
	)

	if !uppercase.MatchString(password) {
		return errors.New("Mật khẩu phải chứa ít nhất một chữ hoa")
	}
	if !lowercase.MatchString(password) {
		return errors.New("Mật khẩu phải chứa ít nhất một chữ thường")
	}
	if !number.MatchString(password) {
		return errors.New("Mật khẩu phải chứa ít nhất một chữ số")
	}
	if !special.MatchString(password) {
		return errors.New("Mật khẩu phải chứa ít nhất một ký tự đặc biệt")
	}

	return nil
}
func ValidateEmail(email string) error {
	// Regex đơn giản kiểm tra định dạng email
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)
	if !re.MatchString(email) {
		return errors.New("Email không hợp lệ")
	}
	return nil
}
