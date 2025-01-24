package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTAPP = new(JWT)

type JWT struct{}

var (
	key = []byte("qingYu")
	t   *jwt.Token
	s   string
)

type Claims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(userId uint) (string, error) {

	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "example.com",
		},
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ = t.SignedString(key)
	return s, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	// 解析 JWT，校验签名
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回用于验证签名的密钥
		return key, nil
	})
	// 检查 Token 是否有效
	if c, ok := token.Claims.(*Claims); ok {
		return c, nil
	} else {
		return &Claims{}, err
	}
}

func GetToken(c *gin.Context) string {
	return c.GetHeader("q-token")
}
