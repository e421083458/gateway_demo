package public

import (
	"errors"
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

func JwtDecode(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	//fmt.Println("token.Claims", token.Claims)
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("request expired")
		}
		return claims, nil
	}
	return nil, errors.New("token is not StandardClaims")
}

//jwt加密
//claims := jwt.StandardClaims{
//	ExpiresAt: time.Now().Add(time.Second * 20).Unix(),
//	Issuer:    appID,
//}
func JwtEncode(claims *jwt.StandardClaims) (string, error) {
	mySigningKey := []byte(SignKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
