package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/herizal95/golang-jwt-mux/models"
	"github.com/satori/uuid"
)

var MySecretKey = []byte("lfjalksjlkrj1lk2jlk")

type MyCustomClaims struct {
	Uid   uuid.UUID `json:"uid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		user.Uid,
		user.Name,
		user.Email,

		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySecretKey)
	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	} else {
		fmt.Println(err)
	}
	return claims, nil
}
