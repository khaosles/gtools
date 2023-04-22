package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   @File: cross.go
   @Author: khaosles
   @Time: 2023/4/12 12:07
   @Desc: 跨域
*/

// Cross 处理跨域请求,支持options访问
func Cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if len(origin) == 0 {
			c.Next()
			return
		}
		// 同源直接过
		host := c.GetHeader("Host")
		if origin == "http://"+host || origin == "https://"+host {
			c.Next()
			return
		}
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,X-Requested-With,Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// OPTIONS
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			c.Abort()
		}
		c.Next()
	}
}
