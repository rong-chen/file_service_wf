package middleware

import (
	"file_service/model/common/requests"
	"github.com/gin-gonic/gin"
)

func ValidUserIsSuperManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("authorityId")
		if id == 888 {
			c.Next()
		} else {
			requests.NoAuth("非法访问", c)
			c.Abort()
		}
	}
}

func ValidUserIsOrdinaryUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("authorityId")
		if id == 88 {
			c.Next()
		} else {
			requests.NoAuth("非法访问", c)
			c.Abort()
		}
	}
}

func ValidUserIsMusicUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("authorityId")
		if id == 8 {
			c.Next()
		} else {
			requests.NoAuth("非法访问", c)
			c.Abort()
		}
	}
}
