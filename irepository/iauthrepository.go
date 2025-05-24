package irepository

import (
	"github.com/phy749/LearnEnglish/model"
)

type IAuthRepository interface {
	Register(user model.Useraccount) (model.Useraccount, error)
	FindUserByUsername(username string) (model.Useraccount, error)
	FindUserByEmail(email string) (model.Useraccount, error)
	Update(user model.Useraccount) (model.Useraccount, error)
}
