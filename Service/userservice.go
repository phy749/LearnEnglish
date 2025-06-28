package service

import (
	"errors"

	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/irepository"
	"github.com/phy749/LearnEnglish/model"
	"github.com/phy749/LearnEnglish/utils"
)

type UserService struct {
	UserRepo irepository.IUserRepository
}

// Constructor
func NewUserService(repo irepository.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) GetAllUser() ([]dataoject.UpdateImformationUser, error) {
	users, err := s.UserRepo.FindAll()
	if err != nil {
		return []dataoject.UpdateImformationUser{}, err
	}

	var updateUsers []dataoject.UpdateImformationUser
	for _, user := range users {
		updateUsers = append(updateUsers, dataoject.UpdateImformationUser{
			Id:        user.User_id,
			Username:  user.Username,
			Fullname:  user.Fullname,
			Email:     user.Email,
			Birthdate: user.Birthdate,
			Phone:     user.Phone,
			Gender:    user.Gender,
		})
	}
	return updateUsers, nil
}

func (s *UserService) CreateUser(req dataoject.User) (dataoject.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dataoject.User{}, err
	}
	user := model.Useraccount{
		Username:  req.Username,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Password:  hashedPassword,
		Birthdate: req.Birthdate,
		Phone:     req.Phone,
		Gender:    req.Gender,
	}
	_, err = s.UserRepo.Create(user)
	if err != nil {
		return dataoject.User{}, err
	}
	return req, err
}

func (s *UserService) DeactivateUser(id int) (string, error) {
	user, err := s.UserRepo.FindByID(int(id))
	if err != nil {
		return "Không tìm thấy user", err
	}
	n := "N"
	user.Is_active = &n
	user, err = s.UserRepo.Update(user)
	return "Xóa thành công user", err
}

func (s *UserService) UpdateUser(req dataoject.UpdateImformationUser) (string, error) {
	existingUser, err := s.UserRepo.FindByID(req.Id)
	if err != nil {
		return "", err
	}

	if userByUsername, err := s.UserRepo.FindUserByUsername(req.Username); err == nil && userByUsername.User_id != req.Id {
		return "", errors.New("username already exists")
	}

	if userByEmail, err := s.UserRepo.FindUserByEmail(req.Email); err == nil && userByEmail.User_id != req.Id {
		return "", errors.New("email already exists")
	}

	existingUser.User_id = req.Id
	existingUser.Username = req.Username
	existingUser.Fullname = req.Fullname
	existingUser.Email = req.Email
	existingUser.Birthdate = req.Birthdate
	existingUser.Phone = req.Phone
	existingUser.Gender = req.Gender

	existingUser, err = s.UserRepo.Update(existingUser)
	return "Update user thành công", err
}

func (s *UserService) FindUserById(id int) (dataoject.UpdateImformationUser, error) {
	user, err := s.UserRepo.FindByID(int(id))
	if err != nil {
		return dataoject.UpdateImformationUser{}, errors.New("Không tìm thấy User")
	}
	users := dataoject.UpdateImformationUser{
		Id:        user.User_id,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Phone:     user.Phone,
		Gender:    user.Gender,
	}
	return users, nil
}
