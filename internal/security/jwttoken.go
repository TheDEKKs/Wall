package security

import (
	"fmt"
	"time"
	"thedekk/WWT/internal/env"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type tokenJWT struct{
	Password string 
	Id int
	UserID uuid.UUID 
	jwt.StandardClaims
}

func JwtCreate(password string, userID uuid.UUID) (string, error){
	var config env.Config

	config.Load()

	userTOKENgen := tokenJWT {
		Password: password,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userTOKENgen)
	
	return token.SignedString([]byte(config.SecretKey))

}

func ValidateToken(tokenStrUser string) (*tokenJWT, error){
	var config env.Config

	config.Load()

	tokenSTR := &tokenJWT{}
	token, err := jwt.ParseWithClaims(tokenStrUser, tokenSTR, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token invalid") 
	}
	return tokenSTR, nil

}