package gcrypt

/*
   @File: gjwt.go
   @Author: khaosles
   @Time: 2023/2/18 15:37
   @Desc:
*/

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(userID, username, jwtSecret string) (string, error) {
	// 设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)

	payload := Payload{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// 生成签名
	token, err := tokenClaims.SignedString([]byte(jwtSecret))
	return token, err
}

// ParseToken 解析token
func ParseToken(token, jwtSecret string) (*Payload, error) {

	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	// 判断token是否有效
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Payload); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
