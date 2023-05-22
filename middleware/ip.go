package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

/*
   @File: ip.go
   @Author: khaosles
   @Time: 2023/5/22 21:40
   @Desc:
*/

func GetIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// GetIP returns the client IP and whether the request has been
		ip := c.ClientIP()

		// Check the X-Forwarded-For header to see if the request
		if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
			parts := strings.Split(xff, ",")
			if len(parts) > 0 {
				ip = strings.TrimSpace(parts[0])
			}
		}
		c.Set("ip", ip)
		c.Next()
	}
}
