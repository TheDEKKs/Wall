package pkg

import (
		"os"
		"fmt"
		"time"
		"github.com/dgrijalva/jwt-go"
)


//Вообще я собираюсь позже зафигачить сюда не юзернейм а ID хотя может и отсавялю юзернейм
//Просто если сразу буду записывать ID то я смогу избавиться от вызова функций по поиску его 
type tokenJWT struct{
	Password string 
	Name string 
	jwt.StandardClaims
}
//Получаем секретный ключ
var sec = []byte(os.Getenv("secretKey"))

func JwtCreate(UserName, PasswordUs string) (string, error){
	//Создаем стурктуру 
	userTOKENgen := tokenJWT {
		Password: PasswordUs,
		Name: UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(),
		},
	}
	//Кодируем
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userTOKENgen)
	
	//Возращаем
	return token.SignedString(sec)

}

func ValidateToken(tokenStrUser string) (tokenJWT, error){
	tokenSTR := &tokenJWT{}
	//Декодируем 
	token, err := jwt.ParseWithClaims(tokenStrUser, tokenSTR, func(t *jwt.Token) (interface{}, error) {
		return sec, nil
	})

	//Какаято ошиибка
	if err != nil {
		return tokenJWT{}, err
	}

	//Токен невалидный
	if !token.Valid {
		return tokenJWT{}, fmt.Errorf("Token invalid") 
	}
	// Если все хорошо возращаем данные 
	return *tokenSTR, nil

}