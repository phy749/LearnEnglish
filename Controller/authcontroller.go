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
	c.JSON(http.StatusOK, gin.H{
		"message": "Đăng kí thành công",
		"data":    user,
	})
}

func (uc *AuthController) Login(c *gin.Context) {
	var request dataoject.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshtoken, accesstoken, err := uc.AuthService.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"refresh_token",
		refreshtoken,
		7*24*60*60, // 7 days
		"/",
		"",
		false,
		true, // HttpOnly
	)

	// Return access token in response body
	c.JSON(http.StatusOK, gin.H{
		"access_token": accesstoken,
		"message":      "Đăng nhập thành công",
	})
}

func (uc *AuthController) ChangePassword(c *gin.Context) {
	var req dataoject.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := uc.AuthService.ChangePassword(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": message})
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
	message, err := a.AuthService.Logout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": message})
}
