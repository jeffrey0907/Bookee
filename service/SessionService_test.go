package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["uid"] = 123123
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		t.Error(err.Error())
	}

	token1, err :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(token1)
}
