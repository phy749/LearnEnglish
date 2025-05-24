package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phy749/LearnEnglish/dataoject"
	"github.com/phy749/LearnEnglish/iservice"
)

type AuthController struct {
	AuthService iservice.IAuthService
}

func NewAuthController(authService iservice.IAuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (uc *AuthController) Register(c *gin.Context) {
	var request dataoject.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.AuthService.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *AuthController) Login(c *gin.Context) {
	var request dataoject.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uc.AuthService.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *AuthController) ChangePassword(c *gin.Context) {
	var req dataoject.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uc.AuthService.ChangePassword(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) RefreshToken(c *gin.Context) {
	var req dataoject.RefreshToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := a.AuthService.RefreshToken(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) Logout(c *gin.Context) {
	userID := c.Query("user_id")
	token, err := a.AuthService.Logout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) ForgotPassword(c *gin.Context) {
	var req dataoject.ForgotPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := a.AuthService.SendResetLink(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
