package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music-pc-server/internal/app/plus"
	"music-pc-server/pkg/logs"
	"music-pc-server/pkg/util"
	"strconv"
	"time"
)

type DataAnalysisModel struct {
	RequestTime string `json:"datetime"`	//请求时间
	RequstURL string `json:"url"`			//请求url
	UserId   int `json:"userid"`			//用户id
}

var DataAnalysisCh chan DataAnalysisModel

func init()  {
	DataAnalysisCh = make(chan DataAnalysisModel, 1000)
	go HandleChannel()
}

/*
	恢复崩溃中间件
*/

func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if logs.DBCNormalLogger!=nil{
					logs.DBCNormalLogger.Error("API异常"+fmt.Sprint(err))
				}
				plus.RespError(context,plus.MSC_ServerError)
			}
		}()
		//上报用户活跃状态
		go SendDataAnalysis(context)
		context.Next()
	}
}

/*
	上报用户状态
*/
func SendDataAnalysis(c  *gin.Context)  {
	//异常崩溃
	defer func() {
		if err := recover(); err != nil {
			plus.RespError(c,plus.MSC_ServerError)
			return
		}
	}()
	//获取上报信息
	var d  DataAnalysisModel
	d.RequestTime = util.TimeTransformDateString(time.Now())
	d.RequstURL = c.Request.URL.Path
	d.UserId,_= strconv.Atoi(c.Request.Header.Get("userid"))
	DataAnalysisCh<-d
}


/*
	处理数据
*/
func HandleChannel() {
	for {
		select {
		case c :=<-DataAnalysisCh:
			_, err := util.CustomHttpRequest("POST","",c)
			if err!=nil {
				logs.DBCNormalLogger.Error(err.Error())
			}
		default:
			time.Sleep(10*time.Second)  //当前没有数据处理
		}
	}
}