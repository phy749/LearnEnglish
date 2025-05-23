package iservice

import (
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/model"
)

type IUserService interface {
	GetAllUser() ([]model.Useraccount, error)
	CreateUser(dataoject.User) (model.Useraccount, error)
	UpdateUser(dataoject.User) (model.Useraccount, error)
	DeactivateUser(id int) (model.Useraccount, error)
	FindUserById(id int) (model.Useraccount, error)
}
