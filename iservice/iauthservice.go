package iservice

import (
	"github.com/phy749/LearnEnglish/dataoject"
)

type IAuthService interface {
	CreateUser(req dataoject.Register) (dataoject.Register, error)
	Login(req dataoject.LoginRequest) (string, string, error)
	ChangePassword(req dataoject.ChangePasswordRequest) (string, error)
	RefreshToken(req dataoject.RefreshToken) (string, error)
	Logout(userID string) (string, error)
}
