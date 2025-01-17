package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net"
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

//// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
//func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
//	v, err, _ := global.QY_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
//		return j.CreateToken(claims)
//	})
//	return v.(string), err
//}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

//func NewJWT() *JWT {
//	return &JWT{
//		[]byte(global.QY_CONFIG.JWT.SigningKey),
//	}
//}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	return c.GetHeader("q-token")
}

// ParseToken 解析 token
//func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
//
//}

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}
