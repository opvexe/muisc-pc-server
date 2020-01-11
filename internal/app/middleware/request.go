package middleware

import "github.com/gin-gonic/gin"

/*
	未找到请求方法的处理函数
 */
func NoMethodHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

/*
	身份授权
 */
