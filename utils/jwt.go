package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

var (
	AccessTokenDuration  = time.Minute * 15
	RefreshTokenDuration = time.Hour * 24 * 7
)

// create access token
func CreateAccessToken(userID int, secretkey []byte) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenObj.SignedString(secretkey)
}

// create refresh token
func CreateRefreshToken(userID int, secretkey []byte) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenObj.SignedString(secretkey)
}

// verify or parse token
func VerifyToken(tokenStr string, secretkey []byte) (*MyClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenObj.Claims.(*MyClaims)
	if !ok || !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
