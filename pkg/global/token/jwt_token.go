package token

import (
	"github.com/dgrijalva/jwt-go"
	"gomall/config"
	"time"
)

// MemberToken 用户认证token
type MemberToken struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.GetConfig().Token.ExpireTime * time.Second)
	issuer := username
	claims := MemberToken{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().Token.Key))
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*MemberToken, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MemberToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Token.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MemberToken); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
