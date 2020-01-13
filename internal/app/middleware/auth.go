package middleware

import (
	"github.com/gin-gonic/gin"
	"music-pc-server/internal/app/plus"
)

/*
	校验Token
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if t := plus.GetToken(context); t != "" {

		}
	}
}
