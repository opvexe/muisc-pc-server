package application

import (
	"fmt"
)

type Application struct {

}

/*
	启动服务
*/
func (this *Application) Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
		}
	}()
	fmt.Println("Start....")
}
