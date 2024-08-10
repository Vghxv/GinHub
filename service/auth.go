package service

import (
	"fmt"
	"time"

	"github.com/Vghxv/GinHub/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)


type CustomClaims struct {
	Credentials string `json:"credentials"`
	jwt.StandardClaims
}

func GenerateJWT(username string, credentials string) (string, error) {
	claims := CustomClaims{
		Credentials: credentials,
		StandardClaims: jwt.StandardClaims{
			Subject: username,
			Issuer: setting.JWTSetting.Issuer,
			Audience: "dummy audience",
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(setting.JWTSetting.Expires)).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(setting.JWTSetting.Key))
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %v", err)
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.JWTSetting.Key), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing JWT token: %v", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing JWT token: invalid claims")
	}

	return claims, nil
}
