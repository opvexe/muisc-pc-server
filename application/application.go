package application

import (
	"context"
	"music-pc-server/routes"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {

}

/*
	初始化
 */
func NewApplication() *Application {
	return &Application{}
}

/*
	启动服务
*/
func (this *Application) Start() {

	r  := routes.Register()

	srv := &http.Server{
		Addr:"",
		Handler:r,
	}

	go func() {
		if err := srv.ListenAndServe();err!=nil{
			logrus.Error("START SERVER ERROR:",err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
	<- quit
	logrus.Infof("SHUTDOWN SERVER ...")

	ctx,cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	if err:=srv.Shutdown(ctx);err!=nil{
		logrus.Error("SERVER SHUTDOWN ERROR:",err)
	}
	logrus.Infof("SERVER EXITING")
	os.Exit(0)
}


