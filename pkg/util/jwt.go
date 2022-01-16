package util

import (
	"gin-example/pkg/setting"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//header 和 payload
type Claim struct {
	UserName string
	Password string
	jwt.StandardClaims
}

var jwtSecretKey = []byte(setting.AppSetting.JwtSecret)

//生成JWT
func GenerateJwt(username, password string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	// Create the Claims
	claims := Claim{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//添加签名，生成完整token
	ss, err := token.SignedString(jwtSecretKey)
	if err != nil {
		log.Fatal("生成token失败")
	}
	// fmt.Printf("%v %v", ss, err)
	return ss
}

func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
