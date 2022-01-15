package util

import (
	"fmt"
	"gin-example/pkg/setting"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	User     string
	Password string
	jwt.StandardClaims
}

var jwtSecret = setting.AppSetting.JwtSecret

func GenerateJwt() {
	mySigningKey := []byte(jwtSecret)

	// Create the Claims
	claims := Claim{
		"zzy",
		"1230",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "gin-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//添加签名，生成完整token
	ss, err := token.SignedString(mySigningKey)

	fmt.Printf("%v %v", ss, err)
}

func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
