package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("Blue's-very-secret-secretKey"))

type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(id int64, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Blue",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
