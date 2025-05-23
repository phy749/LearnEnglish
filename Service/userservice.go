package service

import (
	"errors"

	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
)

type UserService struct {
	UserRepo irepository.IUserRepository
}

// Constructor
func NewUserService(repo irepository.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

// Triển khai các method của interface IUserService

func (s *UserService) GetAllUser() ([]model.Useraccount, error) {
	return s.UserRepo.FindAll()
}

func (s *UserService) CreateUser(req dataoject.User) (model.Useraccount, error) {
	user := model.Useraccount{
		Username:  req.Username,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Password:  req.Password, // Nên mã hóa trước khi lưu
		Birthdate: req.Birthdate,
		Phone:     req.Phone,
		Gender:    req.Gender,
	}
	return s.UserRepo.Create(user)
}

func (s *UserService) DeactivateUser(id int) (model.Useraccount, error) {
	user, err := s.UserRepo.FindByID(int(id))
	if err != nil {
		return model.Useraccount{}, err
	}
	n := "N"
	user.Is_active = &n
	return s.UserRepo.Update(user)
}

func (s *UserService) ChangePassword(req dataoject.ChangePasswordRequest, id int) (model.Useraccount, error) {
	user, err := s.UserRepo.FindByID(int(id))
	if err != nil {
		return model.Useraccount{}, err
	}
	if req.Password != req.ConfirmPassword {
		return model.Useraccount{}, errors.New("password confirmation does not match")
	}
	user.Password = req.Password // Nên mã hóa trước khi lưu
	return s.UserRepo.Update(user)
}

func (s *UserService) UpdateUser(req dataoject.User) (model.Useraccount, error) {
	// Tìm user theo email (giả sử email là unique)
	users, err := s.UserRepo.FindAll()
	if err != nil {
		return model.Useraccount{}, err
	}
	var user model.Useraccount
	found := false
	for _, u := range users {
		if u.Email == req.Email {
			user = u
			found = true
			break
		}
	}
	if !found {
		return model.Useraccount{}, errors.New("user not found")
	}
	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Password = req.Password
	user.Birthdate = req.Birthdate
	user.Phone = req.Phone
	user.Gender = req.Gender
	return s.UserRepo.Update(user)
}

func (s *UserService) FindUserById(id int) (model.Useraccount, error) {
	return s.UserRepo.FindByID(int(id))
}
