package grpc

import (
	"github.com/smallnest/rpcx/client"
	"sync"
)

var (
	LogicRpcClient client.XClient
	once           sync.Once
	RpcLogicObj    *RpcLogic
)

/*
	定义空结构体
*/
type RpcLogic struct {
}

/*
	初始化单利
*/
func InitLogicRpcClient() {
	once.Do(func() {
		d := client.NewEtcdV3Discovery("", "", []string{}, nil)
		LogicRpcClient = client.NewXClient("", client.Failtry, client.RandomSelect, d, client.DefaultOption)
		RpcLogicObj = new(RpcLogic)
	})
	if LogicRpcClient == nil {
		panic("get logic grpc client nil")
	}
}
