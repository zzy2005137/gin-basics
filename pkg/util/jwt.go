// package util

// import (
// 	"time"

// 	"gin-example/pkg/setting"

// 	jwt "github.com/dgrijalva/jwt-go"
// )

// var jwtSecret = []byte(setting.AppSetting.JwtSecret)

// //header 和 payload
// type Claims struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	jwt.StandardClaims
// }

// //生成 JWT
// func GenerateToken(username, password string) (string, error) {
// 	nowTime := time.Now()
// 	expireTime := nowTime.Add(3 * time.Hour)

// 	claims := Claims{
// 		username,
// 		password,
// 		jwt.StandardClaims{
// 			ExpiresAt: expireTime.Unix(),
// 			Issuer:    "gin-blog",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenWithSignature, err := token.SignedString(jwtSecret)

// 	return tokenWithSignature, err
// }

// //解析 JWT
// func ParseToken(token string) (*Claims, error) {
// 	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
// 			return claims, nil
// 		}
// 	}

// 	return nil, err
// }
