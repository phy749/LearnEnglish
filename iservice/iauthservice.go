package iservice

import (
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/model"
)

type IAuthService interface {
	CreateUser(req dataoject.Register) (model.Useraccount, error)
	Login(req dataoject.LoginRequest) (string, error)
	ChangePassword(req dataoject.ChangePasswordRequest, id int) error
	RefreshToken(req dataoject.RefreshToken) string
	Logout(userID string) error
	SendResetLink(email string) error
}
