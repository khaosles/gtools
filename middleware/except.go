package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khaosles/gtools/gresult"
)

/*
   @File: exception.go
   @Author: khaosles
   @Time: 2023/3/6 10:13
   @Desc: 全局异常捕获
*/

func Except() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Abort()
				c.JSON(http.StatusOK, gresult.JsonResult{}.CatchErr(err))
				return
			}
		}()
		c.Next()
	}
}
