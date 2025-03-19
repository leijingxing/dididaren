package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your-secret-key")

// Claims 自定义JWT声明
type Claims struct {
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时过期
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
