package irepository

import "github.com/phy749/LearnEnglish/model"

type IUserRepository interface {
	FindAll() ([]model.Useraccount, error)
	FindByID(id int) (model.Useraccount, error)
	Create(user model.Useraccount) (model.Useraccount, error)
	Update(user model.Useraccount) (model.Useraccount, error)
	Delete(id int) error
}
