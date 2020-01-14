package middleware

import "github.com/gin-gonic/gin"

/*
	不处理业务的中间件
 */
func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
