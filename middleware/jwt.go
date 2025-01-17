package middleware

import (
	"file_service/model/common/requests"
	"file_service/utils"
	"github.com/gin-gonic/gin"
)

// var jwtService = router.RouterGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetToken(c)
		if token == "" {
			requests.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			requests.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		if claims.UserId == 0 {
			requests.NoAuth("非法访问", c)
			c.Abort()
		}
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
