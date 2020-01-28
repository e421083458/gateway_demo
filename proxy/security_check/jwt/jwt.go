package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SignKey = "my_sign_key"
)

type MyCustomClaims struct {
	Foo string `json:"foo"` //自定义字段
	jwt.StandardClaims
}

func Decode(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Println("expiresAt", claims.StandardClaims.ExpiresAt)
		return claims.Foo, nil
	} else {
		return "", err
	}
}

func Encode(foo string) (string, error) {
	mySigningKey := []byte(SignKey)
	claims := MyCustomClaims{
		foo,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
