package handle

import (
	"github.com/gin-gonic/gin"
	"music-pc-server/internal/app/plus"
)

func Get(c *gin.Context) {
	plus.RespSuccess(c,"测试数据库")
}
