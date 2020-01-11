package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	MSC_SUCCESS    = 2000
	MSC_FAILED     = 4001
	MSC_UNKNOWERR  = 4501
	MSC_SESSIONERR = 4101
)

var recodeDict = map[int]string{
	MSC_SUCCESS:    "success",
	MSC_FAILED:     "fail",
	MSC_UNKNOWERR:  "unKnow error",
	MSC_SESSIONERR: "session error",
}

/*
	ParseJSON 解析请求JSON
*/
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}

/*
	ParseQuery 解析Query参数
*/
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}

/*
	ParseForm 解析Form请求
*/
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}
