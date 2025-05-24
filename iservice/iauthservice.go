package iservice

import (
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/model"
)

type IAuthService interface {
	CreateUser(req dataoject.Register) (model.Useraccount, error)
	Login(req dataoject.LoginRequest) (string, error)
	ChangePassword(req dataoject.ChangePasswordRequest) (string, error)
	RefreshToken(req dataoject.RefreshToken) (string, error)
	Logout(userID string) (string, error)
	SendResetLink(email string) (string, error)
}
