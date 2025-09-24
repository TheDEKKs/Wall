package pkg

import (
		"os"
		"fmt"
		"time"
		"github.com/dgrijalva/jwt-go"
)


type tokenJWT struct{
	Password string 
	Name string 
	jwt.StandardClaims
}

var sec = []byte(os.Getenv("secretKey"))

func JwtCreate(UserName, PasswordUs string) (string, error){

	userTOKENgen := tokenJWT {
		Password: PasswordUs,
		Name: UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userTOKENgen)

	return token.SignedString(sec)

}

func ValidateToken(tokenStrUser string) (tokenJWT, error){
	tokenSTR := &tokenJWT{}

	token, err := jwt.ParseWithClaims(tokenStrUser, tokenSTR, func(t *jwt.Token) (interface{}, error) {
		return sec, nil
	})

	if err != nil {
		return tokenJWT{}, err
	}

	if !token.Valid {
		return tokenJWT{}, fmt.Errorf("Token invalid") 
	}

	return *tokenSTR, nil

}