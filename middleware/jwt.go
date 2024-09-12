package middleware

import (
	"demo/pkg/ctl"
	"demo/pkg/e"
	"demo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.Success
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(code, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "Token is empty",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.Error
			code = http.StatusNotFound
			c.JSON(code, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "Token is invalid",
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(),
			&ctl.UserInfo{Id: claims.Id, UserName: claims.UserName}))
		c.Next()
	}
}
