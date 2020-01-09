package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	cors跨域
*/
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method :=c.Request.Method
		var isOpenCore = true
		if isOpenCore {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK,nil)
		}
		c.Next()
	}
}


