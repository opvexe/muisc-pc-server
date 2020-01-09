package routes

import (
	"github.com/gin-gonic/gin"
	"music-pc-server/handle"
	"music-pc-server/middleware"
	"music-pc-server/util"
)
/*
	注册路由
 */
func Register() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	initWithUserRouter(r)
	r.NoRoute(func(context *gin.Context) {
		util.Error(context,"请求的路径不存在")
	})
	return r
}

func initWithUserRouter(r *gin.Engine) {
	ug := r.Group("/user")
	ug.POST("/login", handle.Login)
}
