package utils

import (
	"chat-service/internal/config"
	"chat-service/internal/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func AuthenticationResponse(userID string, cfg config.Config) (*model.AuthenticationResponse, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")

	expirationTime := time.Now().In(location).Add(time.Duration(cfg.JwtExpiration) * time.Hour)
	refreshTokenExpirationTime := time.Now().In(location).Add(time.Duration(cfg.JwtExpiration) * time.Hour)

	accessToken, err := GenerateToken(userID, cfg.JWTSecret, expirationTime)
	if err != nil {
		return nil, err
	}

	refreshStr := fmt.Sprintf("%s_%s", userID, strconv.FormatInt(time.Now().Unix(), 10))
	refreshToken, err := GenerateToken(refreshStr, cfg.JWTSecret, refreshTokenExpirationTime)
	if err != nil {
		return nil, err
	}

	return &model.AuthenticationResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		ExpiresIn:        expirationTime,
		RefreshExpiresIn: refreshTokenExpirationTime,
	}, nil
}

func GenerateToken(userID string, jwtKey string, expireTimes time.Time) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTimes.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func ParseToken(tokenStr string, jwtKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
