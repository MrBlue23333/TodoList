package api

import (
	"demo/pkg/utils"
	"demo/service"
	"demo/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserRegisterReq
		//解析context中的JSON，将解析结果绑定到req结构体，如果内容不是JSON将返回ErrorResponse并写日志
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			//
			l := service.GetUserSrv()
			resp, err := l.UserRegister(c.Request.Context(), &req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			c.JSON(http.StatusOK, resp)
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserLoginReq
		//解析context中的JSON，将解析结果绑定到req结构体，如果内容不是JSON将返回ErrorResponse并写日志
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			l := service.GetUserSrv()
			resp, err := l.UserLogin(c.Request.Context(), &req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			c.JSON(http.StatusOK, resp)
		}
	}
}
