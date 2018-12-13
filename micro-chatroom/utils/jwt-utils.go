package utils

import (
	"time"

	upb "github.com/ddosakura/go-micro-demo/micro-chatroom/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

// 可通过环境变量修改
var jwtPrivateKey = []byte("#&$43^64%&*")

// 自定义的 metadata
type CustomClaims struct {
	User *upb.User
	// 使用标准的 payload
	jwt.StandardClaims
}

func DecodeJWT(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtPrivateKey, nil
		})
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func EncodeJWT(user *upb.User) (string, error) {
	expireTime := time.Now().Add(time.Second * 60).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user", // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(jwtPrivateKey)
}
