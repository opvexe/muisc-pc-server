package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/mail"
	"music-pc-server/internal/app/config"
	"music-pc-server/internal/app/plus"
	"music-pc-server/pkg/logs"
	"music-pc-server/pkg/util"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type DataAnalysisModel struct {
	RequestTime string `json:"datetime"` //请求时间
	RequstURL   string `json:"url"`      //请求url
	UserId      int    `json:"userid"`   //用户id
}

var DataAnalysisCh chan DataAnalysisModel

func init() {
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
				if logs.DBCNormalLogger != nil {
					logs.DBCNormalLogger.Error("API异常" + fmt.Sprint(err))
				}
				/*
					崩溃发邮箱
				*/
				debugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					debugStack += v + "<br>"
				}
				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.Global().AppName)
				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", util.TimeTransformDateString(time.Now()))
				body = strings.ReplaceAll(body, "{RequestURL}", context.Request.Method+"  "+context.Request.Host+context.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", context.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", context.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", debugStack)

				cfg := config.Global().Email
				options := &mail.Options{
					MailHost: cfg.Host,
					MailPort: cfg.Port,
					MailUser: cfg.Send,
					MailPass: cfg.Pass,
					MailTo:   cfg.Recive,
					Subject:  subject,
					Body:     body,
				}
				err := mail.Send(options)
				logs.DBCNormalLogger.Error("邮箱异常" + fmt.Sprint(err))
				plus.RespError(context, plus.MSC_ServerError)
			}
		}()
		context.Next()
	}
}

/*
	上报用户信息
*/

func AnalysisMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//上报用户活跃状态
		go SendDataAnalysis(context)

		userId := context.Request.Header.Get("userid")

		sign := context.Request.Header.Get("sign")

		if userId == "" || sign == "" {
			if logs.DBCNormalLogger != nil {
				logs.DBCNormalLogger.Error("参数不全：user_id" + userId + "sign:" + sign)
			}
			plus.RespError(context, plus.MSC_ServerError)
			context.Abort()
			return
		}

		context.Next()
	}
}

/*
	上报用户状态
*/
func SendDataAnalysis(c *gin.Context) {
	//异常崩溃
	defer func() {
		if err := recover(); err != nil {
			plus.RespError(c, plus.MSC_ServerError)
			return
		}
	}()
	//获取上报信息
	var d DataAnalysisModel
	d.RequestTime = util.TimeTransformDateString(time.Now())
	d.RequstURL = c.Request.URL.Path
	d.UserId, _ = strconv.Atoi(c.Request.Header.Get("userid"))
	DataAnalysisCh <- d
}

/*
	处理数据
*/
func HandleChannel() {
	for {
		select {
		case c := <-DataAnalysisCh:
			_, err := util.CustomHttpRequest("POST", "", c)
			if err != nil {
				logs.DBCNormalLogger.Error(err.Error())
			}
		default:
			time.Sleep(10 * time.Second) //当前没有数据处理
		}
	}
}
