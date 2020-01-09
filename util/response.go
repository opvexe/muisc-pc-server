package util

import (
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MSC_SUCCESS   = 2000
	MSC_FAILED    = 4001
	MSC_UNKNOWERR = 4501
	MSC_SESSIONERR= 4101
)

var recodeDict = map[int]string{
	MSC_SUCCESS:   "success",
	MSC_FAILED:    "fail",
	MSC_UNKNOWERR: "unKnow error",
	MSC_SESSIONERR:"session error",
}

/*
	成功返回 util.Success(c,data)
 */
func Success(c *gin.Context, data interface{}) {
	Response(c,MSC_SUCCESS,"",data)
}
/*
	自定义错误 util.Success(c,err,msg)
*/
func CustomError(c *gin.Context,err error,msg string)  {
	Response(c,MSC_FAILED,errors.Wrap(err, msg).Error(),nil)
}

/*
	系统错误 util.CheckError(c,err,msg)
 */
func CheckError(c *gin.Context,code int)  {
	Response(c,code,"",nil)
}

func Error(c *gin.Context,msg string)  {
	Response(c,MSC_FAILED,msg,nil)
}

func Response(c *gin.Context, code int, msg string, data interface{}) {
	if len(msg) == 0 {
		str, ok := recodeDict[code]
		if ok {
			msg = str
		}
		msg = recodeDict[MSC_UNKNOWERR]
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
