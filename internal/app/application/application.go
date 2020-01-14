package application

import (
	"context"
	"fmt"
	"log"
	"music-pc-server/internal/app/config"
	"music-pc-server/internal/app/routes"
	"music-pc-server/pkg/logs"
	"music-pc-server/pkg/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "../config/config.toml"

type Application struct {
}

func NewApplication() *Application {
	return new(Application)
}

/*
	启动服务
*/
func (this *Application) Start() {
	//获取配置文件
	err := config.LoadGlobal(configPath)
	util.HandleError(err)
	cfg := config.Global()

	//初始化日志
	logs.InitLog()

	log.Printf("服务启动，运行模式：%s，版本号：%s，进程号：%d，端口号：%d", cfg.RunMode, cfg.Version, os.Getpid(), cfg.HTTP.Port)

	//初始化http服务
	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      routes.InitWithWeb(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("START SERVER ERROR:", err)
		}
	}()

	//优雅的关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
	<-quit
	log.Println("SHUTDOWN SERVER ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("SERVER SHUTDOWN ERROR:", err)
	}
	log.Println("SERVER EXITING")
	os.Exit(0)
}
