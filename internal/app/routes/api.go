package routes

import (
	"github.com/gin-gonic/gin"
	"music-pc-server/internal/app/config"
	"music-pc-server/internal/app/handle"
	"music-pc-server/internal/app/middleware"
	"music-pc-server/internal/app/plus"
	"music-pc-server/pkg/util"
)

/*
	初始化路由
*/
func InitWithWeb() *gin.Engine {
	//初始化
	cfg := config.Global()
	gin.SetMode(cfg.RunMode)

	app := gin.New()
	//崩溃恢复
	app.Use(middleware.RecoveryMiddleware())

	//频率设置
	app.Use(middleware.RateLimitMiddleware())

	//跨域请求
	if cfg.CORS.Enable {
		app.Use(middleware.CorsMiddleware())
	}

	//注册/api路由
	err := registerRouter(app)
	util.HandleError(err)

	app.NoRoute(func(context *gin.Context) {
		plus.RespError(context,plus.MSC_NotFound)
	})

	return app
}

/*
	注册路由
*/
func registerRouter(app *gin.Engine) error {

	rg := app.Group("/api")

	//身份授权中间件

	//请求频率限制中间件

	v1 := rg.Group("/v1")
	{
		//注册/api/v1/user
		guser := v1.Group("user")
		{
			guser.GET(":id", handle.Get)
			guser.GET("", handle.Get)
		}

	}
	return nil
}
