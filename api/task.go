package api

import (
	"demo/consts"
	"demo/pkg/utils"
	"demo/service"
	"demo/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateTaskReq
		//解析context中的JSON，将解析结果绑定到req结构体，如果内容不是JSON将返回ErrorResponse并写日志
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.CreateTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}

func ListTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListTaskReq
		//解析context中的JSON，将解析结果绑定到req结构体，如果内容不是JSON将返回ErrorResponse并写日志
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		if req.Limit == 0 {
			req.Limit = consts.BasePageSize
		}
		l := service.GetTaskSrv()
		resp, err := l.ListTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func ShowTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowTaskReq
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.ShowTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.DeleteTaskReq
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.DeleteTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}
func UpdateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateTaskReq
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.UpdateTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}
func SearchTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.SearchTaskReq
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.SearchTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}
