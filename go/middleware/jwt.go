package middleware

import (
	"Campus-forum-system/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("agriculture")

type Claims struct {
	// 在payload里面放入用户id
	UserID uint
	jwt.RegisteredClaims
}

// 发token
func ReleaseToken(userID uint) (string, error) {
	// 过期时间
	expirationTime := time.Now().Add(24 * time.Hour * time.Duration(model.TokenExpireDays))
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  "Admin",
			Subject: "user token",
			// 之前版本是没有的， jwtv4之前这里有漏洞
			NotBefore: jwt.NewNumericDate(time.Now()),     // “nbf”（不是之前）声明标识 JWT 之前的时间不得接受处理。
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 此声明可用于确定 JWT 的年龄
			ExpiresAt: jwt.NewNumericDate(expirationTime), // “exp”（过期时间）声明标识或在此之后，不得接受 JWT 进行处理。
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //NewWithClaims 使用指定的签名方法和声明创建新的令牌。
	tokenString, err := token.SignedString(jwtKey)             // SignedString 创建并返回一个完整的签名 JWT。令牌使用令牌中指定的 SigningMethod 进行签名。
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	var token *jwt.Token
	var err error
	token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
