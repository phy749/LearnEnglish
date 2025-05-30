package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/phy749/LearnEnglish/model"
)

// Secret keys for JWT
var AccessSecret = []byte("ACCESS_SECRET")
var RefreshSecret = []byte("REFRESH_SECRET")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateAccessToken(user model.Useraccount) (string, error) {
	claims := jwt.MapClaims{
		"sub":      user.User_id,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"username": user.Username,
		"email":    user.Email,
		"role":     user.RoleID,
	}
	log.Printf("AccessToken Claims: %+v\n", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessSecret)
}

func GenerateRefreshToken(userID string) (string, string, error) {
	jti := uuid.New().String()
	claims := jwt.MapClaims{
		"sub": userID,
		"jti": fmt.Sprintf("%d", time.Now().Unix()),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(RefreshSecret)
	return signedToken, jti, err
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return RefreshSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return AccessSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
