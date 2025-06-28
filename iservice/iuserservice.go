package iservice

import (
	"github.com/phy749/LearnEnglish/dataoject"
)

type IUserService interface {
	GetAllUser() ([]dataoject.UpdateImformationUser, error)
	CreateUser(dataoject.User) (dataoject.User, error)
	UpdateUser(dataoject.UpdateImformationUser) (string, error)
	DeactivateUser(id int) (string, error)
	FindUserById(id int) (dataoject.UpdateImformationUser, error)
}
