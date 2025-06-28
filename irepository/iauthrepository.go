package irepository

import (
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/model"
)

type IAuthRepository interface {
	Register(user model.Useraccount) (dataoject.Register, error)
	FindUserByUsername(username string) (model.Useraccount, error)
	FindUserById(user_id int) (model.Useraccount, error)
	FindUserByEmail(email string) (model.Useraccount, error)
	Update(user model.Useraccount) (model.Useraccount, error)
}
