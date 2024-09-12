package router

import (
	"demo/api"
	"demo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task_create", api.CreateTaskHandler())
			authed.GET("task_list", api.ListTaskHandler())
			authed.POST("task_delete", api.DeleteTaskHandler())
			authed.POST("task_update", api.UpdateTaskHandler())
			authed.GET("task_search", api.SearchTaskHandler())
			authed.GET("task_show", api.ShowTaskHandler())
		}
	}
	return ginRouter
}
